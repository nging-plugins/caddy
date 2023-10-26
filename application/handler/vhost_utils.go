/*
Nging is a toolbox for webmasters
Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/nging/v5/application/handler"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"

	"github.com/nging-plugins/caddymanager/application/dbschema"
	"github.com/nging-plugins/caddymanager/application/library/cmder"
	"github.com/nging-plugins/caddymanager/application/library/engine"
	"github.com/nging-plugins/caddymanager/application/library/form"
	"github.com/nging-plugins/caddymanager/application/model"
)

const configFilePrefix = `nging_`

func makeConfigFileName(cfg engine.Configer, id uint) string {
	if cfg.Engine() == `default` { // 默认引擎为 Nging 内置服务器，为 Nging 所独有，所以配置文件不用加前缀(同时保持对旧版的兼容)
		return fmt.Sprint(id) + `.conf`
	}
	// 其它引擎由用户配置，可能会将网站配置目录指向旧系统的配置目录，通过加本系统的前缀标识“nging_”来避免删掉旧配置
	return configFilePrefix + fmt.Sprint(id) + `.conf`
}

func DeleteCaddyfileByID(ctx echo.Context, serverIdent string, id uint, serverM ...*dbschema.NgingVhostServer) error {
	cfg, err := getServerConfig(ctx, serverIdent, serverM...)
	if err != nil || cfg == nil {
		return err
	}
	saveDir, err := cfg.GetVhostConfigDirAbsPath()
	if err != nil {
		return err
	}
	saveFile := filepath.Join(saveDir, makeConfigFileName(cfg, id))
	err = os.Remove(saveFile)
	if err == nil {
		item := engine.Engines.GetItem(cfg.Engine())
		if item != nil {
			err = item.X.(engine.Enginer).ReloadServer(ctx, cfg)
		}
	} else if os.IsNotExist(err) {
		err = nil
	}
	return err
}

func deleteCaddyfileByServer(ctx echo.Context, svr *dbschema.NgingVhostServer, restart bool) (err error) {
	for _, v := range engine.Engines.Slice() {
		if v.K != svr.Engine {
			continue
		}
		eng := v.X.(engine.Enginer)
		if eng == nil {
			continue
		}
		cfg := eng.BuildConfig(ctx, svr)
		var saveDir string
		saveDir, err = cfg.GetVhostConfigDirAbsPath()
		if err != nil {
			break
		}
		err = removeAllConf(cfg, saveDir)
		if err != nil {
			break
		}
		os.Remove(saveDir)
		if restart {
			err = eng.ReloadServer(ctx, cfg)
		}
	}
	return err
}

var reSplitRegexp = regexp.MustCompile(`[\s]+`)

func hasEnvVar(v string) bool {
	for _, r := range v {
		if r == '$' || r == '%' {
			return true
		}
	}
	return false
}

func generateHostURL(currentHost string, hosts string) []template.HTML {
	hosts = strings.TrimSpace(hosts)
	hostsSlice := reSplitRegexp.Split(hosts, -1)
	urls := make([]template.HTML, 0, len(hostsSlice))
	for _, v := range hostsSlice {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		parsedValue := com.ParseEnvVar(v)
		if len(parsedValue) > 0 {
			switch {
			case parsedValue[0] == ':':
				v = `<a href="http://` + currentHost + parsedValue + `" target="_blank" rel="noopener noreferrer">` + v + `</a>`
			case strings.HasPrefix(parsedValue, `0.0.0.0:`):
				v = `<a href="http://` + currentHost + strings.TrimPrefix(parsedValue, `0.0.0.0`) + `" target="_blank" rel="noopener noreferrer">` + v + `</a>`
			case !strings.Contains(parsedValue, `//`):
				v = `<a href="http://` + parsedValue + `" target="_blank" rel="noopener noreferrer">` + v + `</a>`
			default:
				v = `<a href="` + strings.ReplaceAll(parsedValue, `*`, `test`) + `" target="_blank" rel="noopener noreferrer">` + v + `</a>`
			}
		}
		urls = append(urls, template.HTML(v))
	}
	return urls
}

func removeAllConf(cfg engine.Configer, rootDir string) error {
	isDefaultEngine := cfg.Engine() == `default`
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, `.conf`) || (!isDefaultEngine && !strings.HasPrefix(info.Name(), configFilePrefix)) {
			return nil
		}
		log.Info(`Delete the WebServer configuration file: `, path)
		return os.Remove(path)
	})
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
}

func getSaveDir(cfg engine.Configer) (saveDir string, err error) {
	saveDir, err = cfg.GetVhostConfigDirAbsPath()
	if err != nil {
		return
	}
	err = com.MkdirAll(saveDir, os.ModePerm)
	return
}

func saveVhostConf(ctx echo.Context, cfg engine.Configer, id uint, values url.Values) error {
	ctx.Set(`values`, form.NewValues(values))
	b, err := ctx.Fetch(`caddy/makeconfig/`+cfg.TemplateFile(), nil)
	if err != nil {
		return err
	}
	b = com.CleanSpaceLine(b)
	saveFile, err := cfg.GetVhostConfigDirAbsPath()
	if err != nil {
		return err
	}
	if !com.FileExists(saveFile) {
		com.MkdirAll(saveFile, os.ModePerm)
	}
	saveFile = filepath.Join(saveFile, makeConfigFileName(cfg, id))
	log.Info(`Generate a `+cfg.Engine()+` configuration file: `, saveFile)
	err = os.WriteFile(saveFile, b, os.ModePerm)
	//jsonb, _ := caddyfile.ToJSON(b)
	//err = os.WriteFile(saveFile+`.json`, jsonb, os.ModePerm)
	return err
}

func getServerConfig(ctx echo.Context, serverIdent string, serverM ...*dbschema.NgingVhostServer) (engine.Configer, error) {
	var cfg engine.Configer
	if serverIdent == `default` {
		cfg = cmder.GetCaddyConfig()
	} else {
		var svrM *dbschema.NgingVhostServer
		if len(serverM) > 0 && serverM[0] != nil {
			svrM = serverM[0]
		} else {
			svrM = dbschema.NewNgingVhostServer(ctx)
			err := svrM.Get(nil, `ident`, serverIdent)
			if err != nil {
				return cfg, err
			}
		}
		item := engine.Engines.GetItem(svrM.Engine)
		if item == nil {
			return cfg, fmt.Errorf(`unsupported engine: %v`, serverIdent)
		}
		cfg = item.X.(engine.Enginer).BuildConfig(ctx, svrM)
	}
	return cfg, nil
}

func saveVhostData(ctx echo.Context, m *dbschema.NgingVhost, values url.Values, restart bool) (err error) {
	var cfg engine.Configer
	cfg, err = getServerConfig(ctx, m.ServerIdent)
	if err != nil || cfg == nil {
		return
	}
	var saveDir string
	saveDir, err = getSaveDir(cfg)
	if err != nil {
		return
	}
	if m.Disabled == `Y` {
		saveFile := filepath.Join(saveDir, makeConfigFileName(cfg, m.Id))
		if err = os.Remove(saveFile); os.IsNotExist(err) {
			err = nil
		}
	} else {
		err = saveVhostConf(ctx, cfg, m.Id, values)
	}
	if err == nil && restart {
		item := engine.Engines.GetItem(cfg.Engine())
		if item != nil {
			err = item.X.(engine.Enginer).ReloadServer(ctx, cfg)
		}
	}
	return
}

func receiveFormData(ctx echo.Context, m *dbschema.NgingVhost) {
	m.Domain = ctx.Form(`domain`)
	m.Disabled = ctx.Form(`disabled`)
	m.Root = ctx.Form(`root`)
	m.Name = ctx.Form(`name`)
	m.GroupId = ctx.Formx(`groupId`).Uint()
	m.ServerIdent = ctx.Form(`serverIdent`)
}

func vhostbuild(ctx echo.Context, groupID uint, serverIdent string, engineType string, serverM ...*dbschema.NgingVhostServer) error {
	cond := db.NewCompounds()
	cond.AddKV(`a.disabled`, `N`)
	if groupID > 0 {
		cond.AddKV(`a.group_id`, groupID)
	}
	if len(serverM) > 0 {
		engineType = serverM[0].Engine
		serverIdent = serverM[0].Ident
	}
	var hasEngine, hasIdent bool
	if len(serverIdent) > 0 {
		hasIdent = true
		cond.AddKV(`a.server_ident`, serverIdent)
	} else if len(serverM) == 0 {
		if len(engineType) > 0 {
			if engineType == `default` {
				serverIdent = engineType
			} else {
				cond.AddKV(`b.engine`, engineType)
				hasEngine = true
			}
		}
		if len(serverIdent) > 0 {
			hasIdent = true
			cond.AddKV(`a.server_ident`, serverIdent)
		}
	}
	var err error
	configs := map[string]engine.Configer{}
	for _, v := range engine.Engines.Slice() {
		if len(engineType) > 0 && v.K != engineType {
			continue
		}
		eng := v.X.(engine.Enginer)
		if eng == nil {
			continue
		}
		if len(serverM) == 0 {
			var rows []engine.Configer
			rows, err = eng.ListConfig(ctx)
			if err != nil {
				return fmt.Errorf(`failed to ListConfig: %w`, err)
			}
			for _, cfg := range rows {
				if hasIdent && cfg.Ident() != serverIdent {
					continue
				}
				if groupID == 0 {
					var saveDir string
					saveDir, err = cfg.GetVhostConfigDirAbsPath()
					if err != nil {
						return fmt.Errorf(`failed to GetVhostConfigDirAbsPath: %w`, err)
					}
					err = removeAllConf(cfg, saveDir)
					if err != nil {
						return fmt.Errorf(`failed to removeAllConf: %w`, err)
					}
					os.Remove(saveDir)
				}
				configs[cfg.Ident()] = cfg
			}
		} else {
			cfg := eng.BuildConfig(ctx, serverM[0])
			if groupID == 0 {
				var saveDir string
				saveDir, err = cfg.GetVhostConfigDirAbsPath()
				if err != nil {
					return fmt.Errorf(`failed to GetVhostConfigDirAbsPath: %w`, err)
				}
				err = removeAllConf(cfg, saveDir)
				if err != nil {
					return fmt.Errorf(`failed to removeAllConf: %w`, err)
				}
				os.Remove(saveDir)
			}
			configs[cfg.Ident()] = cfg
		}
	}
	m := model.NewVhost(ctx)
	n := 100
	serverTable := dbschema.NewNgingVhostServer(ctx).Short_()
	var rowAndGroup []*model.VhostAndGroup
	var makeQuerier = func() *factory.Param {
		p := m.NewParam()
		if hasEngine {
			p.SetCols(`a.*`, `b.name AS serverName`, `b.engine AS serverEngine`).AddJoin(`LEFT`, serverTable, `b`, `b.ident=a.server_ident`)
		} else {
			p.SetCols(`a.*`)
		}
		return p.SetAlias(`a`).SetRecv(&rowAndGroup).AddArgs(cond.And())
	}
	var cnt func() int64
	cnt, err = makeQuerier().SetOffset(0).SetSize(n).List()
	if err != nil {
		return err
	}
	for i, j := 0, cnt(); int64(i) < j; i += n {
		if i > 0 {
			rowAndGroup = rowAndGroup[0:0]
			_, err = makeQuerier().SetOffset(i).SetSize(n).List()
			if err != nil {
				handler.SendFail(ctx, err.Error())
				return ctx.Redirect(handler.URLFor(`/caddy/vhost`))
			}
		}
		for _, m := range rowAndGroup {
			var formData url.Values
			err = json.Unmarshal([]byte(m.Setting), &formData)
			if err == nil {
				cfg, ok := configs[m.ServerIdent]
				if ok {
					err = saveVhostConf(ctx, cfg, m.Id, formData)
				} else {
					err = DeleteCaddyfileByID(ctx, m.ServerIdent, m.Id, serverM...)
				}
			}
			if err != nil {
				return err
			}
		}
	}
	for _, cfg := range configs {
		item := engine.Engines.GetItem(cfg.Engine())
		if item == nil {
			continue
		}
		err = item.X.(engine.Enginer).ReloadServer(ctx, cfg)
		if err != nil {
			ctx.Logger().Error(err)
		}
	}
	return nil
}

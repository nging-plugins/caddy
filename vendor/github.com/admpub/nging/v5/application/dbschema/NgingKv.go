// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingKv []*NgingKv

func (s Slice_NgingKv) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingKv) RangeRaw(fn func(m *NgingKv) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingKv) GroupBy(keyField string) map[string][]*NgingKv {
	r := map[string][]*NgingKv{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingKv{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingKv) KeyBy(keyField string) map[string]*NgingKv {
	r := map[string]*NgingKv{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingKv) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingKv) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingKv) FromList(data interface{}) Slice_NgingKv {
	values, ok := data.([]*NgingKv)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingKv{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingKv(ctx echo.Context) *NgingKv {
	m := &NgingKv{}
	m.SetContext(ctx)
	return m
}

// NgingKv 键值数据
type NgingKv struct {
	base    factory.Base
	objects []*NgingKv

	Id           uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Key          string `db:"key" bson:"key" comment:"键名" json:"key" xml:"key"`
	Value        string `db:"value" bson:"value" comment:"元素值" json:"value" xml:"value"`
	Description  string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Help         string `db:"help" bson:"help" comment:"帮助说明" json:"help" xml:"help"`
	Type         string `db:"type" bson:"type" comment:"类型标识" json:"type" xml:"type"`
	Sort         int    `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Updated      uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	ChildKeyType string `db:"child_key_type" bson:"child_key_type" comment:"子键类型(number/text...)" json:"child_key_type" xml:"child_key_type"`
}

// - base function

func (a *NgingKv) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingKv) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingKv) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingKv) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingKv) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingKv) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingKv) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingKv) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingKv) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingKv) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingKv) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingKv) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingKv) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingKv) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingKv) Objects() []*NgingKv {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingKv) XObjects() Slice_NgingKv {
	return Slice_NgingKv(a.Objects())
}

func (a *NgingKv) NewObjects() factory.Ranger {
	return &Slice_NgingKv{}
}

func (a *NgingKv) InitObjects() *[]*NgingKv {
	a.objects = []*NgingKv{}
	return &a.objects
}

func (a *NgingKv) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingKv) Short_() string {
	return "nging_kv"
}

func (a *NgingKv) Struct_() string {
	return "NgingKv"
}

func (a *NgingKv) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingKv) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingKv) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingKv) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingKv:
			err = DBI.FireReaded(a, queryParam, Slice_NgingKv(*v))
		case []*NgingKv:
			err = DBI.FireReaded(a, queryParam, Slice_NgingKv(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingKv) GroupBy(keyField string, inputRows ...[]*NgingKv) map[string][]*NgingKv {
	var rows Slice_NgingKv
	if len(inputRows) > 0 {
		rows = Slice_NgingKv(inputRows[0])
	} else {
		rows = Slice_NgingKv(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingKv) KeyBy(keyField string, inputRows ...[]*NgingKv) map[string]*NgingKv {
	var rows Slice_NgingKv
	if len(inputRows) > 0 {
		rows = Slice_NgingKv(inputRows[0])
	} else {
		rows = Slice_NgingKv(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingKv) AsKV(keyField string, valueField string, inputRows ...[]*NgingKv) param.Store {
	var rows Slice_NgingKv
	if len(inputRows) > 0 {
		rows = Slice_NgingKv(inputRows[0])
	} else {
		rows = Slice_NgingKv(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingKv) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingKv:
			err = DBI.FireReaded(a, queryParam, Slice_NgingKv(*v))
		case []*NgingKv:
			err = DBI.FireReaded(a, queryParam, Slice_NgingKv(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingKv) Insert() (pk interface{}, err error) {
	a.Id = 0
	if len(a.ChildKeyType) == 0 {
		a.ChildKeyType = "text"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingKv) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.ChildKeyType) == 0 {
		a.ChildKeyType = "text"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingKv) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.ChildKeyType) == 0 {
		a.ChildKeyType = "text"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingKv) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.ChildKeyType) == 0 {
		a.ChildKeyType = "text"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingKv) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.ChildKeyType) == 0 {
		a.ChildKeyType = "text"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingKv) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingKv) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingKv) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["child_key_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["child_key_type"] = "text"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingKv) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["child_key_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["child_key_type"] = "text"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Updatex()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(kvset).Updatex(); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", &m, editColumns, mw, args...)
	return
}

func (a *NgingKv) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingKv) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.ChildKeyType) == 0 {
			a.ChildKeyType = "text"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Id = 0
		if len(a.ChildKeyType) == 0 {
			a.ChildKeyType = "text"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingKv) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingKv) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingKv) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingKv) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingKv) Reset() *NgingKv {
	a.Id = 0
	a.Key = ``
	a.Value = ``
	a.Description = ``
	a.Help = ``
	a.Type = ``
	a.Sort = 0
	a.Updated = 0
	a.ChildKeyType = ``
	return a
}

func (a *NgingKv) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Key"] = a.Key
		r["Value"] = a.Value
		r["Description"] = a.Description
		r["Help"] = a.Help
		r["Type"] = a.Type
		r["Sort"] = a.Sort
		r["Updated"] = a.Updated
		r["ChildKeyType"] = a.ChildKeyType
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Key":
			r["Key"] = a.Key
		case "Value":
			r["Value"] = a.Value
		case "Description":
			r["Description"] = a.Description
		case "Help":
			r["Help"] = a.Help
		case "Type":
			r["Type"] = a.Type
		case "Sort":
			r["Sort"] = a.Sort
		case "Updated":
			r["Updated"] = a.Updated
		case "ChildKeyType":
			r["ChildKeyType"] = a.ChildKeyType
		}
	}
	return r
}

func (a *NgingKv) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "key":
			a.Key = param.AsString(value)
		case "value":
			a.Value = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "help":
			a.Help = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "sort":
			a.Sort = param.AsInt(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "child_key_type":
			a.ChildKeyType = param.AsString(value)
		}
	}
}

func (a *NgingKv) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Key":
			a.Key = param.AsString(vv)
		case "Value":
			a.Value = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Help":
			a.Help = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Sort":
			a.Sort = param.AsInt(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "ChildKeyType":
			a.ChildKeyType = param.AsString(vv)
		}
	}
}

func (a *NgingKv) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["key"] = a.Key
		r["value"] = a.Value
		r["description"] = a.Description
		r["help"] = a.Help
		r["type"] = a.Type
		r["sort"] = a.Sort
		r["updated"] = a.Updated
		r["child_key_type"] = a.ChildKeyType
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "key":
			r["key"] = a.Key
		case "value":
			r["value"] = a.Value
		case "description":
			r["description"] = a.Description
		case "help":
			r["help"] = a.Help
		case "type":
			r["type"] = a.Type
		case "sort":
			r["sort"] = a.Sort
		case "updated":
			r["updated"] = a.Updated
		case "child_key_type":
			r["child_key_type"] = a.ChildKeyType
		}
	}
	return r
}

func (a *NgingKv) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingKv) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingKv) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingKv) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}

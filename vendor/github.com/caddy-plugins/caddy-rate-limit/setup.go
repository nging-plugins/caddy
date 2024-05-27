package ratelimit

import (
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/admpub/caddy"
	"github.com/admpub/caddy/caddyhttp/httpserver"
)

func init() {

	caddy.RegisterPlugin("ratelimit", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	cfg := httpserver.GetConfig(c)

	rules, err := rateLimitParse(c)
	if err != nil {
		return err
	}

	// calculate whitelist IPNet in setup
	for _, rule := range rules {
		for _, s := range rule.Whitelist {
			_, ipNet, err := net.ParseCIDR(s)
			if err == nil {
				rule.whitelistIPNets = append(rule.whitelistIPNets, ipNet)
			}
		}
	}

	rateLimit := RateLimit{Rules: rules}
	cfg.AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		rateLimit.Next = next
		return rateLimit
	})

	return nil
}

func rateLimitParse(c *caddy.Controller) (rules []Rule, err error) {

	for c.Next() {
		var rule Rule

		args := c.RemainingArgs()
		switch len(args) {
		case 4:
			// config block
			rule.Methods = args[0]
			rule.Rate, err = strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return rules, err
			}
			rule.Burst, err = strconv.Atoi(args[2])
			if err != nil {
				return rules, err
			}
			rule.Unit = args[3]
		case 5:
			// one line config
			rule.Methods = args[0]
			rule.Resources = append(rule.Resources, args[1])
			rule.Rate, err = strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return rules, err
			}
			rule.Burst, err = strconv.Atoi(args[3])
			if err != nil {
				return rules, err
			}
			rule.Unit = args[4]
		default:
			return rules, c.ArgErr()
		}

		for c.NextBlock() {
			val := c.Val()
			args = c.RemainingArgs()
			switch len(args) {
			case 0:
				// resources
				rule.Resources = append(rule.Resources, val)
			case 1:
				switch val {
				case "whitelist":
					// check if CIDR is valid
					for _, v := range strings.Split(args[0], ",") {
						_, _, err := net.ParseCIDR(v)
						if err != nil {
							return rules, err
						}
						rule.Whitelist = append(rule.Whitelist, v)
					}
				case "limit_by_header":
					if len(args[0]) == 0 {
						return rules, c.Errf("invalid limit_by_header")
					}
					rule.LimitByHeader = http.CanonicalHeaderKey(args[0])
				case "status":
					// TODO: check status code is valid
					rule.Status = args[0]
				default:
					return rules, c.Errf("expecting whitelist, limit_by_header or status, got %s", val)
				}
			default:
				return rules, c.ArgErr()
			}
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

package controllers

import (
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/ohyo/revelmodules/wiki/app/lib/wikihelper"

	"github.com/revel/revel"
	"github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
)

var (
	// NewRelic agent
	agent *gorelic.Agent
)

func init() {
	// NewRelic
	revel.OnAppStart(initGorelic)
	revel.Filters = append(revel.Filters, gorelicFilter)

	// Call InitDB at server startup for automatic migration
	revel.OnAppStart(InitDB)

	// Start automatic transaction
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)

	// Define a template function that performs four arithmetic operations
	revel.TemplateFuncs["add"] = func(args ...int) int {
		result := 0
		for i, value := range args {
			if i == 0 {
				result = value
			} else {
				result += value
			}
		}
		return result
	}
	revel.TemplateFuncs["subtract"] = func(args ...int) int {
		result := 0
		for i, value := range args {
			if i == 0 {
				result = value
			} else {
				result -= value
			}
		}
		return result
	}
	revel.TemplateFuncs["multiply"] = func(args ...int) int {
		result := 0
		for i, value := range args {
			if i == 0 {
				result = value
			} else {
				result *= value
			}
		}
		return result
	}
	revel.TemplateFuncs["divide"] = func(args ...int) int {
		result := 0
		for i, value := range args {
			if i == 0 {
				result = value
			} else {
				result /= value
			}
		}
		return result
	}

	// Define a template function to format the date
	// Specify 2006-01-02 15:04:05 -0700 MST date and time for formatting
	revel.TemplateFuncs["date"] = func(format string, time time.Time) string {
		return time.Local().Format(format)
	}

	revel.TemplateFuncs["len_gt"] = func(arg interface{}, length int) bool {
		if reflect.TypeOf(arg).Kind() == reflect.Slice {
			return reflect.ValueOf(arg).Len() > length
		}
		return false
	}
	revel.TemplateFuncs["len_lt"] = func(arg interface{}, length int) bool {
		if reflect.TypeOf(arg).Kind() == reflect.Slice {
			return reflect.ValueOf(arg).Len() < length
		}
		return false
	}
	revel.TemplateFuncs["replace"] = func(s string, old string, new string) string {
		return strings.Replace(s, old, new, -1)
	}
	revel.TemplateFuncs["urlencode"] = func(str string) string {
		return wikihelper.UrlEncode(str)
	}
}

// initGorelic Initializes the Gorelic agent
func initGorelic() {

	// Load newrelic.license from app configuration file.
	// Only start gorelic if a license id resent.
	newrelicLicense := revel.Config.StringDefault("newrelic.license", "")
	if len(newrelicLicense) > 0 {
		log.Print("Starting newrelic daemon.")
		agent = gorelic.NewAgent()
		agent.NewrelicLicense = newrelicLicense
		agent.NewrelicName = "Wiki"
		agent.NewrelicPollInterval = 180
		agent.Verbose = true

		// "Manually" init the http timer (will be used in gorelicFilter)
		agent.CollectHTTPStat = true
		agent.HTTPTimer = metrics.NewTimer()

		agent.Run()
	} else {
		log.Print("!! Newrelic license missing from config file -> Not started")
	}
}

// Filter to capture HTTP metrics for gorelic
var gorelicFilter = func(c *revel.Controller, fc []revel.Filter) {
	startTime := time.Now()
	defer func() {
		if agent != nil {
			agent.HTTPTimer.UpdateSince(startTime)
		}
	}()
	fc[0](c, fc[1:])
}

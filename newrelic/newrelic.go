package newrelic

import (
	"crud-echo-postgres-redis/config"
	"github.com/joomcode/errorx"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"time"
)

var (
	app *newrelic.Application
	err error
)

func InitNewRelicApp() (*newrelic.Application, *errorx.Error) {
	env, _ := config.LoadConfig(".")
	app, err = newrelic.NewApplication(
		newrelic.ConfigAppName(env.NewRelicAppName),
		newrelic.ConfigLicense(env.NewRelicLicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigFromEnvironment(),
		func(config *newrelic.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)

	if err != nil {
		panic(err)
	}

	log.Print("Before sleep")
	err = app.WaitForConnection(5 * time.Second)
	if err != nil {
		panic(err)
	}
	log.Print("After sleep")

	go app.RecordCustomMetric("init_application_metric", float64(rand.Intn(1000)))
	go app.RecordCustomEvent("InitApplication", map[string]interface{}{
		"myString": "hello",
		"myFloat":  0.603,
		"myInt":    123,
		"myBool":   true,
	})

	if err != nil {
		return app, errorx.Decorate(err, "Failed to start New Relic integration")
	}

	return app, nil
}

func GetNewRelicApp() *newrelic.Application {
	return app
}

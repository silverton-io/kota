// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/silverton.io/kota/pkg/config"
	"github.com/silverton.io/kota/pkg/constants"
	"github.com/silverton.io/kota/pkg/handler"
	"github.com/spf13/viper"
)

var VERSION string

type App struct {
	config *config.Config
	engine *gin.Engine
	debug  bool
}

func is_debug_mode(debug string) bool {
	if debug != "" && (debug == "true" || debug == "1") {
		return true
	}
	return false
}

func (a *App) configure() {
	log.Info().Msg("configuring kota")
	// Configure application logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	gin.SetMode(gin.ReleaseMode)
	// Load application config from file
	conf_path := os.Getenv(constants.KOTA_CONFIG_PATH)
	if conf_path == "" {
		conf_path = constants.DEFAULT_CONFIG_PATH
	}
	log.Debug().Msg("loading config from file: " + conf_path)
	viper.SetConfigFile(conf_path)
	viper.SetConfigType("yaml") // FIXME: make const
	err := viper.ReadInConfig()
	if err != nil {
		// Blow up if kota cannot read config file.
		log.Fatal().Stack().Err(err).Msg("could not read config from file: " + conf_path)
	}
	a.config = &config.Config{}
	if err := viper.Unmarshal(a.config); err != nil {
		log.Fatal().Stack().Err(err).Msg("could not unmarshal config from file: " + conf_path)
	}

	// Flip into debug if env variable is set
	debug := os.Getenv(constants.DEBUG)
	if is_debug_mode(debug) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Warn().Msg("kota is set to debug mode")
		gin.SetMode(gin.DebugMode)
		a.config.Middleware.RequestLogger.Enabled = true
		a.debug = true
		log.Debug().Interface("config", a.config).Msg("kota config")
	}
	a.config.App.Version = VERSION

}

func (a *App) initializeRouter() {
	log.Info().Msg("initializing router")
	a.engine = gin.New()
	if err := a.engine.SetTrustedProxies(nil); err != nil {
		log.Fatal().Stack().Err(err).Msg("could not set trusted proxies")
	}
	if a.debug {
		log.Debug().Msg("setting up pprof at /debug/pprof")
		pprof.Register(a.engine)
	}
	a.engine.RedirectTrailingSlash = false
}

func (a *App) initializeRoutes() {
	log.Info().Msg("initializing routes")
	a.engine.GET(constants.DEFAULT_HEALTH_ROUTE, handler.HealthcheckHandler)
	a.engine.GET(constants.DEFAULT_OKTA_HOOKS_ROUTE, handler.OktaHookHandler)
}

func (a *App) initializeMiddleware() {
	log.Info().Msg("initializing middleware")
}

func (a *App) initializeConsumption() {
	// Consumers for EventBridge | Kinesis | Kafka go here.
	log.Info().Msg("initializing consumer")
}

func (a *App) initializeRedelivery() {
	log.Info().Msg("initializing redelivery mechanisms")
}

func (a *App) initialize() {
	log.Info().Msg("initializing kota")
	a.configure()
	// Initialize http collecter routes if configured to do so
	a.initializeRouter()
	a.initializeMiddleware()
	a.initializeRoutes()
	// Initialize consumer if configured to do so
	a.initializeConsumption()
	// Initialize redelivery mechanisms
	a.initializeRedelivery()
}

func (a *App) Run() {
	a.initialize()
	server := &http.Server{
		Addr:        ":" + a.config.App.Port,
		Handler:     a.engine,
		ReadTimeout: constants.DEFAULT_HTTP_TIMEOUT,
	}
	// Spin up http server
	go func() {
		log.Info().Msg("kota is running")
		if err := server.ListenAndServe(); err != nil {
			log.Debug().Msg("kota shut down successfully")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), constants.DEFAULT_SHUTDOWN_TIMEOUT)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Stack().Err(err).Msg("server forced to shutdown")
	}
}

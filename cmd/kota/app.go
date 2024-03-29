// Copyright (c) 2024 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/kota/blob/main/LICENSE

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var VERSION string

type App struct {
	engine *gin.Engine
	debug  bool
}

func (a *App) configure() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	gin.SetMode(gin.ReleaseMode)
	log.Debug().Msg("configuring kota")
}

func (a *App) Initialize() {
	log.Debug().Msg("initializing kota")
}

func (a *App) Run() {
	log.Debug().Interface("conf", "TODO").Msg("running kota with config")
	a.configure()
}

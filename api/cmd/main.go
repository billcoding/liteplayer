package main

import (
	"fmt"

	_ "github.com/billcoding/diskplayer/api/cmd/middleware"
	_ "github.com/billcoding/diskplayer/api/cmd/router"
	_ "github.com/billcoding/diskplayer/api/cmd/tasks"

	"github.com/billcoding/diskplayer/api/deps/envs"
	"github.com/rwscode/svc"
)

func main() { fmt.Println(svc.GetApp().Run(envs.ServerAddr)) }

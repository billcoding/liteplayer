package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	_ "github.com/billcoding/liteplayer/api/router"

	"github.com/billcoding/liteplayer/api/app"
	"github.com/billcoding/liteplayer/api/envs"
)

func main() { fmt.Println(app.GetApp().Run(envs.ServerAddr)) }

func init() {
	var (
		cc  int
		ccc int
	)
	filepath.Walk(envs.Root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			cc++
		} else {
			ccc++
		}
		return nil
	})
	fmt.Println("filepath walk end cc:", cc)
	fmt.Println("filepath walk end ccc:", ccc)
}

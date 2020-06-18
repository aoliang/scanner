package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aoliang/sanner/pkg/browser"
	"github.com/aoliang/sanner/pkg/module"
	"github.com/aoliang/sanner/pkg/pkgtree"
	"github.com/aoliang/sanner/pkg/visualization"
)

func main() {
	//获取待解析项目的根目录。
	fileRoot, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//获取待解析的module。
	importRoot := module.Get()

	//解析
	packageImports := pkgtree.ListPackages(fileRoot, importRoot)
	dotFile := "/tmp/importGraph.dot"
	pngFile := "/tmp/importGraph.png"

	visualization.PrintDag(packageImports, dotFile, pngFile)

	importsJson, _ := json.Marshal(packageImports)
	addr := "127.0.0.1:18888"
	go func() {
		time.Sleep(1 * time.Second)
		browser.Request("http://" + addr)
		browser.Request("file://" + pngFile)
	}()
	s := browser.StaticServer{Addr: addr, Data: importsJson}
	s.Serve()
}

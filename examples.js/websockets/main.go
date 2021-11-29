package main

import (
	"log"
	"path/filepath"

	"github.com/breakcount/go-sciter"
	"github.com/breakcount/go-sciter/window"
)

//websocketd 配合做服务器测试
func main() {

	DefaultRect := &sciter.Rect{0, 0, 1280, 800}
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, DefaultRect)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	w.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_FILE_IO | sciter.ALLOW_SOCKET_IO|sciter.ALLOW_EVAL | sciter.ALLOW_SYSINFO)

	fullpath, err := filepath.Abs("./examples.js/websockets/basic-test.htm")

	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)
	w.Show()
	w.Run()
}

package main

import (
	"log"
	"path/filepath"

	"github.com/breakcount/go-sciter"
	"github.com/breakcount/go-sciter/window"
)

//目前还不支持 sqlite
func main() {

	DefaultRect := &sciter.Rect{0, 0, 1280, 800}
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, DefaultRect)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	w.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_FILE_IO | sciter.ALLOW_SOCKET_IO|sciter.ALLOW_EVAL | sciter.ALLOW_SYSINFO)
	// SciterSetGlobalAsset(new sqlite::SQLite()); // adding SQLite as a global namespace object 还没有写
	fullpath, err := filepath.Abs("./examples.js/sqlite/db-basic.htm")

	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)
	w.Show()
	w.Run()
}

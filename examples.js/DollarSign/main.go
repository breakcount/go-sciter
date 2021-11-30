package main

import (
	"log"
	"path/filepath"

	"github.com/breakcount/go-sciter"
	"github.com/breakcount/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	fullpath, err := filepath.Abs("./examples.js/$/minimal-test.html") //fullpath, err := filepath.Abs("index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)

	w.Show()
	w.Run()
}

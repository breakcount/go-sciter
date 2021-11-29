package main

import (
	"log"
	"path/filepath"

	"github.com/breakcount/go-sciter"
	"github.com/breakcount/go-sciter/window"
)


func main() {

	DefaultRect := &sciter.Rect{0, 0, 1280, 800}
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, DefaultRect)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}

	//index := "clock.htm"
	index := "clock-hooks.html"
	//index := "hello-world.html"
	//index := "hello-world-jsx.html"
	//index := "hello-world-jsx-interactive.html"
	//index := "todo.htm"
	fullpath, err := filepath.Abs("./examples.js/preact/"+index)

	if err != nil {
		log.Fatal(err)
	}
	w.LoadFile(fullpath)
	w.Show()
	w.Run()
}

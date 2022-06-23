package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"time"
)

func main() {
	//var inTE, outTE *walk.TextEdit
	var w *walk.MainWindow
	stopchan := make(chan bool)
	MainWindow{
		AssignTo: &w,
		Title:    "moyu 1.1",
		Size:     Size{300, 50},
		Layout:   VBox{},
		Children: []Widget{
			/*HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},*/
			PushButton{
				Text: "mess around",
				OnClicked: func() {
					update(w, "mess around")
					go func() {
						for {
							x, y := robotgo.GetMousePos()
							robotgo.MoveMouse(x+1, y+1)
							robotgo.MoveMouse(x, y)
							time.Sleep(3 * time.Minute)
							select {
							case _ = <-stopchan:
								return
							default:
							}
						}
					}()
				},
			},
			PushButton{
				Text: "work",
				OnClicked: func() {
					update(w, "work")
					go func() {
						stopchan <- true
					}()
				},
			},
		},
	}.Run()
}

func update(w *walk.MainWindow, s string) {
	w.SetTitle(s)
}

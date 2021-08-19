package main

import (
	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"math/rand"
	"time"
)

func main() {
	//var inTE, outTE *walk.TextEdit
	var w *walk.MainWindow
	stopchan := make(chan bool)
	width, height := robotgo.GetScreenSize()
	MainWindow{
		AssignTo: &w,
		Title:    "工作中",
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
				Text:     "开始摸鱼",
				OnClicked: func() {
					update(w, "摸鱼中")
					go func() {
						for {
							x := rand.Intn(width)
							y := rand.Intn(height)
							robotgo.MoveMouse(x, y)
							time.Sleep(5 * time.Second)
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
				Text:     "开始工作",
				OnClicked: func() {
					update(w, "工作中")
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

/*width, height := robotgo.GetScreenSize()

x := rand.Intn(width)
y := rand.Intn(height) //生成0-99之间的随机数
// 将鼠标移动到屏幕 x:800 y:400 的位置（闪现到指定位置）
robotgo.MoveMouse(x, y)

//time.Sleep(10 * time.Second)
// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
//robotgo.MoveMouseSmooth(1000, 200)
time.Sleep(5 * time.Second)*/

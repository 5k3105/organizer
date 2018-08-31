package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var (
	ap   *Application
	font *gui.QFont
)

type Application struct {
	*widgets.QApplication
	Window    *widgets.QMainWindow
	Statusbar *widgets.QStatusBar
	Urlbar    *widgets.QLineEdit
}

func init() {
	font = gui.NewQFont2("corbel", 13, 1, false)
}

func main() {
	ap = &Application{}
	ap.QApplication = widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	ap.Window = window
	ap.Window.SetWindowTitle("Organizer")

	_, edittreelayout := New_EditTree(window)

	panel := widgets.NewQDockWidget("", ap.Window, core.Qt__Widget)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, panel)
	w := widgets.NewQWidget(ap.Window, core.Qt__Widget)
	w.SetLayout(edittreelayout)
	panel.SetWidget(w)

	// Run App
	widgets.QApplication_SetStyle2("fusion")
	window.ShowMaximized()
	widgets.QApplication_Exec()
}

package Ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type MainWindowApp struct {
	App fyne.App
	MainWindow fyne.Window
	MenuFile *fyne.Menu
}

func New (App fyne.App) MainWindowApp {
	ret := MainWindowApp{App: App}
	return ret
}
func (z *MainWindowApp) Run (file string) {
	z.MainWindow = z.App.NewWindow("RZip")
	z.MenuFile = fyne.NewMenu("File",&fyne.MenuItem{
		ChildMenu:   nil,
		IsSeparator: false,
		Label:       "打开",
		Action:      nil,
	})
	z.MainWindow.SetMainMenu(&fyne.MainMenu{Items: []*fyne.Menu{z.MenuFile}})
	z.MainWindow.SetContent(widget.NewHBox(
		widget.NewLabel("test"),
		))
	z.MainWindow.ShowAndRun()
}
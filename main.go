package main

import (
	"RZip/Ui"
	"fyne.io/fyne/app"
)
func main() {
	dealer := GetDealer("test.7z")
	dealer.GetNowFiles()
	return
	app := app.New()
	mainWindow := Ui.New(app)
	mainWindow.Run("")
}

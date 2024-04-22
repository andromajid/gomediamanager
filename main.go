package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/*fmt.Println("Hello, World!")
	var db = Db{}
	err := db.ConnectDB()
	if err != nil {
		fmt.Printf("%v", err)
	}
	ScanDirectory("/home/andro", db)*/
	gomediamanager := app.New()
	mediaManagerWindow := gomediamanager.NewWindow("Go Media Manager")
	mediaManagerWindow.SetContent(widget.NewLabel("Go Media Manager"))
	mediaManagerWindow.Show()
	gomediamanager.Run()

}

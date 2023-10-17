package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"gamch1k.org/todo-win/api"
)

func main() {
	tasks := api.GetTasks()

	mainApp := app.New()
	mainWindow := mainApp.NewWindow("List")

	topLabel := widget.NewLabel("Your To-Do List:")

	data := binding.BindStringList(
		&tasks,
	)

	list := widget.NewListWithData(data, 
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			}, 
			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			})
	
	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})

	mainWindow.SetContent(container.NewBorder(topLabel, add, nil, nil, list))
	mainWindow.ShowAndRun()
}
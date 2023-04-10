// package GUI

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// 	// "fyne.io/fyne/v2/container"
// 	// "fmt"
// 	"fyne.io/fyne/v2/canvas"	
// 	"strconv"
// 	// "fyne.io/fyne/v2/layout"
// )

// func ShowGUI() {
	// a := app.New()
	// w := a.NewWindow("Solving Shortest Path")
	// // resize
	// w.Resize(fyne.NewSize(1000, 1000))
	// // labeling
	// label := widget.NewLabel("Hello World!")
	// // btn
	// btn := widget.NewButton("A* Algorithm", func() {
	// 	fmt.Println("I am clicking the button")
	// })
	// // checkbox
	// check := widget.NewCheck("Optional", func(value bool) {
	// 	if value{
	// 		fmt.Println("I love Charence")
	// 	}else{
	// 		fmt.Println("I love Charence even more")
	// 	}
	// })
	// // container
	// vbox := container.NewVBox(
    //     label,
    //     btn,
	// 	check,
    // )
	// w.SetContent(vbox)
	// w.ShowAndRun()
package GUI

import (
    // "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func ShowGUI() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Graph Example")

    graph := widget.NewLabel("Graph widget placeholder") // Replace with a third-party plotting library

    button1 := widget.NewButton("Button 1", func() {
        // Handle button 1 click event
    })

    button2 := widget.NewButton("Button 2", func() {
        // Handle button 2 click event
    })

    buttons := container.New(layout.NewVBoxLayout(), button1, button2)
    // buttons.Alignment = fyne.TextAlignLeading

    content := container.New(layout.NewHBoxLayout(), graph, buttons)

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}

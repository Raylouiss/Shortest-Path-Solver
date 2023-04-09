package main

import (
	"fmt"
	"TUCIL3_13521054_13521143/src/Algorithm"
)

func main(){
	// include range to goal
	rangeToGoal := map[string]float64{
		"A": 10,
		"B": 8,
		"C": 7,
		"D": 3,
		"E": 2,
		"F": 0,
	}
	// include adj matrix
	adjMatrix := [][]float64{
		{0, 7, 0, 5, 0, 0},
		{7, 0, 8, 9, 7, 0},
		{0, 8, 0, 0, 5, 0},
		{5, 9, 0, 0, 15, 6},
		{0, 7, 5, 15, 0, 8},
		{0, 0, 0, 6, 8, 0},
	}
	nodeIdx := make(map[string]int)
	nodeIdx["A"] = 0
	nodeIdx["B"] = 1
	nodeIdx["C"] = 2
	nodeIdx["D"] = 3
	nodeIdx["E"] = 4
	nodeIdx["F"] = 5
	path, distance := Algorithm.AStar(rangeToGoal, adjMatrix, nodeIdx, "F", "A")
	fmt.Println("Cost :", distance)
	fmt.Println("This is the path:")
	for i:=0; i<len(path); i++{
		fmt.Print(path[i], " ")
	}
}


// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Hello World")

// 	w.SetContent(widget.NewLabel("Hello World!"))
// 	w.ShowAndRun()
// }
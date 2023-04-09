package main

import (
	"fmt"
	"TUCIL3_13521054_13521143/src/Algorithm"
)

func main(){
	// include range to goal
	rangeToGoal := make(map[string]float64)
	rangeToGoal["one"] = 5
	rangeToGoal["two"] = 2.236
	rangeToGoal["three"] = 3
	rangeToGoal["four"] = 2
	rangeToGoal["five"] = 0
	// include adj matrix
	adjMatrix := [][]float64{
		{0, 2.828, 4, 3.606, 0},
		{2.828, 0, 0, 1, 2.236},
		{4, 0, 0, 0, 3},
		{3.606, 1, 0, 0, 2},
		{0, 2.236, 3, 2, 0},
	}
	nodeIdx := make(map[string]int)
	nodeIdx["one"] = 0
	nodeIdx["two"] = 1
	nodeIdx["three"] = 2
	nodeIdx["four"] = 3
	nodeIdx["five"] = 4
	path := Algorithm.AStar(rangeToGoal, adjMatrix, nodeIdx, "five", "one")
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
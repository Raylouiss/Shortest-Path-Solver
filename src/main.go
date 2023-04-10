package main

import (
	"Tucil3_13521054_13521143/src/Algorithm"
	"Tucil3_13521054_13521143/src/Class"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	filepath := "../test/test.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Read totalNodes
	scanner.Scan()
	totalNodes, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Graph
	graph := Class.Graph{
		TotalNodes:      totalNodes,
		AdjacencyMatrix: make([][]float64, totalNodes),
		Nodes:           make([]Class.Node, totalNodes),
	}

	// Read nodes data
	for i := 0; i < totalNodes; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		edge := line[0]
		for j := 1; j < len(line)-2; j++ {
			edge = edge + " " + line[j]
		}
		latitude, _ := strconv.ParseFloat(line[len(line)-2], 64)
		longitude, _ := strconv.ParseFloat(line[len(line)-1], 64)
		graph.Nodes[i] = Class.Node{edge, latitude, longitude}
	}

	// Read adjacency matrix
	for i := 0; i < totalNodes; i++ {
		scanner.Scan()
		adjacencyData := strings.Split(scanner.Text(), "\t")

		graph.AdjacencyMatrix[i] = make([]float64, totalNodes)
		for j := 0; j < totalNodes; j++ {
			weight, err := strconv.ParseFloat(adjacencyData[j], 64)
			if err != nil {
				log.Fatal(err)
			}
			graph.AdjacencyMatrix[i][j] = weight
		}
	}
	// Print all nodes
	fmt.Println("Nodes:")
	Class.PrintAllNodes(graph.Nodes)
	fmt.Print("\n")
	fmt.Println("AdjacencyMatrix:")
	Class.PrintAdjacencyNodes(graph.AdjacencyMatrix)
	fmt.Print("\n")
	fmt.Println("Weight:")
	Class.PrintAllWeight(&graph)
	fmt.Print("\n")
	fmt.Println("WeightedAdjacencyMatrix:")
	Class.PrintAdjacencyNodes(Class.WeightedAdjacencyMatrix(graph.AdjacencyMatrix, &graph))

	// Get distance between two nodes
	fmt.Println("\nDistance from A to C:", graph.IsTetanggaFloat("A", "C"))
	fmt.Println("Distance from D to B:", graph.IsTetanggaFloat("D", "B"))

	// Get index of a node
	fmt.Println("\nIndex of node A:", graph.GetIndex("A"))
	fmt.Println("Index of node C:", graph.GetIndex("C"))

	fmt.Println(graph.GetDistance("A", "C"))
	fmt.Println(Algorithm.UCS("A", "F", graph.AdjacencyMatrix, &graph))

	rangeToGoal := map[string]float64{
		"A": graph.GetDistanceToGoal("A", "F"),
		"B": graph.GetDistanceToGoal("B", "F"),
		"C": graph.GetDistanceToGoal("C", "F"),
		"D": graph.GetDistanceToGoal("D", "F"),
		"E": graph.GetDistanceToGoal("E", "F"),
		"F": graph.GetDistanceToGoal("F", "F"),
		"G": graph.GetDistanceToGoal("G", "F"),
		"H": graph.GetDistanceToGoal("H", "F"),
	}
	// include adj matrix
	nodeIdx := make(map[string]int)
	for i := 0; i < graph.TotalNodes; i++ {
		nodeIdx[graph.Nodes[i].Name] = graph.GetIndex(graph.Nodes[i].Name)
	}
	// fmt.Println(nodeIdx)
	// adjMatrix := [][]float64{
	// 	{0, 7, 0, 5, 0, 0},
	// 	{7, 0, 8, 9, 7, 0},
	// 	{0, 8, 0, 0, 5, 0},
	// 	{5, 9, 0, 0, 15, 6},
	// 	{0, 7, 5, 15, 0, 8},
	// 	{0, 0, 0, 6, 8, 0},
	// }
	// fmt.Println(graph.AdjacencyMatrix)
	// fmt.Println(adjMatrix)
	adjMatrix2 := Class.WeightedAdjacencyMatrix(graph.AdjacencyMatrix, &graph)
	// fmt.Println(adjMatrix2)
	// nodeIdx["A"] = 0
	// nodeIdx["B"] = 1
	// nodeIdx["C"] = 2
	// nodeIdx["D"] = 3
	// nodeIdx["E"] = 4
	// nodeIdx["F"] = 5
	path, distance, found := Algorithm.AStar(rangeToGoal, adjMatrix2, nodeIdx, "F", "A")
	if(found){
		fmt.Println("Cost :", distance)
		fmt.Println("This is the path:")
		for i := 0; i < len(path); i++ {
			fmt.Print(path[i], " ")
		}
	}else{
		fmt.Println("STUCK")
	}
}

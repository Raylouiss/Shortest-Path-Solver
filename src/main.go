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
	filepath := "test/test.txt"
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
	fmt.Println(Algorithm.UCS("B", "G", graph.AdjacencyMatrix, &graph))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	totalNodes      int
	adjacencyMatrix [][]float64
	nodes           []node
}

type node struct {
	name       string
	latitude   float64
	longtitude float64
}

func (graph *Graph) getTotalNodes() int {
	return graph.totalNodes
}

func (graph *Graph) getNodes() []node {
	return graph.nodes
}

func (node *node) getLatitude() float64 {
	return node.latitude
}

func (node *node) getLongtitude() float64 {
	return node.longtitude
}

func (node *node) getName() string {
	return node.name
}

func (graph *Graph) isTetanggaFloat(from string, destination string) float64 {
	index1 := -1
	index2 := -1
	for i := 0; i < len(graph.nodes); i++ {
		if graph.nodes[i].name == from {
			index1 = i
		}
		if graph.nodes[i].name == destination {
			index2 = i
		}
	}
	if index1 == -1 || index2 == -1 {
		log.Fatal("Error : Node tidak ditemukan2")
	}
	return graph.adjacencyMatrix[index1][index2]
}

func (graph *Graph) getDistance(from string, destination string) float64 {
	isTetangga := graph.isTetanggaFloat(from, destination)
	distance := -1.0
	if isTetangga == 1 {
		index1 := graph.getIndex(from)
		index2 := graph.getIndex(destination)
		x := graph.nodes[index2].latitude - graph.nodes[index1].latitude
		y := graph.nodes[index2].longtitude - graph.nodes[index1].longtitude

		distance = math.Sqrt(x*x + y*y)
		return distance
	}
	log.Print("Node tidak bertertangga")
	return distance

}

func (graph *Graph) getIndex(nodeSearched string) int {
	for i := 0; i < len(graph.nodes); i++ {
		if graph.nodes[i].name == nodeSearched {
			return i
		}
	}
	log.Print("Error : node tidak ditemukan3")
	return 0
}

func printAllNodes(nodes []node) {
	count := 0
	for _, tempNode := range nodes {
		count++
		strCount := strconv.Itoa(count)
		fmt.Printf(strCount + ". " + tempNode.name)
		fmt.Printf("\n")
	}
}

func printAllWeight(graph *Graph) {
	for i := 0; i < graph.totalNodes; i++ {
		for j := 0; j < graph.totalNodes; j++ {
			if graph.isTetanggaFloat(graph.nodes[i].name, graph.nodes[j].name) == 1 {
				distance := graph.getDistance(graph.nodes[i].name, graph.nodes[j].name)
				distanceString := strconv.FormatFloat(distance, 'f', -1, 64)
				fmt.Println(graph.nodes[i].name, " - ", graph.nodes[j].name, " = ", distanceString)
			}
		}
	}
}

func printAdjacencyNodes(graph [][]float64) {
	for _, tempGraph := range graph {
		fmt.Println(tempGraph)
	}
}

func main() {
	// Read input from file
	file, err := os.Open("test.txt")
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
	graph := Graph{
		totalNodes:      totalNodes,
		adjacencyMatrix: make([][]float64, totalNodes),
		nodes:           make([]node, totalNodes),
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
		graph.nodes[i] = node{edge, latitude, longitude}
	}

	// Read adjacency matrix
	for i := 0; i < totalNodes; i++ {
		scanner.Scan()
		adjacencyData := strings.Split(scanner.Text(), "\t")

		graph.adjacencyMatrix[i] = make([]float64, totalNodes)
		for j := 0; j < totalNodes; j++ {
			weight, err := strconv.ParseFloat(adjacencyData[j], 64)
			if err != nil {
				log.Fatal(err)
			}
			graph.adjacencyMatrix[i][j] = weight
		}
	}
	// Print all nodes
	fmt.Println("Nodes:")
	printAllNodes(graph.nodes)
	fmt.Print("\n")
	fmt.Println("AdjacencyMatrix:")
	printAdjacencyNodes(graph.adjacencyMatrix)
	fmt.Print("\n")
	fmt.Println("Weight:")
	printAllWeight(&graph)

	// Get distance between two nodes
	fmt.Println("\nDistance from A to C:", graph.isTetanggaFloat("A", "C"))
	fmt.Println("Distance from D to B:", graph.isTetanggaFloat("D", "B"))

	// Get index of a node
	fmt.Println("\nIndex of node A:", graph.getIndex("A"))
	fmt.Println("Index of node C:", graph.getIndex("C"))

	fmt.Println(graph.getDistance("A", "C"))
}

package main

import (
	"TUCIL3_13521054_13521143/src/Algorithm"
	"TUCIL3_13521054_13521143/src/Class"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func read(filepath string, algo string, start string, goal string) ([]string, float64) {
	// Read input from file
	// filepath := "../test/test.txt"
	file, err := os.Open(filepath)
	fmt.Println(filepath)
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
		adjacencyData := strings.Split(scanner.Text(), " ")

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
	rangeToGoal := make(map[string]float64)
	keys := []string{}
	for i := 0; i < graph.TotalNodes; i++ {
		keys = append(keys, graph.Nodes[i].Name)
	}
	for _, key := range keys {
		distance := graph.GetDistanceToGoal(key, goal)
		rangeToGoal[key] = distance
	}
	// fmt.Println(rangeToGoal)
	// include adj matrix
	nodeIdx := make(map[string]int)
	for i := 0; i < graph.TotalNodes; i++ {
		nodeIdx[graph.Nodes[i].Name] = graph.GetIndex(graph.Nodes[i].Name)
	}
	adjMatrix2 := Class.WeightedAdjacencyMatrix(graph.AdjacencyMatrix, &graph)
	if algo == "Option 2" {
		path, distance := Algorithm.UCS(adjMatrix2, nodeIdx, goal, start)
		return path, distance
	} else {
		path, distance := Algorithm.AStar(rangeToGoal, adjMatrix2, nodeIdx, goal, start)
		return path, distance
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Go Algorithm is ready and listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue // Continue listening for new connections
		}

		// Handle connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process input from GUI
	input := string(buf[:n])
	result := processInput(input)

	// Send result back to GUI
	_, err = conn.Write([]byte(result))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Go Algorithm finished")
	fmt.Println(result)
}

func processInput(input string) string {
	inputSplit := strings.Split(input, "@")
	wordArray := []string{}
	for _, word := range inputSplit {
		wordArray = append(wordArray, word)
	}
	path, cost := read(wordArray[0], wordArray[1], wordArray[2], wordArray[3])
	result := ""
	for i := 0; i < len(path); i++ {
		result = strings.Join(path, " ")
	}
	result = result + " " + strconv.FormatFloat(cost, 'f', -1, 64) // This is just a dummy example
	return result
}

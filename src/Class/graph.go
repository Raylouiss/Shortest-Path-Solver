package Class

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type Graph struct {
	TotalNodes      int
	AdjacencyMatrix [][]float64
	Nodes           []Node
}

type Node struct {
	Name       string
	Latitude   float64
	Longtitude float64
}

func (graph *Graph) GetTotalNodes() int {
	return graph.TotalNodes
}

func (graph *Graph) GetNodes() []Node {
	return graph.Nodes
}

func (node *Node) GetLatitude() float64 {
	return node.Latitude
}

func (node *Node) GetLongtitude() float64 {
	return node.Longtitude
}

func (node *Node) GetName() string {
	return node.Name
}

func (graph *Graph) IsTetanggaFloat(from string, destination string) float64 {
	index1 := -1
	index2 := -1
	for i := 0; i < len(graph.Nodes); i++ {
		if graph.Nodes[i].Name == from {
			index1 = i
		}
		if graph.Nodes[i].Name == destination {
			index2 = i
		}
	}
	if index1 == -1 || index2 == -1 {
		log.Fatal("Error : Node tidak ditemukan2")
	}
	return graph.AdjacencyMatrix[index1][index2]
}

func (graph *Graph) GetDistance(from string, destination string) float64 {
	isTetangga := graph.IsTetanggaFloat(from, destination)
	distance := -1.0
	if isTetangga == 1 {
		index1 := graph.GetIndex(from)
		index2 := graph.GetIndex(destination)
		x := graph.Nodes[index2].Latitude - graph.Nodes[index1].Latitude
		y := graph.Nodes[index2].Longtitude - graph.Nodes[index1].Longtitude

		distance = math.Sqrt(x*x + y*y)
		return distance
	}
	log.Print("Node tidak bertertangga")
	return distance

}

func (graph *Graph) GetIndex(nodeSearched string) int {
	for i := 0; i < len(graph.Nodes); i++ {
		if graph.Nodes[i].Name == nodeSearched {
			return i
		}
	}
	log.Print("Error : node tidak ditemukan3")
	return 0
}

func PrintAllNodes(nodes []Node) {
	count := 0
	for _, tempNode := range nodes {
		count++
		strCount := strconv.Itoa(count)
		fmt.Printf(strCount + ". " + tempNode.Name)
		fmt.Printf("\n")
	}
}

func PrintAllWeight(graph *Graph) {
	for i := 0; i < graph.TotalNodes; i++ {
		for j := 0; j < graph.TotalNodes; j++ {
			if graph.IsTetanggaFloat(graph.Nodes[i].Name, graph.Nodes[j].Name) == 1 {
				distance := graph.GetDistance(graph.Nodes[i].Name, graph.Nodes[j].Name)
				distanceString := strconv.FormatFloat(distance, 'f', -1, 64)
				fmt.Println(graph.Nodes[i].Name, " - ", graph.Nodes[j].Name, " = ", distanceString)
			}
		}
	}
}

func PrintAdjacencyNodes(adjacencyMatrix [][]float64) {
	for _, tempMatrix := range adjacencyMatrix {
		fmt.Println(tempMatrix)
	}
}

func WeightedAdjacencyMatrix(adjacencyMatrix [][]float64, graph *Graph) [][]float64 {
	rows := len(adjacencyMatrix)
	cols := len(adjacencyMatrix[0])

	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if graph.IsTetanggaFloat(graph.Nodes[i].Name, graph.Nodes[j].Name) > 0 {
				distance := graph.GetDistance(graph.Nodes[i].Name, graph.Nodes[j].Name)
				result[i][j] = adjacencyMatrix[i][j] * distance
			} else {
				result[i][j] = 0
			}
		}
	}

	return result
}

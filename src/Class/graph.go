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

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func (graph *Graph) GetDistance(from string, destination string) float64 {
	isTetangga := graph.IsTetanggaFloat(from, destination)
	distance := -1.0
	earthRadius := 6371.0
	if isTetangga == 1 {
		index1 := graph.GetIndex(from)
		index2 := graph.GetIndex(destination)

		lat1Rad := degToRad(graph.Nodes[index1].Latitude)
		lat2Rad := degToRad(graph.Nodes[index2].Latitude)
		lon1Rad := degToRad(graph.Nodes[index1].Longtitude)
		lon2Rad := degToRad(graph.Nodes[index2].Longtitude)

		deltaLat := lat2Rad - lat1Rad
		deltaLon := lon2Rad - lon1Rad

		a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
			math.Cos(lat1Rad)*math.Cos(lat2Rad)*
				math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

		// Calculate the angular distance in radians
		c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

		// Calculate the distance in kilometers
		distance := earthRadius * c
		return distance
	}
	log.Print("Node tidak bertertangga")
	return distance

}

func (graph *Graph) GetDistanceToGoal(from string, destination string) float64 {
	distance := -1.0
	earthRadius := 6371.0

	index1 := graph.GetIndex(from)
	index2 := graph.GetIndex(destination)
	lat1Rad := degToRad(graph.Nodes[index1].Latitude)
	lat2Rad := degToRad(graph.Nodes[index2].Latitude)
	lon1Rad := degToRad(graph.Nodes[index1].Longtitude)
	lon2Rad := degToRad(graph.Nodes[index2].Longtitude)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	// Calculate the angular distance in radians
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance in kilometers
	distance = earthRadius * c
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

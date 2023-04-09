package Algorithm

import (
	"Tucil3_13521054_13521143/src/Class"
)

func SearchChild(strings string, graph *Class.Graph, visited map[string]bool) []Class.Node {
	index := graph.GetIndex(strings)
	result := []Class.Node{}
	for i := 0; i < graph.TotalNodes; i++ {
		if !IsVisited(visited, graph.Nodes[i].Name) {
			if graph.AdjacencyMatrix[index][i] > 0 {
				result = append(result, graph.Nodes[i])
			}
		}
	}
	return result
}

func GetSmallest(frontier map[string]float64, graph *Class.Graph) string {
	value := 10000000.0
	smallestNode := "aa"
	for i := 0; i < graph.TotalNodes; i++ {
		currentNode := graph.Nodes[i].Name
		_, found := frontier[currentNode]
		if found {
			if value > frontier[currentNode] {
				value = frontier[currentNode]
				smallestNode = currentNode
			}
		}
	}
	return smallestNode
}

func IsGoal(node string, goal string) bool {
	return node == goal
}

func IsEmpty(frontier map[string]float64) bool {
	return len(frontier) == 0
}

func AddFrontier(frontier map[string]float64, distance float64, key string) map[string]float64 {
	newFrontier := make(map[string]float64)
	for k, v := range frontier {
		newFrontier[k] = v
	}
	newFrontier[key] = distance
	return newFrontier
}

func AddVisited(currentNode string, visited map[string]bool) map[string]bool {
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	newVisited[currentNode] = true

	return newVisited
}

func IsVisited(visited map[string]bool, string string) bool {
	_, ok := visited[string]
	return ok
}

func ContainsNode(nodes []Class.Node, target Class.Node) bool {
	for _, n := range nodes {
		if n == target {
			return true
		}
	}
	return false
}

func ReverseStringArray(arr []string) {
	left := 0
	right := len(arr) - 1

	// Swap elements from left to right
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func CreatePath(parent map[string]string, currentNode string) []string {

	newPath := []string{}
	newPath = append(newPath, currentNode)

	for {
		if parent[currentNode] == "0" {
			break
		}
		currentNode = parent[currentNode]
		newPath = append(newPath, currentNode)
	}
	ReverseStringArray(newPath)

	return newPath
}

func UCS(start string, goal string, adjMatrix [][]float64, graph *Class.Graph) ([]string, float64) {
	visited := make(map[string]bool)
	frontier := make(map[string]float64)
	path := []string{}
	currentNode := start
	// tempNode := node{name: start}
	frontier = AddFrontier(frontier, 0, currentNode)
	weight := make(map[string]float64)
	parent := make(map[string]string)
	for i := 0; i < graph.TotalNodes; i++ {
		weight[graph.Nodes[i].Name] = 0
		parent[graph.Nodes[i].Name] = "0"
	}

	for !IsGoal(currentNode, goal) && !IsEmpty(frontier) && !IsVisited(visited, currentNode) {
		visited = AddVisited(currentNode, visited)
		nodeChild := SearchChild(currentNode, graph, visited)

		for i := 0; i < graph.TotalNodes; i++ {
			destination := graph.Nodes[i].Name
			if graph.IsTetanggaFloat(currentNode, destination) > 0 {
				if ContainsNode(nodeChild, graph.Nodes[i]) {
					parent[graph.Nodes[i].Name] = currentNode
					weights := graph.GetDistance(currentNode, graph.Nodes[i].Name)
					weight[graph.Nodes[i].Name] = weights + weight[parent[graph.Nodes[i].Name]]
					totalWeight := weight[graph.Nodes[i].Name]
					frontier = AddFrontier(frontier, totalWeight, graph.Nodes[i].Name)
				}
			}
		}
		delete(frontier, currentNode)
		currentNode = GetSmallest(frontier, graph)
	}
	totalWeight := weight[currentNode]
	path = CreatePath(parent, currentNode)
	return path, totalWeight
}

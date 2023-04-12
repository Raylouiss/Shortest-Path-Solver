package Algorithm

import (
	"container/heap"
	"fmt"
)

func UCS(adjMatrix [][]float64, nodeIndex map[string]int, goal string, start string) ([]string, float64) {
    fmt.Println("You are using UCS Algorithm")
	pq := make(PriorityQueue, 1)

	tempAdjMatrix := CopyMatrix(adjMatrix)

	distance := 0.0

	path := []string{start}

	visited := make(map[string]bool)
	visited[start] = true

	itemValue := []interface{}{
		path,
		start,
		distance,
	}
	pq[0] = &Item{
		value:    itemValue,
		priority: distance,
		index:    0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Item)
		currentValue := currentNode.value.([]interface{})
		nodeName := currentValue[1].(string)
		visited[nodeName] = true
		currentIdx := nodeIndex[nodeName]
		currentPath := currentValue[0].([]string)
		distanceCost := currentValue[2].(float64)
		if nodeName == goal {
			path = currentPath
			distance = distanceCost
			break
		}

		for idx, val := range tempAdjMatrix[currentIdx] {
			if val != 0 {
				nextNode, found := GetKeyByValue(nodeIndex, idx)
				if found && !visited[nextNode] {
					tempCurrentPath := append([]string{}, currentPath...) // create a copy of currentPath
					tempCurrentPath = append(tempCurrentPath, nextNode)
					tempDistanceCost := distanceCost + val
					tempItemValue := []interface{}{
						tempCurrentPath,
						nextNode,
						tempDistanceCost,
					}
					tempItem := &Item{
						value:    tempItemValue,
						priority: tempDistanceCost,
					}
					heap.Push(&pq, tempItem)
				}
			}
		}
	}

	return path, distance
}

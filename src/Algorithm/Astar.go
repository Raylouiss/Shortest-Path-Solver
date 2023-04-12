package Algorithm

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    interface{}
	priority float64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func CopyMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])
	matrixCopy := make([][]float64, rows)
	for i := range matrixCopy {
		matrixCopy[i] = make([]float64, cols)
		copy(matrixCopy[i], matrix[i])
	}
	return matrixCopy
}

func GetKeyByValue(myMap map[string]int, value int) (string, bool) {
	for key, val := range myMap {
		if val == value {
			return key, true
		}
	}
	return "", false
}

func TurnOffNode(adjMatrix [][]float64, firstNode int, secondNode int) {
	adjMatrix[secondNode][firstNode] = 0
	for i := 0; i < len(adjMatrix[firstNode]); i++ {
		adjMatrix[firstNode][i] = 0
	}
	for i := 0; i < len(adjMatrix); i++ {
		adjMatrix[i][secondNode] = 0
	}
}

func AStar(rangeToGoal map[string]float64, adjMatrix [][]float64, nodeIndex map[string]int, goal string, start string) ([]string, float64) {
	fmt.Println("You are using A* Algorithm")
	
	pq := make(PriorityQueue, 1)

	tempAdjMatrix := CopyMatrix(adjMatrix)

	distance := 0.0

	fValue := distance + rangeToGoal[start]

	path := []string{start}

	itemValue := []interface{}{
		path,
		start,
		distance,
	}
	pq[0] = &Item{
		value:    itemValue,
		priority: fValue,
		index:    0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Item)
		currentValue := currentNode.value.([]interface{})
		nodeName := currentValue[1].(string)
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
				nextNode, _ := GetKeyByValue(nodeIndex, idx)
				tempCurrentPath := append([]string{}, currentPath...)
				tempCurrentPath = append(tempCurrentPath, nextNode)
				tempDistanceCost := distanceCost + val
				tempFValue := tempDistanceCost + rangeToGoal[nextNode]
				tempItemValue := []interface{}{
					tempCurrentPath,
					nextNode,
					tempDistanceCost,
				}
				tempItem := &Item{
					value:    tempItemValue,
					priority: tempFValue,
				}
				heap.Push(&pq, tempItem)
			}
		}
	}

	return path, distance
}

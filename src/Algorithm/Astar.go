package Algorithm

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority float64     // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
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
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
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

func checkStillHaveNode(adjMatrix [][]float64) bool {
	for i := 0; i < len(adjMatrix); i++ {
		for j := 0; j < len(adjMatrix[i]); j++ {
			if adjMatrix[i][j] != 0 {
				return true
			}
		}
	}
	return false
}

func FindNextNode(currentNodeDirection []float64, rangeToGoal map[string]float64, distance *float64, nodeIndex map[string]int) int {
	minIdx, fValue := -1, 0.0
	count := 0
	tempDistance := 0.0
	for idx, value := range currentNodeDirection {
		tempFValue := 0.0
		if value != 0 {
			gValue := *distance + value
			nodeName, found := GetKeyByValue(nodeIndex, idx)
			hValue := 0.0
			if found {
				hValue = rangeToGoal[nodeName]
			}
			tempFValue = gValue + hValue
			if count == 0 {
				minIdx, fValue = idx, tempFValue
				tempDistance = gValue
			} else {
				if fValue > tempFValue {
					minIdx, fValue = idx, tempFValue
					tempDistance = gValue
				}
			}
			count++
		}
	}
	*distance = tempDistance
	return minIdx
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
	pq := make(PriorityQueue, 1)

	tempAdjMatrix := CopyMatrix(adjMatrix)

	distance := 0.0

	fValue := distance + rangeToGoal[start]

	path := []string{start}

	// visited := make(map[string]bool)
	// visited[start] = true

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
	// visited[nodeName] = true
	currentIdx := nodeIndex[nodeName]
	currentPath := currentValue[0].([]string)
	distanceCost := currentValue[2].(float64)
	fmt.Println("This is the path now", currentPath)
	if nodeName == goal {
	path = currentPath
	distance = distanceCost
	break
	}

	for idx, val := range tempAdjMatrix[currentIdx] {
	if val != 0 {
		nextNode, _ := GetKeyByValue(nodeIndex, idx)
		// if found && !visited[nextNode] {
		tempCurrentPath := append([]string{}, currentPath...) // create a copy of currentPath
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

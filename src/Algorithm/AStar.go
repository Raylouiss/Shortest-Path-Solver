package Algorithm

import "fmt"

func CopyMatrix(matrix [][]float64) [][]float64{
	rows := len(matrix)
	cols := len(matrix[0])
	matrixCopy := make([][]float64, rows)
	for i := range matrixCopy {
		matrixCopy[i] = make([]float64, cols)
		matrixCopy[i] = append(matrixCopy[i], matrix[i]...)
	}
	return matrixCopy
}

func GetKeyByValue(myMap map[string]int, value int) (string, bool){
	for key,val := range myMap{
		if(val == value){
			return key, true
		}
	}
	return "", false
}

// func GetAllNode(nodeIndex map[string]int) map[string]bool{
// 	nodes := make(map[string]bool)
// 	for key, val := range nodeIndex{
// 		nodes[key] = false
// 	}
// 	return nodes
// }

// func CheckAllNodeStatus(nodeStatus map[string]bool) bool{
// 	for key, val := range nodeStatus{
// 		if(!val){
// 			return false
// 		}
// 	}
// 	return true
// }

func checkStillHaveNode(adjMatrix [][]float64) bool {
	for i := 0; i<len(adjMatrix); i++{
		for j := 0; j<len(adjMatrix[i]); j++ {
			if(adjMatrix[i][j] != 0){
				return true
			}
		}
	}
	return false
}

func FindNextNode(currentNodeDirection []float64, rangeToGoal map[string]float64, distance *float64, nodeIndex map[string]int) int {
	minIdx, fValue := 0, 0.0
	count := 0
	tempDistance := 0.0
	for idx, value := range currentNodeDirection {
		tempFValue := 0.0
		if(value != 0) {
			gValue := *distance + value
			nodeName, found := GetKeyByValue(nodeIndex, idx)
			hValue := 0.0
			if(found){
				hValue = rangeToGoal[nodeName]
			}
			tempFValue = gValue + hValue
			if(count == 0){
				minIdx, fValue = idx, tempFValue 
				tempDistance = gValue
			}else{
				if(fValue > tempFValue){
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

func TurnOffNode(adjMatrix [][]float64, firstNode int, secondNode int){
	adjMatrix[secondNode][firstNode] = 0
	for i:=0; i<len(adjMatrix[firstNode]); i++ {
		adjMatrix[firstNode][i] = 0;
	}
	for i:=0; i<len(adjMatrix); i++ {
		adjMatrix[i][secondNode] = 0
	}
}

func AStar(rangeToGoal map[string]float64, adjMatrix [][]float64, nodeIndex map[string]int, goal string, start string) []string {
	tempAdjMatrix := CopyMatrix(adjMatrix)
	tempIdx := nodeIndex[start]
	currentNode := start
	distance := 0.0
	path := []string{start}
	fmt.Println("Im here")
	for currentNode != goal && checkStillHaveNode(tempAdjMatrix) {
		fmt.Println("This is current node :", currentNode)
		currentNodeDirection := adjMatrix[tempIdx]
		nextIdx := FindNextNode(currentNodeDirection, rangeToGoal, &distance, nodeIndex)
		TurnOffNode(tempAdjMatrix, tempIdx, nextIdx)
		nodeName, _ := GetKeyByValue(nodeIndex, nextIdx)
		path = append(path, nodeName)
		currentNode = nodeName
		tempIdx = nextIdx
		fmt.Println("DONE")
	}
	return path
}
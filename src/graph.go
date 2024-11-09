package src

import (
	"fmt"
	"slices"
)

// Тип для функции получения друзей
type GetFriendsIDsFunc func(userID int) ([]int, error)

// backtrace восстанавливает путь от userIDa до userIDb
func backtrace(path map[int]int, start, end int) []int {
	if _, exists := path[end]; !exists {
		fmt.Printf("Error: No path to %d found in backtrace\n", end)
		return nil
	}

	shortest := []int{end}
	for shortest[len(shortest)-1] != start {
		prevNode, ok := path[shortest[len(shortest)-1]]
		if !ok || prevNode == shortest[len(shortest)-1] {
			fmt.Printf("Error: Node %d not found in path or stuck in loop\n", shortest[len(shortest)-1])
			return nil
		}
		shortest = append(shortest, prevNode)
	}
	slices.Reverse(shortest)
	return shortest
}

// bidirectionalSearch ищет кратчайший путь между двумя пользователями с помощью двустороннего BFS
func bidirectionalSearch(
	userIDa, userIDb int,
	getFriendIDs GetFriendsIDsFunc,
) ([]int, error) {
	if userIDa == userIDb {
		return []int{userIDa}, nil
	}

	// Очереди для прямого и обратного поиска
	qA := []int{userIDa}
	qB := []int{userIDb}

	// Посещённые узлы
	visitedA := map[int]bool{userIDa: true}
	visitedB := map[int]bool{userIDb: true}

	// Путь для восстановления маршрута
	pathA := map[int]int{userIDa: userIDa}
	pathB := map[int]int{userIDb: userIDb}

	for len(qA) > 0 && len(qB) > 0 {
		// Выполняем шаги поиска из обеих сторон
		if found, path := bfsStep(&qA, visitedA, visitedB, pathA, pathB, getFriendIDs); found {
			return path, nil
		}

		if found, path := bfsStep(&qB, visitedB, visitedA, pathB, pathA, getFriendIDs); found {
			return path, nil
		}
	}

	// Путь не найден
	return nil, nil
}

// bfsStep выполняет один шаг BFS
func bfsStep(
	queue *[]int,
	visited, otherVisited map[int]bool,
	path, otherPath map[int]int,
	getFriendIDs GetFriendsIDsFunc,
) (bool, []int) {
	if len(*queue) == 0 {
		return false, nil
	}

	currentID := (*queue)[0]
	*queue = (*queue)[1:]

	//fmt.Printf("Current node: %d\n", currentID)

	friends, err := getFriendIDs(currentID)
	if err != nil {
		//fmt.Printf("Error fetching friends for %d: %v\n", currentID, err)
		return false, nil
	}

	//fmt.Printf("Friends of %d: %v\n", currentID, friends)

	for _, friendID := range friends {
		if visited[friendID] {
			continue
		}
		visited[friendID] = true
		path[friendID] = currentID

		// Если найден узел, посещённый другой стороной
		if otherVisited[friendID] {
			//fmt.Printf("Meeting point found: %d\n", friendID)
			return true, mergePaths(path, otherPath, friendID)
		}

		*queue = append(*queue, friendID)
	}

	return false, nil
}

// mergePaths объединяет пути из прямого и обратного поиска
func mergePaths(pathA, pathB map[int]int, meetingPoint int) []int {
	//fmt.Printf("Merging paths at meeting point: %d\n", meetingPoint)

	// Восстанавливаем путь от userIDa до meetingPoint
	path1 := backtrace(pathA, pathA[meetingPoint], meetingPoint)
	if path1 == nil {
		fmt.Println("Error: Path1 is nil")
		return nil
	}

	// Восстанавливаем путь от meetingPoint до userIDb
	path2 := backtrace(pathB, pathB[meetingPoint], meetingPoint)
	if path2 == nil {
		fmt.Println("Error: Path2 is nil")
		return nil
	}

	// Исключаем дублирование meetingPoint при объединении
	path2 = path2[:len(path2)-1]
	slices.Reverse(path1)
	fullPath := append(path2, path1...)
	fmt.Printf("Merged path: %v\n", fullPath)
	return fullPath
}

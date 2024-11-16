package src

import (
	"fmt"
	"slices"
)

// Тип для функции получения друзей
type GetFriendsIDsFunc func(userID int) ([]int, error)

// backtrace восстанавливает путь от userIDa до userIDb
func backtrace(path map[int]int, start, end int) []int {
	shortest := []int{end}
	for end != start {
		prevNode, ok := path[end]
		if !ok {
			return nil // Если узел не найден, возвращаем nil
		}
		end = prevNode
		shortest = append(shortest, end)
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
		if found, path := bfsStep(&qA, visitedA, visitedB, pathA, pathB, getFriendIDs, userIDa, userIDb); found {
			return path, nil
		}

		if found, path := bfsStep(&qB, visitedB, visitedA, pathB, pathA, getFriendIDs, userIDb, userIDa); found {
			return path, nil
		}
	}

	return nil, fmt.Errorf("Путь не найден")
}

func bfsStep(
	queue *[]int,
	visited, otherVisited map[int]bool,
	path, otherPath map[int]int,
	getFriendIDs GetFriendsIDsFunc,
	userIDa, userIDb int,
) (bool, []int) {
	if len(*queue) == 0 {
		return false, nil
	}

	currentID := (*queue)[0]
	*queue = (*queue)[1:]

	friends, err := getFriendIDs(currentID)
	if err != nil {
		return false, nil
	}

	for _, friendID := range friends {
		if visited[friendID] {
			continue
		}
		visited[friendID] = true
		path[friendID] = currentID

		if otherVisited[friendID] {
			return true, mergePaths(path, otherPath, friendID, userIDa, userIDb)
		}

		*queue = append(*queue, friendID)
	}

	return false, nil
}

// mergePaths объединяет пути из прямого и обратного поиска
func mergePaths(pathA, pathB map[int]int, meetingPoint int, userIDa, userIDb int) []int {
	// Восстанавливаем путь от userIDa до meetingPoint
	path1 := backtrace(pathA, userIDa, meetingPoint)
	if path1 == nil {
		fmt.Println("Error: Path1 is nil")
		return nil
	}

	// Восстанавливаем путь от meetingPoint до userIDb
	path2 := backtrace(pathB, userIDb, meetingPoint)
	if path2 == nil {
		fmt.Println("Error: Path2 is nil")
		return nil
	}

	// Убираем meetingPoint из второго пути
	if len(path2) > 0 {
		path2 = path2[:len(path2)-1]
	}

	// Соединяем оба пути
	slices.Reverse(path2)
	fullPath := append(path1, path2...)
	return fullPath
}

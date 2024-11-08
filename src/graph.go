package src

import "slices"

type GetFriendsIDsFunc func(userID int) ([]int, error)
type GetUsersDetailsFunc func(userIDs []int) ([]Friend, error)

func backtrace(path map[int]int, userIDa, userIDb int) []int {
	shortest := []int{userIDb}
	for shortest[len(shortest)-1] != userIDa {
		shortest = append(shortest, path[shortest[len(shortest)-1]])
	}
	slices.Reverse(shortest)
	return shortest
}

func BuildGraph(userIDa,
	userIDb int,
	getFriendIDs GetFriendsIDsFunc,
	getUsersDetails GetUsersDetailsFunc) (map[int]map[int]Friend, []int, error) {

	used := make(map[int]bool)
	q := []int{userIDa}
	graph := make(map[int]map[int]Friend)
	path := make(map[int]int)
	path[userIDa] = -1

	for !used[userIDb] {
		if len(q) == 0 {
			break
		}

		topFriendID := q[0]
		q = q[1:]
		if used[topFriendID] {
			continue
		}
		used[topFriendID] = true

		friendsOfTopFriendIDs, err := getFriendIDs(topFriendID)
		if err != nil {
			continue
		}

		friendsOfTopFriendInfo, err := getUsersDetails(friendsOfTopFriendIDs)
		if err != nil {
			continue
		}

		if graph[topFriendID] == nil {
			graph[topFriendID] = make(map[int]Friend)
		}

		for _, friend := range friendsOfTopFriendInfo {
			graph[topFriendID][friend.ID] = friend
			if !used[friend.ID] {
				path[friend.ID] = topFriendID
			}
			if friend.ID == userIDb {
				break
			}
		}

		q = append(q, friendsOfTopFriendIDs...)
	}

	return graph, backtrace(path, userIDa, userIDb), nil
}

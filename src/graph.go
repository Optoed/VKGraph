package src

type GetFriendsIDsFunc func(userID int) ([]int, error)
type GetUsersDetailsFunc func(userIDs []int) ([]Friend, error)

func BuildGraph(userIDa, userIDb int, getFriendIDs GetFriendsIDsFunc, getUsersDetails GetUsersDetailsFunc) (map[int]map[int]Friend, error) {
	used := make(map[int]bool)
	q := []int{userIDa}
	graph := make(map[int]map[int]Friend)

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
			return nil, err
		}

		friendsOfTopFriendInfo, err := getUsersDetails(friendsOfTopFriendIDs)
		if err != nil {
			return nil, err
		}

		if graph[topFriendID] == nil {
			graph[topFriendID] = make(map[int]Friend)
		}

		for _, friend := range friendsOfTopFriendInfo {
			graph[topFriendID][friend.ID] = friend
		}

		q = append(q, friendsOfTopFriendIDs...)
	}

	return graph, nil
}

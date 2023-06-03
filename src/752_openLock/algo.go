package algo

func openLock(deadends []string, target string) int {
	type Node struct {
		Number string
		Move   int
	}
	visited := map[string]bool{}
	q := []Node{{
		Number: "0000",
		Move:   0,
	}}

	deadendMap := map[string]bool{}
	for _, deadend := range deadends {
		deadendMap[deadend] = true
	}
	if len(deadends) >= 8 && checkDeadlock(deadendMap, target) {
		return -1
	}

	for len(q) > 0 {
		// TODO: 什麼時候要+ move
		cursor := q[0]
		q = q[1:]
		visited[cursor.Number] = true

		if _, ok := deadendMap[cursor.Number]; ok {
			continue
		}
		if cursor.Number == target {
			return cursor.Move
		}

		neighbors := findNeighbors(cursor.Number)
		for _, neighbor := range neighbors {
			q = append(q, Node{Number: neighbor, Move: cursor.Move + 1})
		}
	}
	return -1
}

func findNeighbors(cursor string) []string {
	neighbors := []string{}
	for i := 0; i < len(cursor); i++ {
		c := int(cursor[i])
		if c == 48 {
			neighbors = append(neighbors, cursor[0:i]+"1"+cursor[i+1:])
			neighbors = append(neighbors, cursor[0:i]+"9"+cursor[i+1:])
		} else {
			neighbors = append(neighbors, cursor[0:i]+string(rune(c+1))+cursor[i+1:])
			neighbors = append(neighbors, cursor[0:i]+string(rune(c-1))+cursor[i+1:])
		}
	}
	return neighbors
}

func checkDeadlock(deadendMap map[string]bool, target string) bool {

	neighbors := findNeighbors(target)
	for _, neighbor := range neighbors {
		if _, ok := deadendMap[neighbor]; !ok {
			return false
		}
	}
	return true
}

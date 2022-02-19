package algo

func firstBadVersion(n int) int {
	start :=1
	end := n
	for start+1<end{
		middle := (start+end)/2
		if isBadVersion(middle){
			end = middle
		} else{
			start = middle
		}
	}
	if isBadVersion(start){
		return start
	} else{
		return end
	}
}
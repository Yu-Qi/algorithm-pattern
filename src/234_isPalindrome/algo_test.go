package algo
import (
	. "../ListNode"
	"testing"
)
func TestIsPalindrome(t *testing.T){
	node1 := ArrToListNode([]int{1, 2, 2, 1})
	if isPalindrome(node1) == true{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestIsPalindrome2(t *testing.T){
	node1 := ArrToListNode([]int{1, 2, 3, 2, 1})
	if isPalindrome(node1) == true{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

package algo
import (
	"testing"
	. "../ListNode"
	"reflect"
)
		

func TestDeleteDuplicates1(t *testing.T){
	nums := []int{1,1,2,3,3}
	listnode := ArrToListNode(nums)
	expected :=ArrToListNode([]int{1,2,3})
	if reflect.DeepEqual(deleteDuplicates(listnode), expected) {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDeleteDuplicates2(t *testing.T){
	nums := []int{1}
	listnode := ArrToListNode(nums)
	expected :=ArrToListNode([]int{1})
	if reflect.DeepEqual(deleteDuplicates(listnode), expected) {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
func TestDeleteDuplicates3(t *testing.T){
	nums := []int{}
	listnode := ArrToListNode(nums)
	expected :=ArrToListNode([]int{})
	if reflect.DeepEqual(deleteDuplicates(listnode), expected) {
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
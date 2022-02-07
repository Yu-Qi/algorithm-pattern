package algo
import (
	"fmt"
	"testing"
	. "../ListNode"
	"reflect"
)
func TestDeleteDuplicates1(t *testing.T){
	nums := []int{1,2,3,3,4,4,5}
	listnode := ArrToListNode(nums)
	expected := []int{1,2,5}
	actual := deleteDuplicates(listnode)
	fmt.Printf("visit: %v\n", ArrFromListNode(actual))
	fmt.Printf("visit: %v\n", expected)	

	if reflect.DeepEqual(ArrFromListNode(actual), expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDeleteDuplicates2(t *testing.T){
	nums := []int{1,1,1,2,3}
	listnode := ArrToListNode(nums)
	actual := deleteDuplicates(listnode)
	fmt.Printf("visit: %v\n", ArrFromListNode(actual))
	expected := []int{2,3}
	if reflect.DeepEqual(ArrFromListNode(actual), expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDeleteDuplicates3(t *testing.T){
	nums := []int{1,1,1,2,3}
	listnode := ArrToListNode(nums)
	actual := deleteDuplicates(listnode)
	fmt.Printf("visit: %v\n", ArrFromListNode(actual))
	expected := []int{2,3}
	if reflect.DeepEqual(ArrFromListNode(actual), expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}

func TestDeleteDuplicates4(t *testing.T){
	nums := []int{1,2,2}
	listnode := ArrToListNode(nums)
	actual := deleteDuplicates(listnode)
	fmt.Printf("visit: %v\n", ArrFromListNode(actual))
	expected := []int{1}
	if reflect.DeepEqual(ArrFromListNode(actual), expected){
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}
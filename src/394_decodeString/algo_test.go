package algo
import (
	"testing"
	"fmt"
)

func TestDecodeString(t *testing.T){
	s := "3[a]2[bc]"
	actual := decodeString(s)
	// fmt.Printf("acutal: %v\n",actual) 
	if actual == "aaabcbc"{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}

}

func TestDecodeString2(t *testing.T){
	s := "3[a2[c]]"
	actual := decodeString(s)
	fmt.Printf("acutal: %v\n",actual) 
	if actual == "accaccacc"{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDecodeString3(t *testing.T){
	s := "2[abc]3[cd]ef"
	actual := decodeString(s)
	fmt.Printf("acutal: %v\n",actual) 
	if actual == "abcabccdcdcdef"{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDecodeString4(t *testing.T){
	s := "10[a]"
	actual := decodeString(s)
	fmt.Printf("acutal: %v\n",actual) 
	if actual == "aaaaaaaaaa"{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestDecodeString5(t *testing.T){
	s := "3[a10[bc]]"
	actual := decodeString(s)
	fmt.Printf("acutal: %v\n",actual) 
	if actual == "abcbcbcbcbcbcbcbcbcbcabcbcbcbcbcbcbcbcbcbcabcbcbcbcbcbcbcbcbcbc"{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}


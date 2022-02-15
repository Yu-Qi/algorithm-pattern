package algo
import "testing"

func TestEvalRPN(t *testing.T){
	tokens := []string{"2","1","+","3","*"}
	if evalRPN(tokens) ==9{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}

func TestEvalRPN2(t *testing.T){
	tokens := []string{"4","13","5","/","+"}
	if evalRPN(tokens) ==6{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
func TestEvalRPN3(t *testing.T){
	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}
	if evalRPN(tokens) ==22{
		t.Log("Success")
	} else{
		t.Error("Fail")
	}
}
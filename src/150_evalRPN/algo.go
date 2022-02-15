package algo

import (
    "strconv"
)

func operate(stack *[]int, op string){
	num1 := (*stack)[len(*stack)-2]
	num2 := (*stack)[len(*stack)-1]
	(*stack)[len(*stack)-2] = num1+num2
	if op == "+"{
		(*stack)[len(*stack)-2] = num1+num2
	}	else if op == "-"{
		(*stack)[len(*stack)-2] = num1-num2
	}	else if op == "*"{
		(*stack)[len(*stack)-2] = num1*num2
	}	else if op == "/"{
		(*stack)[len(*stack)-2] = num1/num2
	}
	(*stack) = (*stack)[:len(*stack)-1]

}
func evalRPN(tokens []string) int {
    stack := make([]int,0)
	for i:=0;i<len(tokens);i++{
		token := tokens[i]
		if token == "+"{
			operate(&stack,"+")
		}		else if token == "-"{
			operate(&stack,"-")
		}		else if token == "*"{
			operate(&stack,"*")	
		} else if token == "/"{
			operate(&stack,"/")
		} else{
			num, _ := strconv.Atoi(token) // result: i = -18
			stack = append(stack, num)
		}
	}
	return stack[0]
}
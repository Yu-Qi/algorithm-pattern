package algo

import (
    "strconv"
	"fmt"
)
func popNum(nums *[]int) int{
	num := (*nums)[len(*nums)-1]
	(*nums) = (*nums)[:len(*nums)-1]
	return num
}
func popToken(digits *[]string) string{
	idx:=len(*digits)-1
	token := ""
	for ;idx>=0;idx--{
		if (*digits)[idx] == "["{
			break
		}
	}
	
	for i:=idx+1;i<len(*digits);i++{
		token += (*digits)[i]
	}
	// pop
	(*digits) = (*digits)[:idx]

	return token
}
func multiple(s string, num int) string{
	res := ""
	for i:=0;i<num;i++{
		res += s
	}
	return res
}

func decodeString(s string) string {
	nums := make([]int,0)
	digits := make([]string,0)
	num_flag := false
	for i:=0;i<len(s);i++{
		token := s[i:i+1]
		switch token {
		case "1","2","3","4","5","6","7","8","9","0": 
			num, _ := strconv.Atoi(token) // result: i = -18
			if num_flag{
				prev_num := nums[len(nums)-1]
				num += prev_num*10
				nums[len(nums)-1] = num
			} else{
				nums = append(nums, num)
			}
			num_flag = true
		case "]":
			digit := popToken(&digits)
			num := popNum(&nums)
			digits = append(digits, multiple(digit, num))
			fmt.Printf("debug: digit%s,num:%d,digits:%v\n",digit,num,digits)
			num_flag = false
		default:
			digits = append(digits, token)
			num_flag = false
		}
	}
	res := ""
	for i:=0;i<len(digits);i++{
		res += digits[i]
	}
	return res
}


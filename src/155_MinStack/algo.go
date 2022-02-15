package algo

type MinStack struct {
    stack []int
	min []int
}
func Constructor() MinStack {
    return MinStack{
		stack: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)

	// 
    min := this.GetMin()
	if min < val{
		this.min = append(this.min, min)
	} else{
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0{
		return 0
	}
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0{
		return 1 <<31
	}
    return this.min[len(this.min)-1]
}
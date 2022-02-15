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
func (this *MinStack) remove(val int){
	idx :=0
	for ;i<len(this.min);i++{
		if this.min[idx] == val{
			this.min = append(this.min[:idx],this.min[idx+1:]...)
			return
		}
}
func (this *MinStack) insert(val int){
	idx :=0
	for ;i<len(this.min);i++{
		if val <= this.min[i]{
			break
		}
	}
	tmp := append(this.min[:idx],val)
	this.min = append(tmp, this.min[idx:])
}
func (this *MinStack) Push(val int)  {
    this.stack = append([]int{val}, this.stack...)
	insert(val)
}

func (this *MinStack) Pop()  {
	item := this.stack[0]
    this.stack = this.stack[1:]
	remove(item)
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0{
		return 0
	}
    return this.stack[0]
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0{
		return 1 <<31
	}
    return this.min[0]
}
package algo

type MyQueue struct {
	Queue []int
}


func Constructor() MyQueue {
    return MyQueue{
		Queue: []int{},
	}
}

func (this *MyQueue) Push(x int)  {
    this.Queue = append(this.Queue, x)
}


func (this *MyQueue) Pop() int {
	if len(this.Queue) <= 0{
		return 0
	}	
    item := this.Queue[0]
	this.Queue = this.Queue[1:]
	return item
}


func (this *MyQueue) Peek() int {
	if len(this.Queue) <= 0{
		return 0
	}
    return this.Queue[0]
}


func (this *MyQueue) Empty() bool {
    return len(this.Queue) ==0
}


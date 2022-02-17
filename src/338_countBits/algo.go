package algo

func countBits(n int) []int {
    if n==0{
        return []int{0}
    } else if n==1{
        return []int{0,1}
    }
    
    table := make([]int,n+1)
    table[0] = 0
    table[1] =1
    pow :=2
    for i:=2;i<=n;i++{
        if i>=pow && i<(pow<<1){
            table[i] = table[i-pow]+1
        }else{
            table[i] = 1
            pow = pow <<1
        }
    }
    return table
}
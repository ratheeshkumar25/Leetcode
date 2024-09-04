func longestConsecutive(nums []int) int {
    
    numSet := make(map[int]bool)

    for _, num := range nums{
        numSet[num] = true
    }
    start := []int{}
    for k,_:= range numSet{
        if !numSet[k-1]{
            start = append(start,k)
        }
    }
    count,max := 0,0
    for _,v:= range start{
        i:=v
        for numSet[i]{
            count++
            i++
        }
         if max < count{
        max  = count
    }
    count = 0
    }
   

    return max
}
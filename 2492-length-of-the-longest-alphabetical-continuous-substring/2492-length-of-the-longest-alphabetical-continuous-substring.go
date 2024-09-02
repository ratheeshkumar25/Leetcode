func longestContinuousSubstring(s string) int {
    res := 0
    curr :=1
    for i :=1; i<len(s);i++{
         //fmt.Println("----",string(s[i]),string(s[i-1]))
        if s[i-1]+1== s[i]{
            curr++
           // fmt.Println("////",curr)
        }else{
            res = max(curr,res)
            //fmt.Println("*****",curr,res)
            curr = 1
        }
    }
    res = max(curr,res)
    return res
}
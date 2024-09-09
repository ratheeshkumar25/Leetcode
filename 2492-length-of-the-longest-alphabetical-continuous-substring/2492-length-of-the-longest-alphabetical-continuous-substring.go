func longestContinuousSubstring(s string) int {
    res := 0
    curr := 1

    for i:=1;i<len(s);i++{
        if s[i-1]+1 == s[i]{
            curr++
        }else{
            res = max(res,curr)
            curr = 1
        }
    }
    res = max(res,curr)
    return res 
}
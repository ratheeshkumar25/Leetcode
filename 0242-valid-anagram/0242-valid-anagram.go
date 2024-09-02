func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    charCount := make(map[byte]int)

    for i := range s{
        charCount[s[i]]++
    }

    for i:= range t{
        charCount[t[i]]--
        //fmt.Println("char2",charCount[t[i]])
        if charCount[t[i]]<0{
            return false
        }
    }
    return true
}
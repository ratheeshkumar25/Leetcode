func topKFrequent(nums []int, k int) []int {
     freq := make(map[int]int)

     for _,num := range nums{
        freq[num]++
     }

     unique := make([]int,0,len(freq))
     for i := range freq{
        unique = append(unique,i)
     }
     //fmt.Println(unique)
     sort.Slice(unique, func(a,b int)bool{
        return freq[unique[a]] > freq[unique[b]]
     })

     return unique[:k]
}
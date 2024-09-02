func containsDuplicate(nums []int) bool {
     num := make(map[int]bool)
     for _,val := range nums{
        if num[val]{
            return true
        }
        num[val] = true 
     }
     return false 
}
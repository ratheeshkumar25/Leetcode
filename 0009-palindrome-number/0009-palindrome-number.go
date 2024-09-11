func isPalindrome(x int) bool {
    	temp := x
    var num int 

	for temp > 0 {
		rem := temp % 10
		num = (num*10 )+ rem
		temp = temp / 10
	}
	if x != num {
		return false
	} else {
		return true
	}
}
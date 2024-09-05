func findKthLargest(nums []int, k int) int {
     heapSort(nums)
     return nums[len(nums)-k]
}

func heapify(arr[]int,idx int, n int){
     larg := idx
     left := 2*idx+1
     right := 2*idx+2

     if left<n && arr[left]>arr[larg]{
        larg = left
     }
     if right<n && arr[right]>arr[larg]{
        larg = right
     }

     if larg != idx{
        arr[idx],arr[larg] = arr[larg],arr[idx]
        heapify(arr,larg,n)
     }
}

func buildHeap(arr[]int,n int){
    for i:= n/2-1; i>=0 ; i--{
        heapify(arr,i,n)
    }
}

func heapSort(arr[]int){
     n := len(arr)
     buildHeap(arr,n)
     for i:=n-1; i>0;i--{
        arr[0],arr[i]= arr[i],arr[0]
        heapify(arr,0,i)
     }
}
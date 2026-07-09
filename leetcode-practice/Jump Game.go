// 再次想贪心的原则 ：  局部最优就是全局最优-》我开始的时候就想到了一个思路，从后往前数，如果可以达到的话就迭代为前面那个元素是否可达的问题 



func CanJump (nums []int) bool{
  i := len(nums)-1
  for {
    if i==0{
      return true
    }
    for j := 1;j<=i;j++ {
      if nums[i-j] >= j{
        i = i-j
        break
      } 
      if j==i {
        return false
      }
    }
    
  }


  
}

// 之前的错误 ， 12行写成了 if num[i-j] + j >= i 这里有问题，其实也是i j 和平时的不一样导致，这时候i是数组的末尾指针，但是j表示的是从后道歉到i到举例，
// 所以第二个指针式i-j，而不是j ，所以if条件直接写num[i-j]大于等于到i到距离j就可以 

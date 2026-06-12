func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    MaxLen := int(0)
    count := 0

    for _,v :=range nums {
        m[v] = true
    }

    for i:=0;i<len(nums);i++ {
        count = 1
        num := nums[i] //按照go语言的话是每次都开辟新的num空间还是只有第一次 ？
        for {
            if _,ok := m[num-1]; ok{
                count++
                num--
                continue
            }
            break
        }
        MaxLen = max(MaxLen,count)

    }

    return MaxLen

}

// Q 为什么你的代码超出了时间限制？
// 因为写for i:=0;i<len(nums);i++这样循环的话会导致重复计算，比如之前从4算过了3但是遇到3的时候还会当成新的算
//=====================v2========================= 还是超时了
func longestConsecutive(nums []int) int {
    m := make(map[int]bool, len(nums))
    maxLen := int(0)
    count := 0

    for _,v :=range nums {
        m[v] = true
    }
    
    for _,num := range nums {
        if!m[num-1]{
            count = 1
            for {
                if _,ok := m[num+1]; ok{
                    num++
                    count++
                    continue
                }
                break
            }
        }
        maxLen = max(maxLen,count)
    }

    // for i:=0;i<len(nums);i++ {
    //     count = 1
    //     num := nums[i] //按照go语言的话是每次都开辟新的num空间还是只有第一次 
    //     for {
    //         if _,ok := m[num-1]; ok{
    //             count++
    //             num--
    //             continue
    //         }
    //         break
    //     }


    //     maxLen = max(MaxLen,count)

    // }

    return maxLen

}
//=====================v3：遍历map而不是开始给的数组（去重）========================
func longestConsecutive(nums []int) int {
    m := make(map[int]bool, len(nums))
    maxLen := int(0)
   

    for _,v :=range nums {
        m[v] = true
    }
    
    for num,_ := range m { //修改这里
        count := 1 //优化：最小作用域原则
        if!m[num-1]{
            
            for {
                if _,ok := m[num+1]; ok{
                    num++
                    count++
                    continue
                }
                break
            }
        }
        maxLen = max(maxLen,count)
    }

    // for i:=0;i<len(nums);i++ {
    //     count = 1
    //     num := nums[i] //按照go语言的话是每次都开辟新的num空间还是只有第一次 
    //     for {
    //         if _,ok := m[num-1]; ok{
    //             count++
    //             num--
    //             continue
    //         }
    //         break
    //     }


    //     maxLen = max(MaxLen,count)

    // }

    return maxLen

}


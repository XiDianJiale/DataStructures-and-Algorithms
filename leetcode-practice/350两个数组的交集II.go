func intersect(nums1 []int, nums2 []int) []int {
    m1 := make(map[int]int)
    arr := make([]int,0,10)

    for _,v :=range nums1{
        m1[v]++
    }
    for _,v := range nums2{
        if _,ok := m1[v]; ok{
            if m1[v] != 0{
                m1[v]--
                arr = append(arr, v)
            }
        }
    }

    return arr

}

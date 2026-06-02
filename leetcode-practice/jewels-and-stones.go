// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
    m := make(map[rune]bool) //不能用char,第一次出错
    res := int(0)
    for _,v:=range jewels{
        m[v] = true
    }
    for _,v:= range stones{
        if _,ok:=m[v]; ok{
            res++
        }
    }

    return res

}

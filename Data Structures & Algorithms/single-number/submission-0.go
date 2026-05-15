func singleNumber(nums []int) int {
    var xor=0
    for _,i := range nums{
        xor^=i
    }
    return xor
}

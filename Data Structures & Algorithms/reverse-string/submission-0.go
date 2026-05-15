func reverseString(s []byte) {
var l=len(s)
var left,right=0,l-1
for left <right{
    s[left],s[right]=s[right],s[left]
    //var tmp=s[right]
    //s[right]=s[left]
    //s[left]=tmp
    left+=1
    right-=1
}
}

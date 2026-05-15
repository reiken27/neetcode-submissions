func isPalindrome(s string) bool {
	var left, right=0,len(s)-1
	for left<right{
		var l,r=rune(s[left]),rune(s[right])
		if !unicode.IsLetter(l)&&!unicode.IsDigit(l) {
			left+=1
			continue
		}
		if !unicode.IsLetter(r)&&!unicode.IsDigit(r){
			right-=1
			continue
		}
		l=unicode.ToLower(l)
		r=unicode.ToLower(r)
	    if l==r{
			left+=1
		    right-=1
		}else{
		    return false
		}

	}
	return true
}

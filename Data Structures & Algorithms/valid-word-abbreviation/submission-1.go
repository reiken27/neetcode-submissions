func validWordAbbreviation(word string, abbr string) bool {

    var i,j=0,0
	for i<len(word) && j<len(abbr){
		if word[i]==abbr[j]{
			i+=1
			j+=1
		}else if unicode.IsLetter(rune(abbr[j]))|| abbr[j]=='0'{
			return false
		}else if unicode.IsDigit(rune(abbr[j])){
			var start = j
			for j < len(abbr) && unicode.IsDigit(rune(abbr[j])){
				j+=1
			}
			var number,_=strconv.Atoi(abbr[start:j])
			i+=number
		}

	}
	return i == len(word) && j == len(abbr)
}

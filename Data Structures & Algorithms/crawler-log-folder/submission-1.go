func minOperations(logs []string) int {
	var counter=0
	var i=0
	for i<len(logs){
		var log =(logs[i])
		if log[0]=='.'{
			if log[1]=='.'{
				counter-=1
				if counter<=0{counter=0}
			}
		}else{
        counter+=1
		}
		i++
	}
return counter
}

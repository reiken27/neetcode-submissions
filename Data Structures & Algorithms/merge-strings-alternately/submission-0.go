func mergeAlternately(word1 string, word2 string) string {
var res strings.Builder
var word1_p,word2_p=0,0
var word1_l,word2_l=len(word1),len(word2)

for word1_p<word1_l && word2_p<word2_l{
res.WriteByte(word1[word1_p])
res.WriteByte(word2[word2_p])
word1_p++
word2_p++
} 
if word1_p<word1_l{
	res.WriteString(word1[word1_p:])

}
if word2_p<word2_l{
	res.WriteString(word2[word2_p:])
}
return res.String()
}

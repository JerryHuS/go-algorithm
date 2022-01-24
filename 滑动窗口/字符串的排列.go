package 滑动窗口

//换句话说，s1 的排列之一是 s2 的 子串
//abc bcasde
//"adc"
//"dcda"
func checkInclusion(s1 string, s2 string) bool {
	num1 := len(s1)
	num2 := len(s2)
	if num1 > num2 {
		return false
	}
	var cnt1, cnt2 [26]int
	for i := 0; i < num1; i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for j := 1; j+num1-1 < num2; j++ {
		cnt2[s2[j-1]-'a']--
		cnt2[s2[num1+j-1]-'a']++
		if cnt2 == cnt1 {
			return true
		}
	}
	return false
}

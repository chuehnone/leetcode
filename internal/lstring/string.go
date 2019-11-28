package lstring

func BalancedStringSplit(s string) int {
	cnt := 0

	sLen := len(s)
	var last string
	push := 0
	for i := 0 ; i < sLen ; i++ {
		ch := string(s[i])
		if push == 0 {
			last = ch
			cnt++
		}

		if ch == last {
			push++
		} else {
			push--
		}
	}

	return cnt
}

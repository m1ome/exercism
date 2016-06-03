package secret

var wink, dwink, clo, jmp = "wink", "double blink", "close your eyes", "jump"

func Handshake(code int) []string {
	var answer []string

	if code <= 0 {
		return answer
	}

	if code&(1<<0) > 0 {
		answer = append(answer, wink)
	}

	if code&(1<<1) > 0 {
		answer = append(answer, dwink)
	}

	if code&(1<<2) > 0 {
		answer = append(answer, clo)
	}

	if code&(1<<3) > 0 {
		answer = append(answer, jmp)
	}

	if code&(1<<4) > 0 {
		for i, j := 0, len(answer)-1; i < j; i, j = i+1, j-1 {
			answer[i], answer[j] = answer[j], answer[i]
		}
	}

	return answer
}

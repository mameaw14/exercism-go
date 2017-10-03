package secret

const testVersion = 2

func Handshake(code uint) []string {
	var seq []string
	if code%2 >= 1 {
		seq = append(seq, "wink")
	}
	if code%4 >= 2 {
		seq = append(seq, "double blink")
	}
	if code%8 >= 4 {
		seq = append(seq, "close your eyes")
	}
	if code%16 >= 8 {
		seq = append(seq, "jump")
	}
	var reverse bool
	if code%32 >= 16 {
		reverse = true
	} else {
		reverse = false
	}

	if reverse {
		var rev_seq []string
		for i := len(seq) - 1; i >= 0; i-- {
			rev_seq = append(rev_seq, seq[i])
		}
		return rev_seq
	}
	return seq
}

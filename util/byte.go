package util

func ReverseBytes(x []byte) []byte {
	y := make([]byte, len(x))
	max := len(x) - 1
	for i := 0; i <= max; i++ {
		y[max-i] = x[i]
	}
	return y
}

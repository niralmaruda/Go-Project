package iteration

var repeatCount int = 6

func Repeat(ch string) (res string) {
	for i := 0; i < repeatCount; i++ {
		res += ch
	}
	return
}
func ExampleRepeat(ch string, N int) (res string) {
	for i := 0; i < N; i++ {
		res += ch
	}
	return
}

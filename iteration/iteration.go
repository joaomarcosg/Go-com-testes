package iteration

func Iteration(character string) string {
	var iterations string
	for i := 0; i < 5; i++ {
		iterations += character
	}
	return iterations
}

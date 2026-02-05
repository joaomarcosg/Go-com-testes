package iteration

const iterationsQuantity = 5

func Iteration(character string) string {
	var iterations string
	for range iterationsQuantity {
		iterations += character
	}
	return iterations
}

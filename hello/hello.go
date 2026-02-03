package main

const portugueseHelloPrefix = "Olá, "

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo"
	}
	return portugueseHelloPrefix + name
}

func main() {

}

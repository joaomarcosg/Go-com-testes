package main

const spanish = "spanish"
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return portugueseHelloPrefix + name
}

func main() {

}

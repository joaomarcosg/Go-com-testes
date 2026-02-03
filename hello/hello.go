package main

const spanish = "spanish"
const french = "french"

const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjuor, "

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return portugueseHelloPrefix + name
}

func main() {

}

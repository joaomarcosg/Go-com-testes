package main

const spanish = "spanish"
const french = "french"

const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjuor, "

func Hello(name, language string) string {

	prefix := portugueseHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {

}

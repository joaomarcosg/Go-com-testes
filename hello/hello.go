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
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = portugueseHelloPrefix
	}
	return
}

func main() {

}

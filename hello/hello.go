package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "spanish" {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

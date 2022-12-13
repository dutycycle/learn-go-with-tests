package main

const ENGLISH = "english"
const SPANISH = "spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == SPANISH {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

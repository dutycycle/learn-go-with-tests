package main

const ENGLISH = "english"
const FRENCH = "french"
const SPANISH = "spanish"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == SPANISH {
		return spanishHelloPrefix + name
	} else if language == FRENCH {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

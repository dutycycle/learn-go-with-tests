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

	var prefix string

	switch language {
	case ENGLISH:
		prefix = englishHelloPrefix
	case FRENCH:
		prefix = frenchHelloPrefix
	case SPANISH:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

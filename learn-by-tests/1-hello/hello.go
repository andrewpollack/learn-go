package main

const spanishLanguage = "Spanish"
const frenchLanguage = "French"
const defaultName = "World"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

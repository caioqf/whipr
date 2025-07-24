package main

import (
	"log"
)

func Translate(text string) string {
	language, err := assertLanguague(text)

	if err == nil {
		log.Println("error asserting language of text: ", err)
	}

	return translatorCall(text, language)
}

func assertLanguague(text string) (string, error) {
	return "en-EN", nil
}

func translatorCall(textToTranslate, language string) string {
	return textToTranslate
}

package cmd

import (
	"log"
	"os/exec"
)

func TranslateContent(text string) (string, error) {
	language, err := assertLanguage(text)
	if err != nil {
		log.Printf("Error asserting language: %v", err)
		return "", err
	}
	return translatorCall(text, language), nil
}

func assertLanguage(text string) (string, error) {
	return "en-EN", nil
}

func translatorCall(textToTranslate, language string) string {
	return textToTranslate
}

func DisplayTranslated(text string) {
	translated, err := TranslateContent(text)
	if err != nil {
		log.Printf("Translation error: %v", err)
		translated = "Error: " + err.Error()
	}
	if ShouldUsePopup() {
		log.Println("Using popup display")
		err := exec.Command("zenity", "--info", "--text", translated).Run()
		if err != nil {
			log.Printf("zenity error: %v", err)
		}
	}
	if ShouldUseNotify() {
		log.Println("Using notification display")
		err := exec.Command("notify-send", "Translation", translated).Run()
		if err != nil {
			log.Printf("notify-send error: %v", err)
		}
	}
}

// TODO -> para imagens
// func TranslateImage(imagePath string) {
// 	// Extrai texto com OCR (ex: tesseract)
// 	text, err := exec.Command("tesseract", imagePath, "stdout").Output()
// 	if err != nil {
// 		log.Printf("OCR error: %v", err)
// 		return
// 	}
// 	DisplayTranslated(string(text))  // Reutiliza a exibição
// }

package app

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Ramdomize Функция возвраащет сгенерированный id и короткую ссылку
func Ramdomize(fullStr string) string {

	//rand.Seed(time.Now().UnixNano())
	shortStr := randSeq(6)
	// составляем ID из 4 случаенный частей
	id := randID(5) + "-" + randID(5) + "-" + randID(5) + "-" + randID(5)
	//fmt.Println(id)
	// записываем данные в Json
	saveJSON(fullStr, shortStr, id)
	return shortStr
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Функция генерации сокращенной ссылки
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var count = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Функция генерации id
func randID(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = count[rand.Intn(len(letters))]
	}
	return string(b)
}

const settingsFilename = "clientURLs.json"

type URLs struct {
	FullURL string
	ShorURL string
	ID      string
	//Date    string
}

type Settings struct {
	URLs []URLs
}

// Функция записи данных в JSON
func saveJSON(fullSrtToJSON, shortStrToJSON, id string) {
	// Читаем файл
	rawDataIn, err := os.ReadFile(settingsFilename)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var settings Settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	// Записываем данные в структуру
	newClient := URLs{
		FullURL: fullSrtToJSON,
		ShorURL: shortStrToJSON,
		ID:      id,
		//Date: string(time.Now()),
	}

	// Дополняем новые данные
	settings.URLs = append(settings.URLs, newClient)

	rawDataOut, err := json.MarshalIndent(&settings, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	// записываем новые данные в JSON
	err = os.WriteFile(settingsFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}

}

// JSONDecoder Функция поска полной ссылки в JSON по ID
func JSONDecoder(idUrlToResponse string) string {
	var urlToResponse string
	rawDataIn, err := os.ReadFile(settingsFilename)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var settings Settings
	err = json.Unmarshal(rawDataIn, &settings)
	if err != nil {
		log.Fatal("Invalid settings format:", err)
	}
	//json.Unmarshal([]byte(settingsFilename), &settings)

	for _, v := range settings.URLs {
		if v.ID == idUrlToResponse {
			//fmt.Println(v.ID)
			urlToResponse = strings.ReplaceAll(v.FullURL, "url=", "")
		}

	}

	return urlToResponse
}

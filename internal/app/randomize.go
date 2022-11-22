package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

// Ramdomize Функция возвраащет сгенерированный id и короткую ссылку
func Ramdomize(fullStr string) string {

	//rand.Seed(time.Now().UnixNano())
	shortStr := randSeq(6)
	// составляем ID из 4 случаенный частей
	id := randId(5) + "-" + randId(5) + "-" + randId(5) + "-" + randId(5)
	fmt.Println(id)
	// записываем данные в Json
	saveJson(fullStr, shortStr, id)
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
func randId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = count[rand.Intn(len(letters))]
	}
	return string(b)
}

const settingsFilename = "clientURLs.json"

type URLs struct {
	FullUrl string
	ShorUrl string
	ID      string
	//Date    string
}

type Settings struct {
	URLs []URLs
}

// Функция записи данных в JSON
func saveJson(fullSrtToJson, shortStrToJson, id string) {
	// Читаем файл
	rawDataIn, err := ioutil.ReadFile(settingsFilename)
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
		FullUrl: fullSrtToJson,
		ShorUrl: shortStrToJson,
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
	err = ioutil.WriteFile(settingsFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated settings file:", err)
	}

}

// JsonDecoder Функция поска полной ссылки в JSON по ID
func JsonDecoder(idUrlToResponse string) string {
	var urlToResponse string
	rawDataIn, err := ioutil.ReadFile(settingsFilename)
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
			urlToResponse = strings.ReplaceAll(v.FullUrl, "url=", "")
		}

	}

	return urlToResponse
}

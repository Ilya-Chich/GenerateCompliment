package controllers

import (
	"GenerateCompliment/models/words"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GenerateCompliment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": concatCompliment()})
}

func concatCompliment() string {
	compliment := []string {getTemplateSample(), getAddressSample(), getAdverbSample(), getAdjectiveSample()}
	return strings.Join(compliment[:], " ")
}

func getTemplateSample() string {
	template := words.Templates()
	return getRandomSample(template)
}

func getAddressSample() string {
	address := words.Addresses()
	return getRandomSample(address)
}

func getAdverbSample() string {
	adverb := words.Adverbs()
	return getRandomSample(adverb)
}

func getAdjectiveSample() string {
	adjective := words.Adjectives()
	return getRandomSample(adjective)
}

func getRandomSample(sample []string)  string {
	sampleLength := len(sample)
	rand.Seed(time.Now().UnixNano())
	return sample[rand.Intn(sampleLength)]
}

func addArticle(word string) string {
	trimmedWord := strings.Trim(word, " ")
	if trimmedWord == "" {
		log.Fatal("empty string passed")
	}
	firstLetter := trimmedWord[0:1]
	switch firstLetter {
	case
		"a",
		"e",
		"i",
		"o",
		"u":
		return "an " + trimmedWord
	}
	return "a " + trimmedWord
}
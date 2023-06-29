package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"encoding/json"
	"os"
	"strings"

)

type Article struct {
	Title string
}

func filterArticles(articles []Article, keyword string) []Article {
	filteredArticles := []Article{}

	for _, article := range articles {
		if strings.Contains(strings.ToLower(article.Title), keyword) {
			filteredArticles = append(filteredArticles, article)
		}
	}

	return filteredArticles
}

func main() {
	c := colly.NewCollector()
	articles := []Article{}

	c.OnHTML(".link-gray", func(e *colly.HTMLElement) {
		title := e.Text
		article := Article{Title: title}
		articles = append(articles, article)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.buzzfeed.com/tvandmovies")
	for _, article := range articles {
		fmt.Println(article.Title)
	}

	barbieArticles := filterArticles(articles, "barbie")

	jsonData, err := json.MarshalIndent(barbieArticles, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))

	file, err := os.Create("outputBarbie.json")
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}
	
}
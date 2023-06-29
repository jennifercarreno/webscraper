package main

import (
    "fmt"
	"github.com/gocolly/colly"
	"encoding/json"
	"os"
)

type Article struct {
	Title string
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

	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing to JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))

	file, err := os.Create("output.json")
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
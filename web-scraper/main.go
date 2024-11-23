package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"web-scraper/db"
	"web-scraper/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-co-op/gocron"
)

func main() {
	db.InitDB()
	f, err := os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("failed to open logs: ", err)
	}
	log.SetOutput(f)

	var baseUrl = "https://www.scrapingcourse.com/ecommerce/page/"
	var pages = [3]int{1, 2, 3}
	var wg sync.WaitGroup

	s := gocron.NewScheduler(time.Local)
	_, err = s.Every(5).Minutes().Do(func() {
		for i := 0; i < len(pages); i++ {
			wg.Add(1)
			url := fmt.Sprintf("%s%d", baseUrl, pages[i])
			go fetchAndSaveProducts(url)
			pages[i] = pages[i] + len(pages)
		}

		defer wg.Done()
	})

	if err != nil {
		log.Fatalln("Failed job with error: ", err)
	}

	s.StartBlocking()
}

func fetchAndSaveProducts(url string) {

	log.Println("Now Fetching: ", url)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln("Failed to scrape website")
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("Failed to parse response body", err)
	}

	doc.Find("li.product").Each(func(i int, s *goquery.Selection) {
		product := models.Product{}

		product.Name = s.Find("h2").First().Text()
		product.Price = s.Find("span.price").First().Text()

		err := product.Save()

		log.Println(product)

		if err != nil {
			log.Fatal("Failed to store product", err)
		}
	})
}

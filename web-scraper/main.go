package main

import (
	"log"
	"net/http"
	"os"
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

	s := gocron.NewScheduler(time.Local)
	_, err = s.Every(5).Minutes().Do(fetchAndSaveProducts)

	if err != nil {
		log.Fatalln("Failed job with error: ", err)
	}

	s.StartBlocking()
}

func fetchAndSaveProducts() {

	res, err := http.Get("https://www.scrapingcourse.com/ecommerce/")

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

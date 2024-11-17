package main

import (
	"log"
	"net/http"
	"web-scraper/db"
	"web-scraper/models"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	db.InitDB()

	res, err := http.Get("https://www.scrapingcourse.com/ecommerce/")

	if err != nil {
		log.Fatalln("Failed to scrape website")
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("HTTP Error %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("Failed to parse response body", err)
		return
	}

	doc.Find("li.product").Each(func(i int, s *goquery.Selection) {
		product := models.Product{}

		product.Name = s.Find("h2").First().Text()
		product.Price = s.Find("span.price").First().Text()

		err := product.Save()

		if err != nil {
			log.Fatal("Failed to store product", err)
		}
	})

	// byteBody, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatal("Failed to read the buffer")
	// 	return
	// }

	// html := string(byteBody)
	// fmt.Println(html)
}

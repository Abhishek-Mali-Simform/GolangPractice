package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func get(url string) (res *http.Response) {
	var err error
	res, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func close(res *http.Response) {
	res.Body.Close()
}
func docReader(res *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func journalDev() {
	res := get("http://journaldev.com")
	defer close(res)

	//num, err := io.Copy(os.Stdout, res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Number of Bytes Copied: ", num)
	doc := docReader(res)
	doc.Find("p").Each(func(index int, s *goquery.Selection) {
		fmt.Println("next")
		txt := s.Text()
		fmt.Printf("Article %d: %s\n\n", index, txt)
	})
}

func metalSucks() {
	res := get("http://metalsucks.net")
	defer close(res)
	doc := docReader(res)
	doc.Find(".left-content article .post-title").Each(func(index int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n\n", index, title)
	})
}

func twitterTrending() {
	res := get("https://twitter.com/explore/tabs/trending")
	defer close(res)
	doc := docReader(res)
	doc.Find("span.css-901oao css-16my406 r-poiln3 r-bcqeeo r-qvutc0").Each(func(index int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Trending %d: %s\n\n", index, title)
	})
}

func main() {
	journalDev()
	metalSucks()
	twitterTrending()
}

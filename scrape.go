package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk"),
	)

	//Title
	c.OnHTML(".title--3s1R8", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	//Date
	c.OnHTML(".sub-title--37mkY", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	//Price
	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[1]/div", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[1]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[2]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[3]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[4]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	//Condtion
	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[5]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[6]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[7]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[8]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[9]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[2]/div[10]", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[1]/div/div[2]/div[3]/div/div[2]", func(e *colly.XMLElement) {
		fmt.Println("Description :")
		fmt.Println(e.Text)
	})

	c.OnXML("//*[@id=\"app-wrapper\"]/div[1]/div[2]/div/div/div[2]/div[3]/div[2]/div/div[2]/div/div/div[1]/div[2]/div[1]/button/div[2]/div[2]/div", func(e *colly.XMLElement) {

		fmt.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://ikman.lk/en/ad/toyota-chr-gt-turbo-2018-for-sale-colombo-170")
}

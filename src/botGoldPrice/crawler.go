package botGoldPrice

import (
	"time"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type Crawler struct {
}

func NewCrawler() *Crawler {
	return &Crawler{}
}

func (c *Crawler)GetAll() []string {
	begin := time.Date(2004, time.Month(7), 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	return c.GetDateRange(begin, end)
}

func (c *Crawler)GetDateRange(begin, end time.Time) []string {
	slice := make([]string, 0)
	for date := begin; date.Before(end); date = date.Add(time.Duration(24) * time.Hour) {
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}
		result := c.GetOneDay(date)
		if len(result) > 0 {
			slice = append(slice, result)
		}
	}
	log.Println("Finished getting.")
	return slice
}

func (c *Crawler)GetOneDay(date time.Time) string {
	dateStr := fmt.Sprintf("%04d%02d%02d", date.Year(), date.Month(), date.Day())
	fmt.Printf("Getting %s...\n", dateStr)
	url := "http://rate.bot.com.tw/Pages/UIP005/UIP00511.aspx?lang=zh-TW&whom=GB0030001000&afterOrNot=0&curcd=TWD&date=" + dateStr
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Printf("get url has error, %s", err)
		return ""
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("get url has error, %s", err)
		return ""
	}

	return string(b)
}

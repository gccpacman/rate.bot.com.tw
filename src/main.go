package main

import (
	"log"
	"os"
	"time"

	"github.com/realhu1989/rate.bot.com.tw/src/botGoldPrice"
)

func main() {

	begin := time.Date(2000, time.Month(10), 30, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	crawler := botGoldPrice.NewCrawler()
	htmls := crawler.GetDateRange(begin, end)
	records := make([]botGoldPrice.Record, 0)

	for _, html := range htmls {
		records = append(records, botGoldPrice.NewParser(html).Parse()...)
	}

	logfile := "records.csv"
	if !FileExists(logfile) {
		CreateFile(logfile)
	}
	f, _ := os.Create(logfile)
	log.SetOutput(f)

	for _, record := range records {
		log.Printf("%04d-%02d-%02d, buy: %f, sell: %f\n", record.Date.Year(), record.Date.Month(), record.Date.Day(), record.Buy, record.Sell)
	}

}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

package scraper

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func InitHTTPClient() {
	httpClient = resty.New().
		SetTimeout(10*time.Second).
		SetHeader("User-Agent", "Mozilla/5.0 (compatible; EconomicPulseBot/1.0)").
		SetRetryCount(2).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)

	log.Println("HTTP client initialized")
}

func GetClient() *resty.Client {
	if httpClient == nil {
		InitHTTPClient()
	}
	return httpClient
}

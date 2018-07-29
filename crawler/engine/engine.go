package engine

import (
	"go_crawler_study/crawler/fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, value := range seed {
		requests = append(requests, value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.FetchWebContent(r.Url)
		if err != nil {
			log.Printf("Fetch: error, Fetch Url %s: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)

		}

	}
}

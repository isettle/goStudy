package engine

import (
	"fmt"
	"log"

	"imooc.com/crawler/fetcher"
)

type SimpleEngine struct {
}

func (se SimpleEngine) Run(seeds ...Request) {
	// 入列
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	// 出列
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("+got item: %+v", item)
		}
	}

	fmt.Println("crawler done")
}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.Url)

	log.Printf("Fetching url: %s", r.Url)
	if err != nil {
		log.Printf("error fetcher url: %s, error: %v", r.Url, err)
		return ParserResult{}, err
	}
	parserResult := r.ParserFun(body)

	return parserResult, nil
}

package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			//log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver:error saving item %v:%v", item, err)
			}
		}
	}()
	return out

}

func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	response, err := client.Index().Index("dating_profile").Type("ygdy8").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return response.Id, nil

}

package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func InitElasticSearchClient() (client *elasticsearch.Client) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://111.229.3.130:30920"},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

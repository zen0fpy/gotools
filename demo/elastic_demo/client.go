package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
	"time"
)

type Elastic struct {
	Host string
	Port int64
}

type ServerConfig struct {
	Elastic Elastic
}

func NewClient(conf ServerConfig) *elastic.Client {
	url := fmt.Sprintf("http://%s:%d", conf.Elastic.Host, conf.Elastic.Port)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetHealthcheck(true),
		elastic.SetGzip(true),
		elastic.SetHealthcheckTimeout(5*time.Second),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(os.Stdout, "[ELASTIC] ", log.LstdFlags)),
	)
	if err != nil {
		log.Fatalf("Failed to create elastic client, %s\n", err)
	}

	info, code, err := client.Ping(url).Do(context.Background())
	if err != nil {
		log.Fatalf("can not connect to elastic_search, %s\n", err)
	}
	fmt.Printf("info: %s, code: %d\n", info.Name, code)

	esVersion, err := client.ElasticsearchVersion(url)
	if err != nil {
		log.Fatalf("failed to query es verison: %s\n", err)

	}
	fmt.Printf("ES Version: %s\n", esVersion)

	return client
}

func CreateIndexTwitter(client *elastic.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}

	if exists {
		fmt.Printf("index has exists!!!")
		return
	}

	mapping := `{
		"setting": {
				"number_of_shards": 1,
				"number_of_replicas": 0,
		},
		"mappings": {
			"doc": {
				"properties": {
					"user": {
						"type": "keyword",
					},
					"message": {
						"type": "text",
						"store": true,
						"fielddata": true,
					},
					"retweets": {
						"type": "long",
					},
					"tags": {
						"type": "keyword",
					},
					"location": {
						"type": "geo_point",
					},
					"suggest_field": {
						"type": "completion",
					},
				}
			}
		}
	}
`
	createIndex, err := client.CreateIndex("twitter").Body(mapping).Do(ctx)
	if err != nil {
		log.Fatalf("failed to create index, err: %s\n", err.Error())
	}

	if createIndex.Acknowledged {
		log.Fatalf("not ack! %s\n", createIndex.Index)
	}
}

func Add(client *elastic.Client, index string, obj interface{}) error {
	result, err := client.Index().Index(index).BodyJson(obj).Do(context.Background())
	if err != nil {
		return fmt.Errorf("add obj to es failed, err %s\n", err.Error())
	}
	fmt.Printf("Index: %s, Type: %s, Id: %s, result: %s\n", result.Index, result.Type, result.Id, result.Result)
	return nil
}

func Delete(client *elastic.Client, index string, id string) error {
	result, err := client.Delete().Index(index).Id(id).Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete index: %s, id: %s, err: %s\n", index, id, err.Error())
	}
	fmt.Printf("delete result: %s\n", result.Id)
	return nil

}

func Update(client *elastic.Client, index string, id string, content map[string]interface{}) error {

	result, err := client.Update().Index(index).Id(id).Doc(content).Do(context.Background())
	if err != nil {
		return fmt.Errorf("can not update at index: %s, id: %s, err: %s\n", index, id, err.Error())
	}
	fmt.Printf("update result: %s\n", result.Index)
	return nil
}

func Get(client *elastic.Client, index string, id string) error {
	result, err := client.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get from index: %s, id: %s, err:%s\n", index, id, err.Error())
	}

	encodeObj, _ := result.Source.MarshalJSON()
	fmt.Printf("result: %s\n", encodeObj)
	return nil
}

func Search(client *elastic.Client, index string, obj interface{}, query elastic.Query) error {
	result, err := client.Search(index).Query(query).Do(context.Background())
	if err != nil {
		fmt.Printf("can not search from index: %s\n", err.Error())
	}

	parseResult(result, obj)
	return nil
}

func List(client *elastic.Client, index string, obj interface{}, page, size int) error {
	if page < 1 {
		page = 1
	}

	result, err := client.Search(index).Size(size).From(size * (page - 1)).Do(context.Background())
	if err != nil {
		return fmt.Errorf("search failed, err: %s\n", err.Error())
	}
	parseResult(result, obj)
	return nil
}

func parseResult(result *elastic.SearchResult, obj interface{}) {
	for _, item := range result.Each(reflect.TypeOf(obj)) {
		fmt.Printf("list item: %v\n", item)
	}
}

func DeleteIndex(client *elastic.Client, index string) error {
	result, err := client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete index: %s, err: %s\n", index, err.Error())

	}
	fmt.Printf("delete index succ! %v\n", result.Acknowledged)
	return nil

}

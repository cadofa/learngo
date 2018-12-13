package main

import (
	"gopkg.in/olivere/elastic.v5"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"log"
)

func main(){
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host, index string) error{
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return  rpcsupport.ServeRpc(":1234",
		    &persist.ItemSaverService{
			    Client: client,
			    Index: index,
		    })
}

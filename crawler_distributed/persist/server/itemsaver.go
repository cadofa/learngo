package main

import (
	"gopkg.in/olivere/elastic.v5"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"log"
	"fmt"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	"flag"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main(){
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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

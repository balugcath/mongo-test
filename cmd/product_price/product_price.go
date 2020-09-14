package main

import (
	"flag"
	"log"
	"net"

	"github.com/jamiealquiza/envy"
	"google.golang.org/grpc"

	"github.com/balugcath/mongo-test/api"
	"github.com/balugcath/mongo-test/internal/handler"
	"github.com/balugcath/mongo-test/internal/storage"
)

var (
	mongoURI        string
	mongoDatabase   string
	mongoCollection string
	gRPCListen      string
)

func main() {
	flag.StringVar(&mongoURI, "mongo_uri", "mongodb://toor:toor@localhost:27017", "mongodb connect uri")
	flag.StringVar(&mongoDatabase, "database", "test", "mongodb database name")
	flag.StringVar(&mongoCollection, "collection", "test", "mongodb database name")
	flag.StringVar(&gRPCListen, "listen", ":5005", "grpc listen address")

	envy.Parse("PRODUCT_PRICE")
	flag.Parse()

	cli, err := storage.NewMongoClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	coll := storage.NewMongoCollection(cli, mongoDatabase, mongoCollection)

	h, err := handler.NewGRPCHandler(coll, coll)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", gRPCListen)
	if err != nil {
		log.Fatal(err)
	}

	serv := grpc.NewServer()
	api.RegisterProductPriceServer(serv, h)

	if err := serv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

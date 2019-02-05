package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/matalmeida/hashlab-service-1-golang/product"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing sub command: get")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server %v", err)
	}
	client := product.NewProductsClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "get":
		err = get(context.Background(), client, flag.Arg(1), flag.Arg(2))
	default:
		err = fmt.Errorf("unknown sub command %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func get(ctx context.Context, client product.ProductsClient, userID string, productID string) error {
	queryIDS := &product.Ids{
		UserID:    userID,
		ProductID: productID,
	}

	p, err := client.WithDiscount(ctx, queryIDS)
	if err != nil {
		return fmt.Errorf("could not fetch discount")
	}

	fmt.Println(proto.MarshalTextString(p))

	return nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("please specify credentials json as first argument")
		return
	}
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Args[1])
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc("Ada").Set(ctx,
		map[string]interface{}{
			"first": "Ada",
			"last":  "Lovelace",
			"born":  1815,
		})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	_, err = client.Collection("users").Doc("Alan").Set(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

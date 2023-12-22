package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

var ()

func main() {
	// fmt.Println(rdbActivity.Get(context.Background(), "lucky_egg_banner:6119f5358a6f3e62d1f5be81:9"))

	// uri := "mongodb://root:Oemkvdjo&h8yU%23rz@s-j6c08d4954641db4.mongodb.rds.aliyuncs.com:3717/admin"
	// // Use the SetServerAPIOptions() method to set the Stable API version to 1
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// // Create a new client and connect to the server
	// client, err := mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// // Send a ping to confirm a successful connection
	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// // b, _ := client.Database("activity").ListCollectionNames(context.Background(), bson.M{})
	// // fmt.Println(b)
	// // client.Database("admin").RunCommand(context.Background(), bson.D{{"status", 1}}).Decode(&result)
	// // fmt.Println(result)

	initMongoClient()
	initRedisClient()
	// indexView := mdbClient["test"].Collection("user_play_card").Indexes()
	// c, _ := indexView.List(context.Background())
	// var result []bson.M
	// c.All(context.Background(), &result)
	// for _, el := range result {
	// 	for k, v := range el {
	// 		fmt.Printf("%v: %v\n", k, v)
	// 	}
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	http.Handle("/home", gziphandler.GzipHandler(http.HandlerFunc(httpHomeHandler)))

	http.Handle("/home/redis", gziphandler.GzipHandler(http.HandlerFunc(httpRedisHandler)))
	http.Handle("/home/redis/list", gziphandler.GzipHandler(http.HandlerFunc(httpListHandler)))
	http.Handle("/home/redis/search", gziphandler.GzipHandler(http.HandlerFunc(httpSearchHandler)))
	http.Handle("/home/redis/detail", gziphandler.GzipHandler(http.HandlerFunc(httpDetailHandler)))

	http.Handle("/home/mongo", gziphandler.GzipHandler(http.HandlerFunc(httpMongoHandler)))
	http.Handle("/home/mongo/collections", gziphandler.GzipHandler(http.HandlerFunc(httpCollectionsHandler)))
	http.Handle("/home/mongo/indexes", gziphandler.GzipHandler(http.HandlerFunc(httpIndexesHandler)))

	port := ":7000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}

package main

import (
	"fmt"
	"gin-first/platform/newsFeed"
	"gin-first/server/controllers"
	db_database "gin-first/server/db"
	"gin-first/server/helpers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	ips := helpers.GetIP()

	fmt.Println("main ---", ips)

	feed := newsFeed.New()

	db, err := db_database.Connect()
	if err != nil {
		log.Fatal(fmt.Printf("Error connecting: %v \n", err))
		return
	}
	defer db.Close()

	// user := &db_database.User{
	// 	Email:    "vasia@gmail.com",
	// 	Password: "11111111",
	// }

	// user.Create()

	// fmt.Println("User", user)

	// uData, err := user.Get()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("User data", uData)

	// transport, err := tlsTcpConn()

	// if err != nil {
	// 	fmt.Println(fmt.Printf("Error connecting: %v \n", err))
	// 	// log.Fatal(fmt.Printf("Error connecting: %v \n", err))
	// }

	// fmt.Println("transport connected ---", transport)

	router := gin.Default()

	router.GET("/ping", controllers.PingGet())
	router.GET("/newsfeed", controllers.NewsFeedGet(feed))
	router.POST("/newsfeed", controllers.NewsFeedPost(feed))
	router.POST("/pvbki_auth/token", controllers.GetToken())
	router.POST("/idsub/token", controllers.GetTokenIdsub())

	router.GET("/v1/health/service/addsvc", controllers.HealthResponceAddsvc())

	router.GET("/v1/health/service/stringsvc", controllers.HealthResponceStringsvc())

	router.GET("/api/v1/health/service/addsvc", controllers.HealthResponceAddsvc())
	router.GET("/api/v1/user/logout", controllers.HealthResponceAddsvc())

	router.Run() // get port from PORT env if not exist listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// interface {}

// func WithtlsTcpConn(handler func(*tls.Conn) gin.HandlerFunc) (gin.HandlerFunc) {
// 	log.SetFlags(log.Lshortfile)

// 	conf := &tls.Config{
// 		InsecureSkipVerify: true,
// 	}

// 	conn, err := tls.Dial("tcp", "127.0.0.1:4444", conf)
// 	if err != nil {
// 		log.Println("tls connection error", err)
// 		// log.Fatal("tls connection error", err)

// 		var tr gin.HandlerFunc

// 		return tr
// 	}
// 	// defer conn.Close()

// 	// n, err := conn.Write([]byte("hello\n"))
// 	// if err != nil {
// 	// 	log.Println(n, err)
// 	// 	return
// 	// }

// 	// buf := make([]byte, 100)
// 	// n, err = conn.Read(buf)
// 	// if err != nil {
// 	// 	log.Println(n, err)
// 	// 	return
// 	// }

// 	// println(string(buf[:n]))

// 	return handler(conn)
// }

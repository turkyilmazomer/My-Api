package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`..\config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

/*
 * Tag... - a very simple struct
 */
type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {

	fmt.Println("Go MySQL Tutorial")

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := sql.Open("mysql", connection)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))

	//results, err := dbConn.Query("SELECT id, title FROM article")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}

	//for results.Next() {
	//	var tag Tag
	// for each row, scan the result into our tag composite object
	//	err = results.Scan(&tag.ID, &tag.Title)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	// and then print out the tag's Name attribute
	//	log.Printf(tag.Title)
	//}

}

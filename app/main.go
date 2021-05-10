package main

import (
	"log"

	"github.com/iswanulumam/dependency-injection-go/business/following"
	"github.com/iswanulumam/dependency-injection-go/db"
)

var (
	mysqlHost     = "localhost"
	mysqlUser     = "root"
	mysqlPassword = "admin123"
)

func main() {
	// initialize database connection
	dbCon, err := db.NewSql(mysqlHost, mysqlUser, mysqlPassword)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	// initite following repository
	followingRepo := following.NewMysqlRepository(dbCon)

	// initiate following service
	followingService := following.NewService(followingRepo)

	if err := followingService.Insert("root xixixix", "I'm Root! xixiix"); err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Println("Done!")
}

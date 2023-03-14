package db

import (
	"boardgame/app/config/toml"
	"boardgame/ent"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MySQL *ent.Client

func InitMySQL(c *toml.Config) {

	mySqlUrl := c.Local.User + ":" + c.Local.Password + "@" + c.Local.Host + "/" + c.Local.Db + "?parseTime=True"

	client, err := ent.Open(c.Local.Dbms, mySqlUrl)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	MySQL = client

	fmt.Println(c.Local.Dbms, " connected successfully !!")
}

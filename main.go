package main

import (
	"fmt"
	"os"

	"github.com/MarkTBSS/040-connectDatabase/config"
	"github.com/MarkTBSS/040-connectDatabase/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	db := databases.DbConnect(cfg.Db())
	fmt.Println(db)
}

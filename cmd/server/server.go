package main

import (
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/jeotic/nmapper/pkg/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := server.Server{}
	s.Init()
	s.Run(":8000")

	defer s.Env.DB.Close()
}

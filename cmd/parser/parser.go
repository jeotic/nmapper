package main

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/jeotic/nmapper/pkg"
	"github.com/jeotic/nmapper/pkg/parser"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	f, err := os.Open("nmap.results.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	db, err := sql.Open("sqlite3", "./dev-database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	dialect := goqu.Dialect("sqlite3")

	env := pkg.ENV{
		DB:      db,
		Builder: dialect,
	}

	nmap_id, err := parser.ParseReader(env, f)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(nmap_id)
}

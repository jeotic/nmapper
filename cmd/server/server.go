package main

import (
	"database/sql"
	"flag"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/gorilla/mux"
	"github.com/jeotic/nmapper/pkg"
	"github.com/jeotic/nmapper/pkg/router"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "Javascript files")
	flag.Parse()

	muxRouter := *mux.NewRouter().StrictSlash(true)
	muxRouter.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(dir))))

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

	router.AddRoutes(env, &muxRouter)

	log.Fatal(http.ListenAndServe(":3000", &muxRouter))
}

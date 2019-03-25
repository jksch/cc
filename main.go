package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jksch/cc/internal/db"
	"github.com/jksch/cc/internal/server"
)

var dbPath = flag.String("db_path", "app.db", "Path to app db.")
var srvAddr = flag.String("addr", server.DefaultAddr, "Address on which the server will listen.")
var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Printf("Error on %v", run())
}

func run() error {
	flag.Parse()

	var err error
	var DB *db.Database
	var srv *server.Server

	setup := []func() (description string, err error){
		func() (string, error) {
			DB, err = db.New(*dbPath)
			return "database", err
		},
		func() (string, error) {
			srv, err = server.New(*srvAddr, DB, logger)
			return "server", err
		},
	}
	for _, do := range setup {
		if des, err := do(); err != nil {
			return fmt.Errorf("app %s setup, %+v", des, err)
		}
	}

	return fmt.Errorf("app run, %v", srv.Run())
}

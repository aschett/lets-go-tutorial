package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	// Apparently this is common usage for sql drivers in go to surpress compile errors because the package is not even used but only specific function
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
)

type application struct {
	logger *slog.Logger
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	addr := flag.String("addr", ":4000", "Http network address")
	//asks for the mysql user assumes standard is web as used in the tutorial.
	//The following lines could be kinda messy as in the tutorial he hardcoded the passcode which doesn't seem very "safe"
	user := flag.String("dbuser", "web", "MySQL User")
	dbname := flag.String("dbname", "snippetbox", "MySQL Database Name")
	//Add new flag for mysql
	flag.Parse()

	fmt.Printf("Enter your MySQL password for User %s for Database %s:", *user, *dbname)
	//Load password securely with term
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	password := string(passwordBytes)

	//Building the dsn dynamically
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", *user, password, *dbname)

	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
	}

	logger.Info("Starting Server on", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

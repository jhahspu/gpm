package main

import (
	"context"
	"database/sql"
	"errors"
	"gpm/models"
	"gpm/repository"
	"io/fs"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx     context.Context
	sqlRepo repository.SqliteRepository
}

// NewApp creates a new App application struct
func NewApp() *App {
	// check if DB exists
	// if no, create DB
	if _, err := os.Stat("./gpm.db"); errors.Is(err, fs.ErrNotExist) {
		// message DB doesn't exist
		log.Println(err.Error())

		// create DB
		file, err := os.Create("gpm.db")
		if err != nil {
			log.Println(err.Error())
		}
		file.Close()
		log.Println("----- DB created")
	} else {
		log.Println("----- DB exists")
	}
	db, err := sql.Open("sqlite3", "./gpm.db")
	if err != nil {
		log.Println("----- error opening DB")
	}
	// defer db.Close()

	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS gpm (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"website" TEXT,
		"username" TEXT,
		"email" TEXT,
		"password" TEXT,
		"notes" TEXT
	  );` // SQL Statement for Create Table

	log.Println("----- Creating table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("----- Table created")
	repo, err := repository.NewSqliteRepository(db)
	if err != nil {
		log.Println("---- error plug-in sqlite repo ")
	}

	return &App{
		sqlRepo: repo,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// return all rows
func (a *App) HandleGetGPMs() models.AllRows {
	var response models.AllRows

	gpms, err := a.sqlRepo.DisplayGPMs()
	if err != nil {
		response.Status = err.Error()
		return response
	}

	response.Status = "success"
	response.Rows = gpms

	return response
}

// insert new row
func (a *App) HandleInsertGPM(req models.ReqInsertGPM) string {
	log.Println("----- got data")
	log.Println(req)

	var newRow models.GPM

	newRow.Website = req.Website
	newRow.Username = req.Username
	newRow.Email = req.Email
	newRow.Password = req.Password
	newRow.Notes = req.Notes

	err := a.sqlRepo.InsertGPM(newRow)
	if err != nil {
		return err.Error()
	}

	return "added"
}

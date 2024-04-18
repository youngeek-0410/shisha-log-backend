package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"shisha-log-backend/lib"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var migrationFilePath = "file://./migrations/"

func main() {
	fmt.Println("start migration")
	flag.Parse()
	command := flag.Arg(0)

	if command == "" {
		showUsage()
		os.Exit(1)
	}

	m, err := migrate.New(migrationFilePath, lib.GenerateDsn())
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "migrate":
		migration(m)
	case "up":
		up(m)
	case "down":
		down(m)
	case "drop":
		drop(m)
	case "version":
		showVersionInfo(m.Version())
	default:
		fmt.Println("\nerror: invalid command '", command, "'")
		showUsage()
		os.Exit(0)
	}
}

func showUsage() {
	fmt.Println(`
-------------------------------------
Usage:
  go run migration/main.go <command>
Commands:
  migrate   Apply all up migrations
  up        Apply up migrations
  down      Apply down migrations
  drop      Drop everything
  version   Check current migrate version
-------------------------------------`)
}

func migration(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Up()
	if err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
		fmt.Println("\nno change")
	} else {
		fmt.Println("\nUpdated:")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func up(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Steps(1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

func down(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Steps(-1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

func drop(m *migrate.Migrate) {
	err := m.Drop()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Dropped all migrations")
		return
	}
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version : ", version)
	fmt.Println("dirty   : ", dirty)
	fmt.Println("error   : ", err)
	fmt.Println("-------------------")
}

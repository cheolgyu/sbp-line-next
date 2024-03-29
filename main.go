package main

import (
	"log"

	"github.com/cheolgyu/sbm-base/db"
	"github.com/cheolgyu/sbm-base/logging"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}
func project_run() {
	run()
}

func run() {
	query := `select 1 from project.func_lines()`

	_, err := db.Conn.Exec(query)
	if err != nil {
		log.Fatalln(err, query)
		panic(err)
	}
}

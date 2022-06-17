package main

import (
	"flag"
	"fmt"

	"github.com/jtheo/boot_dev_go_project_sm/internal/http"
)

func main() {
	bind := flag.String("addr", "localhost", "host to send the requests")
	port := flag.Int("port", 8080, "port to send the requests")
	db := flag.String("db", "./db.json", "path to json file")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *bind, *port)

	http.New(addr, *db)
}

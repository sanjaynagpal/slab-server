package main

import (
	"os"

	"github.com/sanjaynagpal/slab-server/internal/server"
)

func main() {
	os.Exit(server.Main())
}

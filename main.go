package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	switch args[1] {
	case List.String():
		handleList(args)
	case Add.String():
		handleAdd(args)
	case Delete.String():
		handleDelete(args)
	case Update.String():
		handleUpdate(args)
	case MarkInProgress.String():
		handleMarkInProgress(args)
	case MarkDone.String():
		handleMarkDone(args)
	default:
		log.Fatalf("unknown command: %s", args[1])
	}
}

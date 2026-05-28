package main

import (
	"log"
	"strconv"
)

func handleList(args []string) {
	s, err := ParseStatus(args[2])

	if err != nil {

		log.Fatalf("%v. use: %s,%s,%s", err, Todo, InProgress, Done)
	}

	tasks := getTasksByStatus(s)
	log.Printf("%+v", tasks)
}

func handleAdd(args []string) {
	addTask(TaskDescription(args[2]))
}

func handleDelete(args []string) {
	id, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal("Please provide correct id format (int)")
	}

	deleteTask(TaskID(id))
}

func handleMarkInProgress(args []string) {
	id, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal("Please provide correct id format (int)")
	}

	updateTask(
		TaskID(id),
		UpdateDto{status: InProgress.Pointer()},
	)
}

func handleMarkDone(args []string) {
	id, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal("Please provide correct id format (int)")
	}

	updateTask(
		TaskID(id),
		UpdateDto{status: Done.Pointer()},
	)
}

func handleUpdate(args []string) {
	id, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal("Please provide correct id format (int)")
	}

	if args[3] != "" {
		updateTask(
			TaskID(id),
			UpdateDto{
				description: TaskDescription(args[3]).TaskDescriptionPointer(),
			})
	} else {
		log.Fatal("Please provide description")
	}
}

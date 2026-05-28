package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) == 2 && args[1] == List.String() {
		tasks := getTasks()
		log.Printf("%+v", tasks)
	}

	if len(args) == 3 {
		if args[1] == List.String() {
			s, err := ParseStatus(args[2])

			if err != nil {
				log.Fatalf("%v. use: %s,%s,%s", err, Todo, InProgress, Done)
			}

			tasks := getTasksByStatus(s)
			log.Printf("%+v", tasks)

			return
		}

		if args[1] == Add.String() && args[2] != "" {
			addTask(TaskDescription(args[2]))

			return
		}

		if args[2] != "" {
			id, err := strconv.Atoi(args[2])

			if err != nil {
				log.Fatal("Please provide correct id format (int)")
			}

			if args[1] == Delete.String() {
				deleteTask(TaskID(id))
			}

			if args[1] == MarkInProgress.String() {
				updateTask(
					TaskID(id),
					UpdateDto{status: InProgress.Pointer()},
				)
			}

			if args[1] == MarkDone.String() {
				updateTask(
					TaskID(id),
					UpdateDto{status: Done.Pointer()},
				)
			}

		}

		return
	}

	if len(args) == 4 && args[1] == Update.String() {
		id, err := strconv.Atoi(args[2])

		if err != nil {
			log.Fatal("Please provide correct id format (int)")
		}

		if args[3] != "" {
			updateTask(
				TaskID(id),
				UpdateDto{
					description: new(TaskDescription(args[3])),
				})
		} else {
			log.Fatal("Please provide description")
		}

	}

	log.Printf("%+v", getTasks())
}

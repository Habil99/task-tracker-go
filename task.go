package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"time"
)

type (
	TaskId          int
	TaskDescription string

	task struct {
		Id          TaskId
		Description TaskDescription
		Status      Status
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	UpdateDto struct {
		description *TaskDescription
		status      *Status
	}
)

func addTask(d TaskDescription) {
	var tasks []task = getTasks()

	id := getIncrementalId(tasks)
	now := time.Now()

	t := task{
		Id:          id,
		Description: d,
		Status:      Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	newTasks := append(tasks, t)

	err := updateJSONFile(newTasks)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Printf("Task added successfully (ID: %d)", t.Id)
}

func getTasks() []task {
	file, err := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error while opening/creating json ", err)
	}

	defer file.Close()

	var todos []task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		if err == io.EOF {
			return []task{}
		}

		log.Fatal("Error while decoding file: ", err)
	}

	return todos
}

func getTasksByStatus(s Status) []task {
	tasks := getTasks()
	filtered := make([]task, 0)

	for _, t := range tasks {
		if t.Status == s {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func getTaskById(id TaskId) (task, int, error) {
	tasks := getTasks()

	for i, task := range tasks {
		if task.Id == id {
			return task, i, nil
		}
	}

	return task{}, 0, fmt.Errorf("task not found by ID: %d", id)
}

func deleteTask(id TaskId) {
	tasks := getTasks()
	_, taskIndex, err := getTaskById(id)

	if err != nil {
		return
	}

	updatedTasks := slices.Delete(tasks, taskIndex, taskIndex+1)

	updateErr := updateJSONFile(updatedTasks)

	if updateErr != nil {
		log.Fatal("Error while updating json ", updateErr)
	}

	log.Printf("Task deleted successfully (ID: %d)", id)
}

func updateTask(id TaskId, dto UpdateDto) {
	_, ti, err := getTaskById(id)

	if err != nil {
		log.Fatal("Error finding task by id ", err)
	}

	tasks := getTasks()

	if dto.description != nil {
		tasks[ti].Description = *dto.description
	}

	if dto.status != nil {
		tasks[ti].Status = *dto.status
	}

	tasks[ti].UpdatedAt = time.Now()

	ue := updateJSONFile(tasks)

	if ue != nil {
		log.Fatal("Error happened updating json file ", ue)
	}

	log.Printf("Task updated successfully (ID: %d)", id)
}

func updateJSONFile(tasks []task) error {
	file, oe := os.OpenFile("todo.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if oe != nil {
		return oe
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	ee := encoder.Encode(tasks)

	return ee
}

func getIncrementalId(tasks []task) TaskId {
	var bi TaskId

	for _, t := range tasks {
		if bi < t.Id {
			bi = t.Id
		}
	}

	return bi + 1
}

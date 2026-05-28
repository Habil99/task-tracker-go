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
	TaskID          int
	TaskDescription string

	task struct {
		ID          TaskID          `json:"id"`
		Description TaskDescription `json:"description"`
		Status      Status          `json:"status"`
		CreatedAt   time.Time       `json:"createdAt"`
		UpdatedAt   time.Time       `json:"updatedAt"`
	}

	UpdateDto struct {
		description *TaskDescription
		status      *Status
	}
)

func (td TaskDescription) TaskDescriptionPointer() *TaskDescription {
	return &td
}

func addTask(d TaskDescription) {
	tasks, _ := getTasks()

	id := getIncrementalId(tasks)
	now := time.Now()

	t := task{
		ID:          id,
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

	log.Printf("Task added successfully (ID: %d)", t.ID)
}

func getTasks() ([]task, error) {
	file, err := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error while opening/creating json ", err)
	}

	defer file.Close()

	var todos []task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		if err == io.EOF {
			return []task{}, nil
		}

		return []task{}, fmt.Errorf("Error while decoding file: ", err)
	}

	return todos, nil
}

func getTasksByStatus(s Status) []task {
	tasks, _ := getTasks()
	filtered := make([]task, 0)

	for _, t := range tasks {
		if t.Status == s {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func getTaskById(id TaskID) (task, int, error) {
	tasks, _ := getTasks()

	for i, t := range tasks {
		if t.ID == id {
			return t, i, nil
		}
	}

	return task{}, 0, fmt.Errorf("task not found by ID: %d", id)
}

func deleteTask(id TaskID) {
	tasks, _ := getTasks()
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

func updateTask(id TaskID, dto UpdateDto) {
	_, ti, err := getTaskById(id)

	if err != nil {
		log.Fatal("Error finding task by id ", err)
	}

	tasks, _ := getTasks()

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

func getIncrementalId(tasks []task) TaskID {
	var bi TaskID

	for _, t := range tasks {
		if bi < t.ID {
			bi = t.ID
		}
	}

	return bi + 1
}

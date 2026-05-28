package main

type Command int

const (
	Add Command = iota
	Update
	Delete
	List
	MarkDone
	MarkInProgress
)

func (c Command) String() string {
	return [...]string{"add", "update", "delete", "list", "mark-done", "mark-in-progress"}[c]
}

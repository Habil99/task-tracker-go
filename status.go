package main

import (
	"encoding/json"
	"fmt"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	return [...]string{"todo", "in-progress", "done"}[s]
}

func ParseStatus(v string) (Status, error) {
	switch v {
	case Todo.String():
		return Todo, nil
	case InProgress.String():
		return InProgress, nil
	case Done.String():
		return Done, nil
	default:
		return 0, fmt.Errorf("invalid status %q", v)
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("status must be a string: %w", err)
	}

	switch raw {
	case "todo":
		*s = Todo
	case "in-progress":
		*s = InProgress
	case "done":
		*s = Done
	default:
		return fmt.Errorf("invalid status %s", raw)
	}

	return nil
}

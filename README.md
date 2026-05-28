# Task Tracker CLI (Go)

A command-line task manager built with Go that supports creating, listing, updating, deleting, and status-tracking tasks persisted in a local JSON file.

This is a portfolio project focused on learning practical Go fundamentals: CLI workflows, type modeling, JSON serialization, and clean project organization.

## Highlights

- CRUD operations from the terminal
- Task status lifecycle: `todo`, `in-progress`, `done`
- Status filtering (`list <status>`)
- Persistent storage in `todo.json`
- Incremental task IDs
- Split command handlers (`handlers.go`) for maintainability

## Tech Stack

- Go
- Standard library only (`os`, `encoding/json`, `time`, etc.)

## Project Structure

```text
task-tracker/
  main.go       # command dispatch
  handlers.go   # CLI handlers per command
  command.go    # command enum/string mapping
  task.go       # task model + service + JSON persistence
  status.go     # status type + parsing + JSON marshal/unmarshal
  todo.json     # local data storage
```

## Getting Started

### Prerequisites

- Go 1.22+

### Run

From the `task-tracker` directory:

```bash
go run .
```

## Commands

```bash
go run . add "Buy groceries"
go run . list
go run . list todo
go run . list in-progress
go run . list done
go run . update 1 "Buy groceries and cook dinner"
go run . mark-in-progress 1
go run . mark-done 1
go run . delete 1
```

## Example JSON Output

```json
[
  {
    "id": 1,
    "description": "Buy groceries",
    "status": "todo",
    "createdAt": "2026-05-28T12:00:00Z",
    "updatedAt": "2026-05-28T12:00:00Z"
  }
]
```

## What I Learned

- Designing small CLI applications in Go
- Safely reading/writing JSON files
- Using custom types (`TaskID`, `TaskDescription`, `Status`)
- Modeling optional updates with pointers
- Refactoring into handlers as complexity grows

## Current Limitations

- No automated tests yet
- User-facing output is still basic (`log.Printf`)
- Errors are mostly handled via `log.Fatal` rather than propagated

## Next Improvements

- Add unit tests for status parsing and task operations
- Improve CLI UX with `help` and formatted output
- Move to a clearer package layout as the project grows
- Optionally switch storage from JSON file to SQLite

## License

MIT

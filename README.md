# ToGo - CLI Todo Application

A simple and efficient command-line todo application written in Go. ToGo helps you manage your tasks directly from the terminal without switching context.

## Features

- Add new tasks
- Edit existing tasks
- Delete tasks
- Toggle task completion status
- List all tasks with formatting
- Persistent storage using JSON

## Installation

### Prerequisites
- Go 1.21 or higher

### Building from source
```bash
# Clone the repository
git clone https://github.com/istiaq-fuad/togo.git
cd togo

# Build the application
go build -o togo

# Make it available system-wide (optional)
mv togo /usr/local/bin/
```

## Usage

### Adding a new task
```bash
./togo -add "Complete the project documentation"
```

### Editing a task
```bash
./togo -edit "0:Updated task description"
```

### Deleting a task
```bash
./togo -del 2
```

### Toggling completion status
```bash
./togo -toggle 1
```

### Listing all tasks
```bash
./togo -list
```

## Project Structure

- `main.go` - Entry point of the application
- `todo.go` - Todo data structures and methods
- `command.go` - Command line flag processing
- `storage.go` - Data persistence handling

## Data Storage

All tasks are stored in a JSON file (`todos.json`) in the current directory.

## Dependencies

- [github.com/aquasecurity/table](https://github.com/aquasecurity/table) - For formatted table output

## License

MIT

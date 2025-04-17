package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

// Task represents a single todo task
type Task struct {
    ID        int    `json:"id"`        // Task identifier
    Name      string `json:"name"`      // Task name or description
    Completed bool   `json:"completed"` // Completion status
}

const fileName = "tasks.json" // File to store tasks

// loadTasks loads tasks from the JSON file
func loadTasks() ([]Task, error) {
    var tasks []Task

    data, err := ioutil.ReadFile(fileName)
    if os.IsNotExist(err) {
        return tasks, nil // Return empty list if file doesn't exist yet
    } else if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &tasks)
    return tasks, err
}

// saveTasks writes tasks to the JSON file
func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(fileName, data, 0644)
}

// addTask adds a new task with the given name
func addTask(name string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    newTask := Task{
        ID:        len(tasks) + 1,
        Name:      name,
        Completed: false,
    }

    tasks = append(tasks, newTask)

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving task:", err)
        return
    }

    fmt.Println("Task added:", name)
}

// listTasks displays all tasks with their status
func listTasks() {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    for _, task := range tasks {
        status := "[ ]"
        if task.Completed {
            status = "[x]"
        }
        fmt.Printf("%d. %s %s\n", task.ID, status, task.Name)
    }
}

// main handles command-line arguments and routes to commands
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go [add|list] [task name]")
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Please provide a task name.")
            return
        }
        taskName := os.Args[2]
        addTask(taskName)
    case "list":
        listTasks()
    default:
        fmt.Println("Unknown command:", command)
    }
}

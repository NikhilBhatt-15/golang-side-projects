package main

import (
	"errors"
	"fmt"
)

type Task struct {
    ID          int
    Name        string
    Description string
    IsDone      bool
}

var taskList []Task
var idCounter = 1

func main() {
    create("Buy groceries", "Milk, Bread, Eggs")
    create("Read book", "Clean Code")
    listTasks()

    fmt.Println("\nMarking task 1 as done:")
    taskList[0].markDone()
    listTasks()

    fmt.Println("\nDeleting task with ID 2:")
    err := deleteTask(2)
    if err != nil {
        fmt.Println("Error:", err)
    }
    listTasks()
	taskList[0].print()
}

func create(title string, description string) *Task {
    newtask := Task{
        ID:          idCounter,
        Name:        title,
        Description: description,
        IsDone:      false,
    }
    taskList = append(taskList, newtask)
    idCounter++
    fmt.Println("Added the task:", title)
    return &newtask
}

func (t *Task) markDone() {
    t.IsDone = true
}

func deleteTask(ID int) error {
    for i, t := range taskList {
        if t.ID == ID {
            taskList = append(taskList[:i], taskList[i+1:]...)
            fmt.Printf("Deleted task with ID %d\n", ID)
            return nil
        }
    }
    return errors.New("task not found")
}

func listTasks() {
    if len(taskList) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    for _, t := range taskList {
        status := "❌"
        if t.IsDone {
            status = "✅"
        }
        fmt.Printf("[%d] %s - %s (%s)\n", t.ID, t.Name, t.Description, status)
    }
}

func (t *Task) print(){
	fmt.Println("------Task-----")
	fmt.Println("Title: ",t.Name)
	fmt.Println("Description: ",t.Description)
	if t.IsDone{
		fmt.Println("Task is Completed")
	}else{
		fmt.Println("Task is Pending")
	}
}
package main

import (
	"fmt"
	"os"
)


type TaskStatus int

// TaskStatus enum.
const(
	NotStarted TaskStatus = iota
	InProgress
	Completed
)

// Map TaskStatus String.
var statuses= map[TaskStatus]string{
	NotStarted: "Not Started",
	InProgress: "In Progress",
	Completed: "Completed",
}

// Task struct.
type Task struct{
	Title string
	Description string
	Status TaskStatus
}

type TaskManager struct{
	Tasks map[string]Task
}

// statusString to convert int enum to string.
func (ts TaskStatus)statusString() string {
	if status,found:=statuses[ts];found{
		return status
	}
	return "Unknown"
}

// addTask to add a task to taskmanager.
func (t *TaskManager)addTask(title,description string){
		t.Tasks[title]= Task{Title: title,Description: description,Status: TaskStatus(NotStarted)}
		fmt.Println("Task  Added successfully")
}

// deleteTask to delete a task from taskmanager.
func (t *TaskManager)deleteTask(title string)  {
	if _,found:= t.Tasks[title];found{
		delete(t.Tasks,title)
		fmt.Println("Task deleted")
	}else{
		fmt.Println("Task not found to delete")
	}
}

// viewTask to view a task from taskmanager.
func (t *TaskManager)viewTask(title string)  {
	if task,found:= t.Tasks[title];found{
		fmt.Printf("Task Title: %s, Description: %s, Status: %s\n",task.Title,task.Description,task.Status.statusString())
	}else{
		fmt.Println("Task not found to view")
	}
}

// listTasks to list all tasks.
func (t *TaskManager)listTasks(){
	tasks:= t.Tasks
	if len(tasks) == 0{
		fmt.Println("No tasks found")
		return
	}
	for _, task := range t.Tasks {
		fmt.Printf("Title: %s, Description: %s, Status: %s\n", task.Title, task.Description,task.Status.statusString())
	}
}

// MarkTaskStatus marks the task status. 
func (t *TaskManager)MarkTaskStatus(title string,status TaskStatus)  {
	if val,found:=t.Tasks[title];found{
		val.Status = TaskStatus(status)
		t.Tasks[title] = val
		fmt.Printf("Task is marked %s\n",val.Status.statusString())
	}else{
		fmt.Println("No task found to change status")
	}
}
func main() {

  tasks:=  make(map[string]Task)	
  taskManager := TaskManager{
	Tasks: tasks,
  }
	
	for {
		fmt.Printf("Hello !! Welcome to task management app\n")
		fmt.Println("Enter 1 to add a task")
		fmt.Println("Enter 2 to view lists of tasks")
		fmt.Println("Enter 3 to view task")
		fmt.Println("Enter 4 to delete a task")
		fmt.Println("Enter 5 to mark task as inprogress")
		fmt.Println("Enter 6 to mark task as  completed")
		fmt.Println("Enter 7 to exit")

		fmt.Println("Enter choice between 1 to 6")
		var choice int
		fmt.Scanln(&choice)

		switch choice{
		case 1:
			var title,description string
			fmt.Println("Enter task title")
			_,err:=fmt.Scanln(&title)
			if err!=nil{
				fmt.Println("Error reading title",err)
				return
			}
			fmt.Println("Enter task description")
			_,err=fmt.Scanln(&description)
			if err!=nil{
				fmt.Println("Error reading description",err)
				return
			}
			taskManager.addTask(title,description)
		case 2:
			taskManager.listTasks()
		case 3:
			var title string
			fmt.Println("Enter task title")
			_,err:=fmt.Scanln(&title)
			if err!=nil{
				fmt.Println("Error reading title",err)
				return
			}
			taskManager.viewTask(title)
		case 4:
			var title string
			fmt.Println("Enter task title")
			_,err:=fmt.Scanln(&title)
			if err!=nil{
				fmt.Println("Error reading title",err)
				return
			}
			taskManager.deleteTask(title)
		case 5:
			var title string
			fmt.Println("Enter task title")
			_,err:=fmt.Scanln(&title)
			if err!=nil{
				fmt.Println("Error reading title",err)
				return
			}
			taskManager.MarkTaskStatus(title,InProgress)
		case 6:
			var title string
			fmt.Println("Enter task title")
			_,err:=fmt.Scanln(&title)
			if err!=nil{
				fmt.Println("Error reading title",err)
				return
			}
			taskManager.MarkTaskStatus(title,Completed)	
		case 7:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
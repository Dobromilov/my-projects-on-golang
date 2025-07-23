package taskmanager

import (
	"fmt"
)

type Tasks struct {
    Title string
    Done  bool
}

var TasksList []Tasks

func AddTask(tasks *[]Tasks, title string){
	*tasks = append(*tasks, Tasks{Title: title, Done: false})
	fmt.Println("Задача добавлена")
}

func ViewTasks(tasks []Tasks){
	if len(tasks) == 0 {
		fmt.Println("📭 Список задач пуст")
		return
	}

	fmt.Println("\n📋 Список задач:")
	for i, task := range tasks {
		status := "X"
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
}

func CompleteTask(tasks *[]Tasks, index int){
	if index < 1 || index > len(*tasks) {
		fmt.Printf("⚠️ Неверный номер задачи! Допустимый диапазон: 1-%d\n", len(*tasks))
		return
	}

	if (index-1 >= 0 && index <= len(*tasks)){
		(*tasks)[index-1].Done = true	
	}
}	
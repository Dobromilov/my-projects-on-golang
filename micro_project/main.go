package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	
    "micro_project/internal/storage"
    "micro_project/internal/taskmanager"
)

func main() {
	if err := storage.LoadTasks(); err != nil {
		log.Printf("⚠️ Ошибка загрузки задач: %v", err)
	}

	// var choice int
	// fmt.Scanln(&choice) - оказывается так тоже можно было, но до этого учась по книжке я делал по другому, поэтому закроем на это глаза

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Посмотреть задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Сохранить и выйти")
		fmt.Print("-> ")

		choice,err := reader.ReadString('\n')

		if (err != nil){
			log.Fatal(err)
		} else {

			choice, err := strconv.Atoi(strings.TrimSpace(choice))
			
			if err != nil {
				log.Fatal(err)
			}else{
				switch choice {
				case 1:
					title, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					taskmanager.AddTask(&taskmanager.TasksList, strings.TrimSpace(title))
				case 2:
					taskmanager.ViewTasks(taskmanager.TasksList)
				case 3:
					taskmanager.ViewTasks(taskmanager.TasksList)
					index, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					} else {
						index, err := strconv.Atoi(strings.TrimSpace(index))
						if err != nil {
							log.Fatal(err)
						} else {
							taskmanager.CompleteTask(&taskmanager.TasksList, index)
						}
					}
				case 4:
					if err := storage.SaveTasks(); err != nil {
						log.Println("⚠️ Ошибка сохранения:", err)
					} else {
						fmt.Println("💾 Задачи сохранены. До свидания!")
					}
					return
				default:
					fmt.Println("Неизвестное действие")
				}
			}
		}
	}
}
package storage

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"micro_project/internal/taskmanager"
)

func SaveTasks() error {
	file, err := os.Create("tasks.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, task := range taskmanager.TasksList {
		_, err := fmt.Fprintf(file, "%s | %t\n", task.Title, task.Done)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadTasks() error {
	file, err := os.Open("tasks.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return nil 
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue 
		}
		
		done, err := strconv.ParseBool(strings.TrimSpace(parts[1]))
		if err != nil {
			continue 
		}
		
		taskmanager.TasksList = append(taskmanager.TasksList, taskmanager.Tasks{
			Title: strings.TrimSpace(parts[0]),
			Done:  done,
		})
	}
	return scanner.Err()
}
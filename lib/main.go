package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
	"strings"

	"lib/models"
	"lib/services"
)

func main(){
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1) Добавить книгу в библиотеку")
		fmt.Println("2) Вывести список книг")
		fmt.Println("3) Удалить из списка книгу по ID")
		fmt.Println("4) Очистить библиотеку")
		fmt.Println("5) Выход")
		fmt.Print("-->")	
		
		reader := bufio.NewReader(os.Stdin)
		choiceStr, err := reader.ReadString('\n')

		if err!=nil{
			log.Fatal(err)
		} else {
			choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))

			if err!=nil{
				log.Fatal(err)
			} else {
				switch choice {
				case 1:
					services.AddBook()
				case 2:
					services.PrintLib()
				case 3:
					if len(models.Lib) == 0 {
						fmt.Println("Библиотека пуста.")
					} else {
						fmt.Print("Введите ID книги для удаления: ")
						idStr, err := reader.ReadString('\n')
						if err != nil {
							log.Fatal(err)
						} else {
							idToDelete, err := strconv.Atoi(strings.TrimSpace(idStr))
							if err != nil {
								log.Fatal(err)
							} else {
								services.DeleteBook(idToDelete)
							}
						}
					}
				case 4:
					services.DeleteEverything()
				case 5:
					return
				default:
					fmt.Println("Неизвестная команда!")
				}
			}
		}
	}
}
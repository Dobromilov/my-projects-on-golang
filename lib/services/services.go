package services

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

	"lib/models"
)

func AddBook(){
	reader := bufio.NewReader(os.Stdin)
	id_book := models.Id
	models.Id += 1

	fmt.Print("Введите название: ")
	title_book,_ := reader.ReadString('\n')
	title_book = strings.TrimSpace(title_book)

	fmt.Print("Введите автора: ")
	author_book,_ := reader.ReadString('\n')
	author_book = strings.TrimSpace(author_book)

	fmt.Print("Введите год: ")
	yearStr,_ := reader.ReadString('\n')
	year_book,_ := strconv.Atoi(strings.TrimSpace(yearStr))

	book := models.Book{
		ID: id_book,
		Title: title_book,
		Author: author_book,
		Year: year_book,
	}

	models.Lib = append(models.Lib, book)
	fmt.Println("Книга добавлена!")
}

func PrintLib(){
	if len(models.Lib) == 0 {
		fmt.Println("Библиотека пуста.")
		return
	}
	fmt.Printf("%-4s | %-40s | %-15s | %-4s\n", "ID", "Название", "Автор", "Год")
	fmt.Println(strings.Repeat("-", 75))
	for _, book := range models.Lib {
		fmt.Printf("%-4d | %-40s | %-15s | %-4d\n", book.ID, book.Title, book.Author, book.Year)
	}
}

func DeleteBook(id_book int){
	// for ind, book := models.Lib {
	// 	if id_book == ind {
	// 		models.Lib = append(models.Lib[:ind],models.Lib[i-1]...) // со срезами до этой програмки я не работал
	// 		fmt.Printf("Книга с ID %d удалена!\n", id_book)
    //         return
	// 	}
	// }
	// fmt.Printf("Книга с ID %d не найдена.\n", id_book)
	// не хочу этим способом, т.к. он достаточно долгий, хочу реализовать бин поиск

	left := 0
	right := len(models.Lib)-1

	for left <= right {
		mid := (left + right)/2
		if models.Lib[mid].ID == id_book {
			models.Lib = append(models.Lib[:mid], models.Lib[mid+1:]...)
			fmt.Printf("Книга с ID %d удалена!\n", id_book)
			return
		} else if models.Lib[mid].ID < id_book {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("Книга с ID %d не найдена.\n", id_book)
}

func DeleteEverything() {
    models.Lib = models.Lib[:0] // со срезами до этой програмки я не работал
    fmt.Println("Библиотека очищена!")
}
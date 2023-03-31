package Floors

import (
	"fmt"
	"log"
)

func FloorNum() (numfloor int, err error) {
	fmt.Println("Введите нужный этаж: ")
	_, err = fmt.Scan(&numfloor)
	if err != nil {
		log.Println("Ошибка ввода")
		return
	}
	return
}

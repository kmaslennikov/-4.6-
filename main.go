package main

import "fmt"

func main() {

	var N int             //кол-во студентов на курсе
	var K int             //кол-во групп
	var StudentNumber int //порядковый номер студента
	var groupNumber int   //номер группы
	fmt.Println("Введите количество студентов на курсе:")
	fmt.Scan(&N)
	fmt.Println("Введите количество групп:")
	fmt.Scan(&K)
	fmt.Println("Введите порядковый номер студента:")
	fmt.Scan(&StudentNumber)

	//количество студентов в группе
	count := N / K
	//определяем группу студента по его порядковому номеру и числу студентов в группе
	groupNumber = StudentNumber/count + 1

	//если значение граничное, то предыдущая группа
	if StudentNumber%count == 0 {
		groupNumber -= 1
	}

	fmt.Println(groupNumber)
}

package main

import (
	"fmt"
)

func CompVariant(num int) string {
	// Получаем последнюю цифру числа
	o := num % 10

	// Создаем словарь со склонениями слова "компьютер"
	s := map[int]string{
		1: "компьютер",
		2: "компьютера",
		3: "компьютера",
		4: "компьютера",
		5: "компьютеров",
		6: "компьютеров",
		7: "компьютеров",
		8: "компьютеров",
		9: "компьютеров",
		0: "компьютеров",
	}

	// Используем склонение из словаря и выводим string через Sprintf
	return fmt.Sprintf("%d %s", num, s[o])
}

func main() {
	// Можно было сделать без возврата, но пусть будет так
	fmt.Println(CompVariant(121))
	fmt.Println(CompVariant(132))
	fmt.Println(CompVariant(243))
	fmt.Println(CompVariant(354))
	fmt.Println(CompVariant(465))
	fmt.Println(CompVariant(576))
	fmt.Println(CompVariant(1048))
}

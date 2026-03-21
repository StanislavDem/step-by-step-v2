package actioninfo

import (
	"fmt"
)

type DataParser interface {
	// Объявление сигнатур
	Parse(datastring string) error
	ActionInfo() (string, error)

}

func Info(dataset []string, dp DataParser) {
	
	for _, data := range dataset {
		// парсим строку
		if err := dp.Parse(data); err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}

		// получаем строку с результатом
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка формирования информации:", err)
			continue
		}

		// выводим результат
		fmt.Println(info)
	}
}

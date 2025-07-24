package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Инициализация map
	mp1 := make(map[int]string)

	//Заполнение map
	for i := 0; i < 5; i++ {
		mp1[i] = strings.Repeat("a", i+1)
	}

	//Инициализация нового приложения
	app := fiber.New()

	//Get запрос
	app.Get("/get_value/:id", func(c *fiber.Ctx) error {

		//Вытаскивание ключа и преобразование в числовой тип
		idStr := c.Params("id")
		id := 0
		_, err := fmt.Sscan(idStr, &id)

		//Проверка на корректность ключа
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Incorrect key format!")
		}

		//Получение значения
		value, ok := mp1[id]

		//Вывод результата
		if !ok {
			return c.Status(fiber.StatusNotFound).SendString("Key not found!")
		}

		return c.SendString("Value: " + value)
	})
	//Запуск сервера
	log.Fatal(app.Listen(":3000"))
}

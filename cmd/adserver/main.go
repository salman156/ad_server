package main

import (
	"ad_server/handlers"
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
		return (handlers.GetFromMap(c))

	})
	//Запуск сервера
	log.Fatal(app.Listen(":3000"))
}

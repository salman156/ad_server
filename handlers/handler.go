package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateMap() map[int]string {
	// Инициализация map
	mp := make(map[int]string)

	// Заполнение map
	for i := 0; i < 5; i++ {
		mp[i] = strings.Repeat("a", i+1)
	}

	return mp
}

func GetFromMap(c *fiber.Ctx) error {
	//Вытаскивание ключа и преобразование в числовой тип
	mp1 := CreateMap()

	idStr := c.Params("id")
	id := 0
	_, err := fmt.Sscan(idStr, &id)

	//Проверка на корректность ключа
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Incorrect key format!")
	}

	//Получение значения
	value, ok := mp1[id]

	//Вывод результата
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("Key not found!")
	}

	return c.SendString("Value: " + value)
}

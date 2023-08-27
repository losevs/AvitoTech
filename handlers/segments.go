package handlers

import (
	"avitoModRhino/database"
	"avitoModRhino/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

//SEGMENTS

// Добавление сегмента
func NewSegment(c *fiber.Ctx) error {
	query := new(models.SegmentQuery)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//Проверка существует ли сегмент который пытаемся добавить
	emptyQuery := new(models.SegmentQuery)
	if txGorm := database.DB.Db.Where("segment = ?", query.Segment).Find(&emptyQuery); txGorm.RowsAffected != 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Segment you are trying to add already exists.",
		})
	}

	database.DB.Db.Create(&query)

	return c.Status(fiber.StatusOK).JSON(query)
}

// Вывод сегментов
func ShowSegment(c *fiber.Ctx) error {
	query := []models.SegmentQuery{}

	txGorm := database.DB.Db.Find(&query)

	if txGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There are no segments.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

// Удаление сегментов
func DeleteSegment(c *fiber.Ctx) error {
	query := new(models.SegmentQuery)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	txGorm := database.DB.Db.Where("segment = ?", query.Segment).Delete(&query)
	if txGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no segment like %s", query.Segment),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Segment deleted successfully.",
	})
}
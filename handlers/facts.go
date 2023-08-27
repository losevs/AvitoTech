package handlers

import (
	"avitoModRhino/database"
	"avitoModRhino/models"
	"fmt"
	"strconv"

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

//USERS

// Создание/Изменение пользователя
func UserSegment(c *fiber.Ctx) error {
	query := new(models.UserQuery)
	emptyUser := new(models.User)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//Если id существует
	if txGorm := database.DB.Db.Where("id = ?", query.ID).First(&emptyUser); txGorm.RowsAffected != 0 {

		//Если есть запрос на добавление сегментов
		message := ""
		if len(query.AddSeg) != 0 {
		loop:
			for _, seg := range query.AddSeg {
				segmentsPool := new(models.SegmentQuery)
				if gormDb := database.DB.Db.Where("segment = ?", seg).First(&segmentsPool); gormDb.RowsAffected != 0 {
					//Проверка есть ли у существующего User элемент который пытаются добавить
					for _, segmentsUser := range emptyUser.Segments {
						if segmentsUser == seg {
							continue loop
						}
					}
					emptyUser.Segments = append(emptyUser.Segments, seg)
				} else {
					message = " One or more segments are not in the list."
				}
			}
			database.DB.Db.Model(&emptyUser).UpdateColumns(models.User{Segments: emptyUser.Segments})
		}

		//Если есть запрос на удаление сегментов
		if len(query.DelSeg) != 0 {
			for _, seg := range query.DelSeg {
				index := -1
				for i, s := range emptyUser.Segments {
					if s == seg {
						index = i
						break
					}
				}
				if index != -1 {
					emptyUser.Segments = append(emptyUser.Segments[:index], emptyUser.Segments[index+1:]...)
				}
			}
			database.DB.Db.Model(&emptyUser).UpdateColumns(models.User{Segments: emptyUser.Segments})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Users segments are changed." + message,
			"User":    emptyUser,
		})
	}

	newUser := new(models.User)
	message := ""
addLoop:
	for _, seg := range query.AddSeg {
		segmentsPool := new(models.SegmentQuery)
		//Проверка есть ли сегмент в списке сегментов
		if gormDb := database.DB.Db.Where("segment = ?", seg).Find(&segmentsPool); gormDb.RowsAffected != 0 {
			//Проверка есть ли повторяющиеся сегменты в списке, который пытаемся добавить
			for _, segms := range newUser.Segments {
				if segms == seg {
					continue addLoop
				}
			}
			newUser.Segments = append(newUser.Segments, seg)
		} else {
			message = " One or more segments are not in the list."
		}
	}
	database.DB.Db.Create(models.User{ID: query.ID, Segments: newUser.Segments})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "New user created." + message,
		"User":    models.User{ID: query.ID, Segments: newUser.Segments},
	})
}

// Вывод пользователей
func UserShow(c *fiber.Ctx) error {
	query := []models.User{}
	txGorm := database.DB.Db.Order("id asc").Find(&query)

	if txGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User list is empty.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

// Удаление пользователя
func UserDelete(c *fiber.Ctx) error {
	needId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	txGorm := database.DB.Db.Where("id = ?", needId).Delete(&models.User{})
	if txGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no ID like %d", needId),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully.",
	})
}

// Вывод одного пользователя по ID
func ShowExactUser(c *fiber.Ctx) error {
	query := new(models.User)
	needId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	txGorm := database.DB.Db.Where("id = ?", needId).First(&query)
	if txGorm.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no ID like %d", needId),
		})
	}

	return c.Status(fiber.StatusOK).JSON(query)
}

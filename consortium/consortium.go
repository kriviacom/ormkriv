package consortium

import (
        "github.com/gofiber/fiber/v2"
)

func GetConsortiums(c *fiber.Ctx) {
        c.Send("All Books")
}

func GetConsortium(c *fiber.Ctx) {
        c.Send("Single Book")
}

func NewGetConsortium(c *fiber.Ctx) {
        c.Send("New Book")
}

func DeleteConsortium(c *fiber.Ctx) {
        c.Send("Delete Book")
}
package consortium

import (
        "github.com/gofiber/fiber"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
)

func GetConsortiums(c *fiber.Ctx, db *sql.DB) {
        c.SendString("All Books x")
}

func GetConsortium(c *fiber.Ctx) {
        c.SendString("Single Book")
}

func NewGetConsortium(c *fiber.Ctx) {
        c.SendString("New Book")
}

func DeleteConsortium(c *fiber.Ctx) {
        c.SendString("Delete Book")
}

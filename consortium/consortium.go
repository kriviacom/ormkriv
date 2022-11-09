package consortium

import (
        "github.com/gofiber/fiber"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
)

var db *sql.DB

func setDatabase(db2 *sql.DB)
{
        
        db := db2
           
}

func GetConsortiums(c *fiber.Ctx) {
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

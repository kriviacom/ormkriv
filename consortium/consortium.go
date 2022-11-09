package consortium

import (
        "github.com/gofiber/fiber"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
)

var db *sql.DB

func setDatabase(s string)
{
        var err error
        db, err = *sql.Open("mysql", s)
        var v = "Não conseguiu conectar ao banco de dados"
        
        if err != nil {
                fmt.Println(v)
                panic(err)
                }

        fmt.Println("conexão OK!")
        fmt.Println(db)
        
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

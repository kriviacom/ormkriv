package consortium

import (
        "github.com/gofiber/fiber"
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
)

func GetConsortiums(c *fiber.Ctx, db *sql.DB) {
     
	rows, err := db.Query("SELECT name FROM consortium")
	if err != nil {
		fmt.Println( "error llamado" )
		panic(err)
		}

	for rows.Next() {
		var name string

		rows.Scan(&name)
                c.SendString(""+name)
        
		fmt.Println( "Post title: "+name )
		}
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

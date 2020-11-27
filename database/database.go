package database

import (
	"database/sql"
	"fmt"
)

/*InitDB permite establecer la conexion a la BD*/
func InitDB() *sql.DB {
	connectionString := "upvmdnz7hynr0nlg:NuPdXP8q08f8qVcUrSYU@tcp(bcgpasrchxpx214jvw7d-mysql.services.clever-cloud.com:3306)/bcgpasrchxpx214jvw7d"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Conexion invalida a la BD")
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}

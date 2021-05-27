package main
import(
	"fmt"
	"TODO/Config"
	//Controllers "github.com/SidBthegr8/TO-DOApp/Controllers"
	"TODO/Models"
	"TODO/Routes"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)
var err error
func main(){
	// Creating a connection to the database
	//dsn := Config.DbURL(Config.BuildDBConfig())
	//fmt.Println(dsn)
  
	Config.DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("status: ", err)
	}

	//defer Config.DB.Close()

	// run the migrations: todo struct
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup routes 
	r := Routes.SetupRouter()

	// running
	r.Run()
}
/*import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
  )
  
  func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	fmt.Println(err)
	fmt.Println(db)
  }*/
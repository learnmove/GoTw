# GoTw
GoTw is a simple Mvc structure Demo

repository
model
controller

your can do :

go run tw.go -do make:model -name User
go run tw.go -do make:repository -name User
go run tw.go -do make:controller -name User

Migrate Database like :
go run tw.go -do migrate
go run tw.go -do seed      // it is not clear table data and seed into table
go run tw.go -do migrate:seed // it is refresh all data in table



All you need to do:
1.change project name  and import pakcage name like:
  	"github.com/goblog/app/repositories"  to 	"github.com/your_project_name/app/repositories"

2.change Mysql Password ,User ,Database Name
  in  github/your_project_name/database/migrations/migration.go
  in  github/your_project_name/database/seeders/seeder.go
  in	"github.com/your_project_name/app/repositories/DB.go"
   like this
   	db, err := gorm.Open("mysql", "root:password@/databasename?charset=utf8&parseTime=True&loc=Local")
  







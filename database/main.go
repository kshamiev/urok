package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql/driver"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"

	"urok/database/most"
	"urok/database/types"
)

// https://github.com/volatiletech/sqlboiler
//
// Установка в режиме модуля (без вендоринга)
// go get -u github.com/volatiletech/sqlboiler@v3.4.0
// go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@v3.4.0
// go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql@v3.4.0
//
// Проектируем БД
//
// Генерация моделей по БД
// sqlboiler psql -p post -o post
// sqlboiler mysql -p most -o most
//
// Проверка и тестирование моделей
// go test ./models
func main() {
	var err error
	var db *sql.DB
	// Open handle to database like normal
	db, err = sql.Open("mysql", "root:qwerty@/test?charset=utf8&parseTime=True&loc=Local")
	def(err)
	err = db.Ping()
	def(err)
	boil.DebugMode = true

	cust := &most.Customer{
		ID:types.NewUUID(),
		Name:null.String{
			String:"anufriy 222",
			Valid:true,
		},
	}
	ctx := context.Background()
	err = cust.Insert(ctx, db, boil.Infer())
	def(err)


	// uuid.FromStringOrNil("a7bc5424-70ad-4d6c-8961-fcb63140470e")

	custNew := &most.Customer{
		ID: types.UUID{uuid.FromStringOrNil("a7bc5424-70ad-4d6c-8961-fcb63140470e")},
	}

	err = custNew.Reload(ctx, db)
	def(err)

	fmt.Println(custNew)

	// boil.DebugWriter
	// u := &models.Role{
	// 	Name: null.NewString("Qwerty", false),
	// }
	// err = u.Insert(context.Background(), db, boil.Infer())
	// def(err)
	//
	// data, _ := models.Roles().All(context.Background(), db)
	// for i := range data {
	// 	fmt.Println(data[i])
	// }
	fmt.Println("ok")
}

func def(err error) {
	if err != nil {
		panic(err)
	}
}

func postgres() {

	// var uid uuid.UUID
	// var err error
	// var db *sql.DB
	// // Open handle to database like normal
	// db, err = sql.Open("postgres", "dbname=test host=localhost user=postgres password=postgres")
	// def(err)
	// err = db.Ping()
	// def(err)
	// boil.DebugMode = true

	// uid, err = uuid.NewV4()
	// u := &models.User{
	// 	ID:       uid.String(),
	// 	UserName: "popcorn 5",
	// }
	// ctx := context.Background()
	// // ctx = boil.SkipTimestamps(ctx)
	// err = u.Insert(ctx, db, boil.Infer())
	// def(err)

	// fmt.Println(u)

	// boil.DebugWriter
	// u := &models.Role{
	// 	Name: null.NewString("Qwerty", false),
	// }
	// err = u.Insert(context.Background(), db, boil.Infer())
	// def(err)
	//
	// data, _ := models.Roles().All(context.Background(), db)
	// for i := range data {
	// 	fmt.Println(data[i])
	// }
}

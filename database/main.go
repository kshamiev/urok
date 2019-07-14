package main

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
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
// sqlboiler psql
// sqlboiler psql -p qwerty -o qwerty --add-global-variants --no-context
//
// Проверка и тестирование моделей
// go test ./models
func main() {

	// var uid uuid.UUID
	var err error
	var db *sql.DB
	// Open handle to database like normal
	db, err = sql.Open("postgres", "dbname=test host=localhost user=postgres password=postgres")
	def(err)
	err = db.Ping()
	def(err)
	boil.DebugMode = true

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

func def(err error) {
	if err != nil {
		panic(err)
	}
}

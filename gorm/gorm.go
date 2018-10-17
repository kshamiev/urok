package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/kshamiev/sungora.v2/lg"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/world?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		lg.Dumper(db, err.Error())
	}
	defer db.Close()

	//db.CreateTable(&User{})

	// create
	//u := new(User)
	//u.Name = "Вася 1"
	//u.Age = 2365
	//u.Birthday = time.Now().Add(time.Hour * 5)
	//u.Description = "Описание 1"
	//lg.Dumper(u)
	// db.Create(u)
	//db.Table("users_copy").Create(u)
	//lg.Dumper(u)

	// QUERY

	// ORM
	//var users []*User
	//err = db.Find(&users, "name = ? AND ID >= ?", "Вася", 3).Error
	//lg.Dumper(users, err)

	//var users []*User
	//var count int32
	//err = db.Select("id, name").Table("users_copy").Where("ID > ?", 1056).Order("ID ASC").Limit(2).Find(&users).Error
	//err = db.Select("id, name").Table("users_copy").Where("ID > ?", 10).Order("ID ASC").Find(&users).Count(&count).Error
	//err = db.Select("id, name").Where("ID > ?", 10).Order("ID ASC").Find(&users).Count(&count).Error
	//lg.Dumper(users, count, err)

	// получение только количества записей
	//var count1 int32
	//err = db.Model(&User{}).Count(&count1).Error
	//lg.Dumper(count1, err)
	//err = db.Table("users_copy").Count(&count1).Error
	//lg.Dumper(count1, err)

	// если нужно получить данные из одной колонки в срез
	//var names []string
	//db.Model(&User{}).Pluck("name", &names)
	//lg.Dumper(names)

	// если нужно создать объект с учетом наличия его в БД по условиям
	// var user User
	// db.Where(User{Name: "Вася"}).FirstOrInit(&user) // не создает просто ищет  ( см. Attrs(User{Age: 30}) Assign(User{Age: 20}) )
	// db.Where(User{Name: "Муся"}).FirstOrCreate(&user) // создает если нет ( см. Attrs(User{Age: 30}) Assign(User{Age: 20}) )
	// lg.Dumper(user)

	// Raw SQL 1
	//type Result struct {
	//	Name string
	//	Age  int
	//}
	//var result []Result
	//db.Select("name, age").Table("users").Where("name = ?", "Вася").Scan(&result)
	//lg.Dumper(result, err)

	// Raw SQL 2
	//var users []*User
	//db.Raw("SELECT c.name, c.age FROM users as c WHERE c.name = ?", "Вася").Scan(&users)
	//lg.Dumper(users)

	// транзакции
	// db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)

}

type User struct {
	//	gorm.Model
	ID          uint64    `gorm:"primary_key"`
	Name        string    `gorm:"default:NULL"`                      // если значение может быть NULL
	Age         int64     ``                                         // если ничего - для форматирвоания
	Birthday    time.Time `gorm:"default:NULL"`                      //
	Cnt         int32     `gorm:"default:20"`                        // если значение не может быть NULL
	Description string    `gorm:"column:opisanie;default:'popcorn'"` // если поле в бд имеет другое имя и не null
	Data        []byte    `gorm:"-"`                                 // исключает использование с БД
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// переопредение таблицы источника обьектов
func (self *User) TableName() string {
	return "users_copy"
}

// функция - хук вызовется перед вставкой - созданием записи
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	// scope.SetColumn("ID", 1050)
	return nil
}

// функция - хук вызовется после вставки - создания записи
func (user *User) AfterCreate(scope *gorm.Scope) error {

	return nil
}

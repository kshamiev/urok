package stmc

type Index interface {
	GetIndexName() string // Индекс (таблица) в мантикоре
}

type Parser interface {
	Parse(map[string]interface{}) // Метод принимающий результаты запроса по строчно для наполнения данным
}

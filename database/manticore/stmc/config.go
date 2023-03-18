package stmc

import (
	"time"

	"github.com/shopspring/decimal"

	"gitlab.tn.ru/golang/app/scheduler"
)

type Config struct {
	DriverName      string          `yaml:"driverName"`      // mysql
	Dsn             string          `yaml:"dsn"`             // tcp(127.0.0.1:9306)/
	MaxIdleConns    int             `yaml:"maxIdleConns"`    // Свободные соединения (необязательно)
	ConnMaxIdleTime time.Duration   `json:"connMaxIdleTime"` // Время жизни с. с. (необязательно)
	MaxOpenConns    int             `yaml:"maxOpenConns"`    // Открытые соединения (необязательно)
	ConnMaxLifetime time.Duration   `json:"connMaxLifetime"` // Время жизни о. с. (необязательно)
	SearchOptions   string          `yaml:"searchOptions"`   // Опции для поисковых запросов по умолчанию (опционально)
	DecimalDot      int             `yaml:"decimalDot"`      // Точность дробных чисел (к. знаков после запятой)
	DotDecimal      decimal.Decimal `yaml:"-"`               // Точность дробных чисел (к. знаков после запятой)
	// Следующие конфигурации только для сервиса индексатора
	DataDir      string              `yaml:"dataDir"`      // Хранилище проиндексированных данных (маппинг, репликации)
	SearchConfig string              `yaml:"searchConfig"` // Конфигурация для мантикоры
	TaskIndexes  []scheduler.Crontab `yaml:"taskIndexes"`  // Настройки индексов
}

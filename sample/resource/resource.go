package resource

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"text/template"
	"time"
)

// Генератор кода встраиваемых ресурсов. Аргументы:
// Ожидаются переменные окружения:
// EMBED_STATIC_RESOURCES - Группы ресурсов и папками самих ресурсов.
//     Формат: group_name1:path/to/folder1,group_name2:/absolut/path/to/folder2
// Относительные пути берутся от текущей рабочей директории
//
// Создаются .go файлы по следующему шаблону:
// resource_{{ groupName }}_{{ resourceNumber }}.go
// Где:
// groupName      - строка приведённая к нижнему регистру не содержащая пробелы.
// resourceNumber - порядковый номер ресурса в формате %020d.
// Пример: resource_test_group_00000000000000000001.go

func init() { singleton = &Storage{res: make(map[string]map[string]*Resource)} }

func Get() *Storage { return singleton }

var singleton *Storage

type Resource struct {
	Size        int64     // Размер ресурса в байтах.
	Time        time.Time // Дата и время создания ресурса.
	ContentType string    // Определённый по расширению имени файла тип контента ресурса.
	Content     []byte    // Контент ресурса.
}

// Execute компиляция html шаблона и сборка контента
func (rec Resource) Execute(fun, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	const nameTpl = "default"
	var tpl *template.Template
	if tpl, err = template.New(nameTpl).Funcs(fun).Parse(string(rec.Content)); err != nil {
		return ret, err
	}

	err = tpl.Execute(&ret, variables)
	return ret, err
}

type Storage struct {
	res     map[string]map[string]*Resource // Карта ресурсов. map[название группы]map[название ресурса]*resource
	resLock sync.RWMutex                    // Защита карты от конкурентного доступа на запись.
}

// Add Добавление ресурса в группу ресурсов.
func (rec *Storage) Add(group, name string, resource Resource) (err error) {
	const (
		errDuplicate = "ресурс %q в группе ресурсов %q уже существует"
		errSize      = "размер контента в описании не соответствует фактическому размеру контента"
	)
	var (
		ok bool
		n  int
	)

	rec.resLock.Lock()
	defer rec.resLock.Unlock()
	if _, ok = rec.res[group]; !ok {
		rec.res[group] = make(map[string]*Resource)
	}
	if _, ok = rec.res[group][name]; ok {
		err = fmt.Errorf(errDuplicate, name, group)
		return
	}
	rec.res[group][name] = &Resource{
		Size:        resource.Size,
		Time:        resource.Time,
		ContentType: resource.ContentType,
		Content:     make([]byte, len(resource.Content)),
	}
	n = copy(rec.res[group][name].Content, resource.Content)
	if int64(n) != rec.res[group][name].Size {
		err = errors.New(errSize)
		return
	}

	return
}

// Group Возвращается список групп ресурсов.
func (rec *Storage) Group() (ret []string) {
	var group string

	rec.resLock.RLock()
	defer rec.resLock.RUnlock()
	ret = make([]string, 0, len(rec.res))
	for group = range rec.res {
		ret = append(ret, group)
	}

	return
}

// ResourceByGroup Возвращает список ресурсов в указанной группе.
func (rec *Storage) ResourceByGroup(group string) (ret []string) {
	var (
		rn string
		ok bool
	)

	rec.resLock.RLock()
	defer rec.resLock.RUnlock()
	if _, ok = rec.res[group]; !ok {
		return
	}
	ret = make([]string, 0, len(rec.res[group]))
	for rn = range rec.res[group] {
		ret = append(ret, rn)
	}

	return
}

// ResourceData Получение ресурса по имени группы и ресурсу. Если ресурса нет, возвращается nil.
func (rec *Storage) ResourceData(group, resource string) (ret *Resource) {
	var ok bool

	rec.resLock.RLock()
	defer rec.resLock.RUnlock()
	if _, ok = rec.res[group]; !ok {
		return
	}
	if _, ok = rec.res[group][resource]; !ok {
		return
	}
	ret = &Resource{
		Size:        rec.res[group][resource].Size,
		Time:        rec.res[group][resource].Time,
		ContentType: rec.res[group][resource].ContentType,
		Content:     make([]byte, len(rec.res[group][resource].Content)),
	}
	copy(ret.Content, rec.res[group][resource].Content)

	return
}

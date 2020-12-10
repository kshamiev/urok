// Инструмент по автоматизации сопоставления типов для GRPC
//
// Генерация описаний прототипов, самих прототипов и методов конвертации типа в обе стороны.
// Обрабатываются только публичные и помеченные тегом json поля структур.
//
// Из коробки обрабатывает базовые типы golang (string, bool, int..., uint..., float..., []byte, []string)
// + typ.UUID - реализация работы с полями UUID
// + time.Time - дата и время
// + time.Duration - время
// + decimal.Decimal - работа с дробными числами
// + ENUM в парадигме GRPC
// + ссылки на другие типы
// + срезы ссылок на другие типы
// + имеет спецификацию работы с типами используемыми в библиотеке boiler
// (null.JSON, null.Bytes, null.String, null.Time, types.StringArray)
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"gitlab.services.mts.ru/3click/million/typs"
)

const (
	suffixFromProto = "FromProto" // config suffix name func
	suffixToProto   = "ToProto"   // config suffix name func
	pkgProtoAlias   = "pb"        // not recommended change
	pkgType         = "typs"      // not recommended change
	suffixSlice     = "Slice"     // config prefix slice types
)

func init() {
	if len(typs.Config) == 0 {
		log.Fatal("config empty")
	}

	var rt = reflect.TypeOf(typs.Config[0])
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	list := strings.Split(rt.PkgPath(), "/")
	pkgProtoImport = strings.Join(list[:len(list)-1], "/")
}

var pkgProtoImport string

func main() {
	if len(typs.Config) == 0 {
		log.Fatal("config empty")
	}
	var err error
	var tplPFull, tplMFull, tplP, tplM string
	gen := Generate{
		controlType: map[string]bool{},
	}

	tplPFull = CreateProtoFile()
	tplMFull = CreateTypeFile()
	for _, t := range typs.Config {
		if tplP, tplM, err = gen.ParseType(t); err != nil {
			log.Fatal(err)
		}
		tplPFull += tplP
		tplMFull += tplM
	}
	// type proto (описание прототипов)
	if err = ioutil.WriteFile(pkgProtoAlias+"/"+pkgType+".proto", []byte(tplPFull), 0600); err != nil {
		log.Fatal(err)
	}
	// golang методы конвертации
	if err = ioutil.WriteFile(pkgType+"/advanced.go", []byte(tplMFull), 0600); err != nil {
		log.Fatal(err)
	}
	fmt.Println(pkgType + ".proto OK")
}

// //// TYPE

type Generate struct {
	fieldCheckParseCpunt map[string]bool
	controlType          map[string]bool
}

// Анализируем тип и формируем его сопряжение с grpc (proto файлы и методы конвертации) (Object = *TypeName)
func (gen *Generate) ParseType(object interface{}) (tplP, tplM string, err error) {
	// разбираем тип
	var value = reflect.ValueOf(object)
	if value.Kind() != reflect.Ptr {
		return tplP, tplM, errors.New("error: " + value.Type().String() + " not ptr")
	}
	if value.IsNil() {
		return tplP, tplM, errors.New("error: " + value.Type().String() + "is null")
	}
	value = value.Elem()

	list := strings.Split(value.Type().String(), ".")
	if _, ok := gen.controlType[list[1]]; ok {
		return tplP, tplM, errors.New("error: type '" + list[1] + "' duplicate (is already defined type)")
	} else if pkgType != list[0] {
		return tplP, tplM, errors.New("error: pkg '" + list[0] + "' invalid (all defined types must one pkg)")
	}
	gen.controlType[list[1]] = true

	// pb
	tplP = "\nmessage " + list[1] + " {\n"

	// one object proto to type
	tplMFrom := "\nfunc New" + list[1] + suffixFromProto + "(pit *" + pkgProtoAlias + "." + list[1] + ") *" + list[1] + " {\n"
	tplMFrom += "\tif pit == nil { return nil }\n"
	tplMFrom += "\treturn &" + list[1] + "{\n"

	// one object type to proto
	tplMTo := "\nfunc New" + list[1] + suffixToProto + "(tip *" + list[1] + ") *" + pkgProtoAlias + "." + list[1] + " {\n"
	tplMTo += "\tif tip == nil { return nil }\n"
	tplMTo += "\treturn &" + pkgProtoAlias + "." + list[1] + "{\n"

	gen.fieldCheckParseCpunt = make(map[string]bool)
	tplP_, tplMFrom_, tplMTo_, n_ := gen.ParseTypeField(value, 0, 0)
	if n_ == -100 {
		return "", "", nil
	}
	tplP += tplP_
	tplMFrom += tplMFrom_
	tplMTo += tplMTo_

	tplP += "}\n"

	// slice type to proto
	tplMTo += "\t}\n}\n\n" + gen.GenerateFuncSliceTypeProto(list[1])

	// slice proto from type
	tplMFrom += "\t}\n}\n\n" + gen.GenerateFuncSliceProtoType(list[1])

	return tplP, tplMTo + tplMFrom, nil
}

// GenerateFuncSliceTypeProto генерация метода конвертации среза типа в срез его прототипа
func (gen *Generate) GenerateFuncSliceTypeProto(typ string) (s string) {
	s += fmt.Sprintf("func new%s%s"+suffixToProto+" (tip []*%s) []*%s.%s {", typ, suffixSlice, typ, pkgProtoAlias, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s.%s, len(tip))", pkgProtoAlias, typ)
	s += "\n\tfor i := range tip {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+suffixToProto+"(tip[i])\n\t}\n\treturn res\n}\n", typ)
	return
}

// GenerateFuncSliceProtoType генерация метода конвертации среза прототипа в соответсвующий ему срез типа
func (gen *Generate) GenerateFuncSliceProtoType(typ string) (s string) {
	s += fmt.Sprintf("func new%s%s"+suffixFromProto+"(list []*%s.%s) []*%s {", typ, suffixSlice, pkgProtoAlias, typ, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s, len(list))", typ)
	s += "\n\tfor i := range list {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+suffixFromProto+"(list[i])\n\t}\n\treturn res\n}\n", typ)
	return
}

// разбираем свойства типа
func (gen *Generate) ParseTypeField(value reflect.Value, num, level int) (tplP, tplMFrom, tplMTo string, i int) {
	tplMFromEnd := ""
	if level > 0 {
		if strings.Split(value.Type().String(), ".")[0] == pkgType {
			tplMFrom += value.Type().Name() + ": " + strings.Split(value.Type().String(), ".")[1] + "{\n"
		} else {
			tplMFrom += value.Type().Name() + ": " + value.Type().String() + "{\n"
		}
		tplMFromEnd = "},\n"
	}

	for i = 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldStruct := value.Type().Field(i)

		// пропускаем приватные свойства
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		// пропускаем исключенные и не обозначенные свойства
		fieldJSON := fieldStruct.Tag.Get(`json`)
		if fieldJSON == `-` || fieldJSON == "" {
			continue
		}

		// встроенные поля
		if fieldStruct.Anonymous {
			tplP_, tplMFrom_, tplMTo_, n_ := gen.ParseTypeField(field, i+num, level+1)
			if n_ == -100 {
				return "", "", "", -100
			}
			tplP += tplP_
			tplMFrom += tplMFrom_
			tplMTo += tplMTo_
			num += n_
			continue
		}

		if _, ok := gen.fieldCheckParseCpunt[fieldStruct.Name]; ok {
			fmt.Printf("ambiguous anonymous property: %s.%s [%s]\n",
				path.Base(value.Type().PkgPath()), value.Type().Name(), fieldStruct.Name)
			return "", "", "", -100
		}
		gen.fieldCheckParseCpunt[fieldStruct.Name] = true

		tplP_, tplMFrom_, tplMTo_ := gen.ParseField(field, fieldStruct, i+num)
		tplP += tplP_
		tplMFrom += tplMFrom_
		tplMTo += tplMTo_
	}
	tplMFrom += tplMFromEnd
	i--

	return
}

// //// FIELD

const msg_int = "[(grpc.gateway.protoc_gen_swagger.options.openapiv2_field) = {type: INTEGER}]"

// Анализируем свойство типа и генерируем описание его прототипа и методы конвертации в обе стороны
// nolint
func (gen *Generate) ParseField(field reflect.Value, fieldStr reflect.StructField, i int) (tplP, tplMFrom, tplMTo string) {
	fieldName := fieldStr.Name
	fieldJSON := strings.Split(fieldStr.Tag.Get(`json`), ",")[0]

	// формируем согласно типу
	fieldKind := field.Type().Kind()
	fieldType := field.Type().String()
	subjErr := "not implemented undefined property: %s.%s [%s] %s"
	subjErr = fmt.Sprintf(subjErr, field.Type().String(), fieldName, fieldKind.String(), fieldType)

	if f, ok := typs.ConfigHandlerFunc[fieldType]; ok {
		return f(i, fieldName, fieldJSON, convFP(fieldJSON))
	}

	switch fieldKind {
	case reflect.String:
		met := field.MethodByName("GetEnum")
		if met.IsValid() {
			tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = gen.GenerateFieldEnum(fieldName, fieldJSON, fieldType)
		} else {
			tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
		}
	case reflect.Bool:
		tplP += "\tbool " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Float32:
		tplP += "\tfloat " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Float64:
		tplP += "\tdouble " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Int:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldInt(i, fieldName, fieldJSON)
	case reflect.Int8:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldInt8(i, fieldName, fieldJSON)
	case reflect.Int16:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldInt16(i, fieldName, fieldJSON)
	case reflect.Int32:
		tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Int64:
		if fieldType == "time.Duration" {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
			tplMFrom, tplMTo = gen.GenerateTimeDuration(fieldName, fieldJSON)
		} else {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
			tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
		}
	case reflect.Uint:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldUint(i, fieldName, fieldJSON)
	case reflect.Uint8:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldUint8(i, fieldName, fieldJSON)
	case reflect.Uint16:
		tplP, tplMFrom, tplMTo = gen.GenerateFieldUint16(i, fieldName, fieldJSON)
	case reflect.Uint32:
		tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Uint64:
		tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
		tplMFrom, tplMTo = gen.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Slice:
		typParse := strings.Split(fieldType, pkgType+".")
		if fieldType == "[]string" {
			return gen.GenerateFieldStringArray(i, fieldName, fieldJSON)
		}
		if fieldType == "[]uint8" {
			return gen.GenerateFieldBytes(i, fieldName, fieldJSON)
		}
		if len(typParse) == 2 {
			typParseAdv := strings.Split(typParse[1], suffixSlice)
			if len(typParseAdv) == 2 {
				return gen.GenerateFieldSlicePtrType(i, typParseAdv[0], fieldName, fieldJSON)
			}
			if typParse[0] == "[]*" {
				return gen.GenerateFieldSlicePtrType(i, typParse[1], fieldName, fieldJSON)
			}
		}
		fmt.Println(subjErr)

	case reflect.Struct:
		fmt.Println(subjErr)

	case reflect.Ptr:
		typParse := strings.Split(fieldType, "*"+pkgType+".")
		if len(typParse) == 2 {
			return gen.GenerateFieldPtrType(i, typParse[1], fieldName, fieldJSON)
		}
		fmt.Println(subjErr)
	default:
		fmt.Println(subjErr)
	}

	return tplP, tplMFrom, tplMTo
}

// GenerateFieldPtrType конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldPtrType(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\t" + typParse + " " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: New%s"+suffixToProto+"(tip.%s),\n", convFP(fieldJSON), typParse, field)
	tplMFrom = fmt.Sprintf("%s: New%s"+suffixFromProto+"(pit.%s),\n", field, typParse, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldSlicePtrType конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldSlicePtrType(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated " + typParse + " " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: new%s%s"+suffixToProto+"(tip.%s),\n", convFP(fieldJSON), typParse, suffixSlice, field)
	tplMFrom = fmt.Sprintf("%s: new%s%s"+suffixFromProto+"(pit.%s),\n", field, typParse, suffixSlice, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateTimeDuration конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateTimeDuration(field, fieldJSON string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("%s: tip.%s.Nanoseconds(),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: time.Duration(pit.%s),\n", field, convFP(fieldJSON))
	return tplMFrom, tplMTo
}

// GenerateFieldUint8 конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldUint8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: uint32(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint8(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint16 конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldUint16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: uint32(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint16(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldUint(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: uint64(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt8 конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldInt8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: int32(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int8(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt16 конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldInt16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: int32(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int16(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldInt(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + " " + msg_int + ";\n"
	tplMTo = fmt.Sprintf("%s: int64(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int(pit.%s),\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldBytes конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldBytes(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s,\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: pit.%s,\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldStringArray конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldStringArray(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s,\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: pit.%s,\n", field, convFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNative конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldEnum(field, fieldJSON, fieldType string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("%s: string(tip.%s),\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: %s(pit.%s),\n", field, fieldType, convFP(fieldJSON))
	return tplMFrom, tplMTo
}

// GenerateFieldNative конвертация - сопоставление туда и обратно
func (gen *Generate) GenerateFieldNative(field, fieldJSON string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("%s: tip.%s,\n", convFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: pit.%s,\n", field, convFP(fieldJSON))
	return tplMFrom, tplMTo
}

// //// support method

// CreateTypeFile инициализация файла с методами конвертации типа
func CreateTypeFile() (proto string) {
	return "// nolint\npackage " + pkgType + "\n\nimport \"" + pkgProtoImport + "/pb\"\n"
}

// инициализация файла с описанием прототипов
func CreateProtoFile() (proto string) {
	if _, err := os.Stat(pkgProtoAlias); err != nil {
		if err := os.Mkdir(pkgProtoAlias, 0700); err != nil {
			log.Fatal(err)
		}
	}
	//
	const separator = "// AFTER CODE GENERATED. DO NOT EDIT //"
	if data, err := ioutil.ReadFile(pkgProtoAlias + "/" + pkgType + ".proto"); err == nil {
		proto = string(data)
		list := strings.Split(proto, separator)
		proto = list[0] + separator + "\n"
	} else {
		if _, err = ioutil.ReadFile(pkgProtoAlias + "/service.proto"); err != nil {
			if data, err = ioutil.ReadFile(pkgType + "/generate/service.proto"); err != nil {
				log.Fatal(err)
			}
			proto = string(data)
			if err = ioutil.WriteFile(pkgProtoAlias+"/service.proto", []byte(proto), 0600); err != nil {
				log.Fatal(err)
			}
		}
		//
		if data, err = ioutil.ReadFile(pkgType + "/generate/" + pkgType + ".proto"); err != nil {
			log.Fatal(err)
		}
		proto = string(data)
	}
	return proto
}

// convFP Получение названия свойства в прототипе (через тег json) для сосоставления
func convFP(fieldTag string) string {
	if fieldTag == "id" {
		return "Id"
	}
	list := strings.Split(fieldTag, "_")
	for i := range list {
		if _, err := strconv.Atoi(list[i]); err == nil {
			list[i] = "_" + list[i]
		} else {
			list[i] = strings.Title(list[i])
		}
	}
	return strings.Join(list, "")
}

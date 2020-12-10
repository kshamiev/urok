package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"

	"gitlab.services.mts.ru/3click/million/pkg/component"

	"github.com/iancoleman/strcase"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
)

var (
	defaultTypeMapping = map[string]TypeName{
		"bool":        {"sql.NullBool", "bool"},
		"int4":        {"sql.NullInt64", "int64"},
		"int8":        {"sql.NullInt64", "int64"},
		"jsonb":       {"[]byte", "[]byte"},
		"text":        {"sql.NullString", "string"},
		"_text":       {"sql.NullString", "string"},
		"varchar":     {"sql.NullString", "string"},
		"_varchar":    {"sql.NullString", "string"},
		"timestamptz": {"sql.NullTime", "time.Time"},
		"timestamp":   {"sql.NullTime", "time.Time"},
		"uuid":        {"sql.NullString", "string"},
		"numeric":     {"sql.NullFloat64", "float64"},
	}

	defaultImportMapping = map[string]string{
		"sql.NullBool":    "database/sql",
		"sql.NullInt64":   "database/sql",
		"sql.NullString":  "database/sql",
		"sql.NullTime":    "database/sql",
		"sql.NullFloat64": "database/sql",
		"time.Time":       "time",
	}
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := readCfg("database.yaml")
	db, err := component.NewPGStorage(&cfg.Postgres)
	panicErr(err)

	columns := columnsFromDB(db.GetDB(), &cfg.Generate)
	scheme := columnsToScheme(&cfg.Generate, columns)
	code := schemeToCode(scheme)
	cruds := schemeToCRUD(scheme)
	writeToFile("typs", code, cruds)
}

type Config struct {
	Postgres component.PostgresConfig `yaml:"postgres"`
	Generate GenerateConfig           `yaml:"generate"`
}

func readCfg(cfgPath string) *Config {
	var cfg Config

	data, err := ioutil.ReadFile(cfgPath)
	panicErr(err)

	err = yaml.Unmarshal(data, &cfg)
	panicErr(err)

	if cfg.Generate.TypeMapping == nil {
		cfg.Generate.TypeMapping = map[string]TypeName{}
	}

	for k, v := range defaultTypeMapping {
		if _, ok := cfg.Generate.TypeMapping[k]; !ok {
			cfg.Generate.TypeMapping[k] = v
		}
	}

	if cfg.Generate.ImportMapping == nil {
		cfg.Generate.ImportMapping = map[string]string{}
	}

	for k, v := range defaultImportMapping {
		if _, ok := cfg.Generate.ImportMapping[k]; !ok {
			cfg.Generate.ImportMapping[k] = v
		}
	}

	return &cfg
}

type ColumnProperties struct {
	SchemaName        string `db:"schema_name"`
	TableName         string `db:"table_name"`
	TableDescription  string `db:"table_description"`
	ColumnName        string `db:"column_name"`
	ColumnType        string `db:"column_type"`
	Nullable          bool   `db:"nullable"`
	ColumnDescription string `db:"column_description"`
}

func columnsFromDB(db *sqlx.DB, cfg *GenerateConfig) []*ColumnProperties {
	query := `select t.table_schema as schema_name, t.table_name, 
                     coalesce(dt.description, '') as table_description,
                     c.column_name, c.udt_name as column_type,
                     c.is_nullable = 'YES' nullable, 
                     coalesce(dc.description, '') as column_description
              from information_schema.tables t
                  left join information_schema.columns c
                       on c.table_name = t.table_name
                       and c.table_schema = t.table_schema
                  left join pg_catalog.pg_statio_all_tables st
                       on st.relname = t.table_name
                       and st.schemaname = t.table_schema
                  left join pg_catalog.pg_description dt
                       on dt.objoid = st.relid
                       and dt.objsubid = 0
                  left join pg_catalog.pg_description dc
                       on dc.objoid = st.relid
                       and dc.objsubid = c.ordinal_position
              where t.table_schema in (%s)
              order by case when t.table_schema = $1 then 0
                  else 1 end, t.table_schema, t.table_name;`

	if len(cfg.DBSchemes) == 0 {
		panic("at least one schema needed")
	}

	args := []interface{}{cfg.DBSchemes[0]}
	in := "$1"

	for i, schema := range cfg.DBSchemes[1:] {
		args = append(args, schema)
		in += ", $" + strconv.Itoa(i+2)
	}

	query = fmt.Sprintf(query, in)
	var result []*ColumnProperties

	err := db.SelectContext(context.Background(), &result, query, args...)
	panicErr(err)

	return result
}

type GenerateConfig struct {
	DBSchemes     []string            `yaml:"db_schemes"`
	TypeMapping   map[string]TypeName `yaml:"type_mapping"`
	ImportMapping map[string]string   `yaml:"import_mapping"`
	CRUD          []string            `yaml:"crud"`
}

type TypeName struct {
	NullName    string `yaml:"null_name"`
	NotNullName string `yaml:"not_null_name"`
}

func columnsToScheme(cfg *GenerateConfig, columns []*ColumnProperties) *Scheme {
	result := &Scheme{imports: map[string]struct{}{}}
	if len(columns) == 0 {
		return result
	}

	crudSet := map[string]struct{}{}
	for _, name := range cfg.CRUD {
		crudSet[name] = struct{}{}
	}

	tablesSet := map[string]struct{}{}
	var currentStruct *Struct
	correctCRUD := true

	for i, qr := range columns {
		if i == 0 || qr.SchemaName != columns[i-1].SchemaName || qr.TableName != columns[i-1].TableName {
			if !correctCRUD {
				panic("couldn't find column 'id' at table '" + qr.TableName + "'")
			}

			structName := strcase.ToCamel(qr.TableName)
			if _, ok := tablesSet[structName]; ok {
				structName += "_" + strcase.ToCamel(qr.SchemaName)
			}

			_, ok := crudSet[qr.TableName]
			tablesSet[structName] = struct{}{}
			currentStruct = &Struct{
				name:        structName,
				tag:         qr.SchemaName + "." + qr.TableName,
				description: qr.TableDescription,
				crud:        ok,
			}
			correctCRUD = !ok

			result.structs = append(result.structs, currentStruct)
		}

		typ := ""
		if qr.Nullable {
			typ = cfg.TypeMapping[qr.ColumnType].NullName
		} else {
			typ = cfg.TypeMapping[qr.ColumnType].NotNullName
		}

		if typ == "" {
			nullable := "null"
			if qr.Nullable {
				nullable = "not_null"
			}

			panic("couldn't find '" + qr.ColumnType + "' " + nullable + " mapping")
		}

		if imp, ok := cfg.ImportMapping[typ]; ok {
			result.imports[imp] = struct{}{}
		}

		if !correctCRUD && qr.ColumnName == "id" {
			correctCRUD = true
		}

		currentStruct.fields = append(currentStruct.fields, &Field{
			name: strcase.ToCamel(qr.ColumnName),
			// description: qr.ColumnDescription,
			typ: typ,
			tag: qr.ColumnName,
		})
	}

	return result
}

type Scheme struct {
	imports map[string]struct{}
	structs []*Struct
}

type Struct struct {
	name        string
	tag         string
	description string
	fields      []*Field
	crud        bool
}

type Field struct {
	name        string
	description string
	typ         string
	tag         string
}

func schemeToCode(scheme *Scheme) string {
	res := "package typs\n\n"

	switch {
	case len(scheme.imports) == 1:
		for imp := range scheme.imports {
			res += "import " + imp + "\n\n"
		}
	case len(scheme.imports) > 1:
		res += "import (\n"
		for imp := range scheme.imports {
			res += "\"" + imp + "\"" + "\n"
		}

		res += ")\n\n"
	}

	for _, str := range scheme.structs {
		if str.description != "" {
			res += "// " + str.description
		}

		res += "\ntype " + str.name + " struct {\n"

		for _, field := range str.fields {
			res += field.name + " " + field.typ + " `json:\"" + field.tag + "\" db:\"" + field.tag + "\"`"
			if field.description != "" {
				res += " // " + field.description
			}

			res += "\n"
		}

		res += "}\n\n"
	}

	return res
}

func schemeToCRUD(scheme *Scheme) string {
	res := "package typs\n\n// nolint:lll // generated\n// language=sql\nconst(\n"
	for _, str := range scheme.structs {
		if !str.crud {
			continue
		}

		res += "SQL" + str.name + "Select = \"SELECT"
		for i, f := range str.fields {
			if i > 0 {
				res += ","
			}

			res += " " + f.tag
		}

		const commaDelimiter = ", "
		args := ""
		res += " FROM " + str.tag + " WHERE id = $1;\"\nSQL" + str.name + "Insert = \"INSERT INTO " + str.tag + " ("

		for i, f := range str.fields {
			if i > 0 {
				res += commaDelimiter
				args += commaDelimiter
			}

			args += "$" + strconv.Itoa(i+1)
			res += f.tag
		}

		res += ") VALUES (" + args + ");\"\n"
	}

	res += ")\n"

	return res
}

func writeToFile(targetPath, code, crud string) {
	err := ioutil.WriteFile(path.Join(targetPath, "/typs.go"), []byte(code), 0600)
	panicErr(err)

	err = ioutil.WriteFile(path.Join(targetPath, "/sql.go"), []byte(crud), 0600)
	panicErr(err)
}

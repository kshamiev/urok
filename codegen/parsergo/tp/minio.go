// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package tp

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// FileService is an object representing the database table.
type FileService struct {
	// ИД
	ID uuid.UUID `db:"id" pg:"id,pk" boil:"id" json:"id" example:"8ca3c9c3-cf1a-47fe-8723-3f957538ce42"`
	// папка хранения - тип объекта
	Bucket string `db:"bucket" pg:"bucket,use_zero" boil:"bucket" json:"bucket"`
	// файл хранения - ид объекта
	ObjectID int64 `db:"object_id" pg:"object_id,use_zero" boil:"object_id" json:"object_id"`
	// имя файла
	Name string `db:"name" pg:"name,use_zero" boil:"name" json:"name"`
	// тип файла
	FileType string `db:"file_type" pg:"file_type,use_zero" boil:"file_type" json:"file_type"`
	// размер файла
	FileSize int `db:"file_size" pg:"file_size,use_zero" boil:"file_size" json:"file_size"`
	// дополнительные параметры файла
	Label null.String `db:"label" pg:"label,use_zero" boil:"label" json:"label,omitempty" swaggertype:"string"`
	// пользователь
	UserLogin string `db:"user_login" pg:"user_login,use_zero" boil:"user_login" json:"user_login"`
	// дата и время создания
	CreatedAt time.Time `db:"created_at" pg:"created_at,use_zero" boil:"created_at" json:"created_at" example:"2006-01-02T15:04:05Z"`
	// подтверждение загрузки
	IsConfirm bool `db:"is_confirm" pg:"is_confirm,use_zero" boil:"is_confirm" json:"is_confirm"`

	R *fileServiceR `db:"-" pg:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L fileServiceL  `db:"-" pg:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FileServiceColumns = struct {
	ID        string
	Bucket    string
	ObjectID  string
	Name      string
	FileType  string
	FileSize  string
	Label     string
	UserLogin string
	CreatedAt string
	IsConfirm string
}{
	ID:        "id",
	Bucket:    "bucket",
	ObjectID:  "object_id",
	Name:      "name",
	FileType:  "file_type",
	FileSize:  "file_size",
	Label:     "label",
	UserLogin: "user_login",
	CreatedAt: "created_at",
	IsConfirm: "is_confirm",
}

var FileServiceTableColumns = struct {
	ID        string
	Bucket    string
	ObjectID  string
	Name      string
	FileType  string
	FileSize  string
	Label     string
	UserLogin string
	CreatedAt string
	IsConfirm string
}{
	ID:        "minio.id",
	Bucket:    "minio.bucket",
	ObjectID:  "minio.object_id",
	Name:      "minio.name",
	FileType:  "minio.file_type",
	FileSize:  "minio.file_size",
	Label:     "minio.label",
	UserLogin: "minio.user_login",
	CreatedAt: "minio.created_at",
	IsConfirm: "minio.is_confirm",
}

// Generated where

type whereHelperuuid_UUID struct{ field string }

func (w whereHelperuuid_UUID) EQ(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperuuid_UUID) NEQ(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperuuid_UUID) LT(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperuuid_UUID) LTE(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperuuid_UUID) GT(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperuuid_UUID) GTE(x uuid.UUID) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var FileServiceWhere = struct {
	ID        whereHelperuuid_UUID
	Bucket    whereHelperstring
	ObjectID  whereHelperint64
	Name      whereHelperstring
	FileType  whereHelperstring
	FileSize  whereHelperint
	Label     whereHelpernull_String
	UserLogin whereHelperstring
	CreatedAt whereHelpertime_Time
	IsConfirm whereHelperbool
}{
	ID:        whereHelperuuid_UUID{field: "\"minio\".\"id\""},
	Bucket:    whereHelperstring{field: "\"minio\".\"bucket\""},
	ObjectID:  whereHelperint64{field: "\"minio\".\"object_id\""},
	Name:      whereHelperstring{field: "\"minio\".\"name\""},
	FileType:  whereHelperstring{field: "\"minio\".\"file_type\""},
	FileSize:  whereHelperint{field: "\"minio\".\"file_size\""},
	Label:     whereHelpernull_String{field: "\"minio\".\"label\""},
	UserLogin: whereHelperstring{field: "\"minio\".\"user_login\""},
	CreatedAt: whereHelpertime_Time{field: "\"minio\".\"created_at\""},
	IsConfirm: whereHelperbool{field: "\"minio\".\"is_confirm\""},
}

// FileServiceRels is where relationship names are stored.
var FileServiceRels = struct {
}{}

// fileServiceR is where relationships are stored.
type fileServiceR struct {
}

// NewStruct creates a new relationship struct
func (*fileServiceR) NewStruct() *fileServiceR {
	return &fileServiceR{}
}

// fileServiceL is where Load methods for each relationship are stored.
type fileServiceL struct{}

var (
	fileServiceAllColumns            = []string{"id", "bucket", "object_id", "name", "file_type", "file_size", "label", "user_login", "created_at", "is_confirm"}
	fileServiceColumnsWithoutDefault = []string{"bucket", "name", "file_type", "user_login"}
	fileServiceColumnsWithDefault    = []string{"id", "object_id", "file_size", "label", "created_at", "is_confirm"}
	fileServicePrimaryKeyColumns     = []string{"id"}
	fileServiceGeneratedColumns      = []string{}
)

type (
	// FileServiceSlice is an alias for a slice of pointers to FileService.
	// This should almost always be used instead of []FileService.
	FileServiceSlice []*FileService
	// FileServiceHook is the signature for custom FileService hook methods
	FileServiceHook func(context.Context, boil.ContextExecutor, *FileService) error

	fileServiceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fileServiceType                 = reflect.TypeOf(&FileService{})
	fileServiceMapping              = queries.MakeStructMapping(fileServiceType)
	fileServicePrimaryKeyMapping, _ = queries.BindMapping(fileServiceType, fileServiceMapping, fileServicePrimaryKeyColumns)
	fileServiceInsertCacheMut       sync.RWMutex
	fileServiceInsertCache          = make(map[string]insertCache)
	fileServiceUpdateCacheMut       sync.RWMutex
	fileServiceUpdateCache          = make(map[string]updateCache)
	fileServiceUpsertCacheMut       sync.RWMutex
	fileServiceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fileServiceAfterSelectHooks []FileServiceHook

var fileServiceBeforeInsertHooks []FileServiceHook
var fileServiceAfterInsertHooks []FileServiceHook

var fileServiceBeforeUpdateHooks []FileServiceHook
var fileServiceAfterUpdateHooks []FileServiceHook

var fileServiceBeforeDeleteHooks []FileServiceHook
var fileServiceAfterDeleteHooks []FileServiceHook

var fileServiceBeforeUpsertHooks []FileServiceHook
var fileServiceAfterUpsertHooks []FileServiceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FileService) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FileService) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FileService) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FileService) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FileService) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FileService) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FileService) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FileService) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FileService) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileServiceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFileServiceHook registers your hook function for all future operations.
func AddFileServiceHook(hookPoint boil.HookPoint, fileServiceHook FileServiceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		fileServiceAfterSelectHooks = append(fileServiceAfterSelectHooks, fileServiceHook)
	case boil.BeforeInsertHook:
		fileServiceBeforeInsertHooks = append(fileServiceBeforeInsertHooks, fileServiceHook)
	case boil.AfterInsertHook:
		fileServiceAfterInsertHooks = append(fileServiceAfterInsertHooks, fileServiceHook)
	case boil.BeforeUpdateHook:
		fileServiceBeforeUpdateHooks = append(fileServiceBeforeUpdateHooks, fileServiceHook)
	case boil.AfterUpdateHook:
		fileServiceAfterUpdateHooks = append(fileServiceAfterUpdateHooks, fileServiceHook)
	case boil.BeforeDeleteHook:
		fileServiceBeforeDeleteHooks = append(fileServiceBeforeDeleteHooks, fileServiceHook)
	case boil.AfterDeleteHook:
		fileServiceAfterDeleteHooks = append(fileServiceAfterDeleteHooks, fileServiceHook)
	case boil.BeforeUpsertHook:
		fileServiceBeforeUpsertHooks = append(fileServiceBeforeUpsertHooks, fileServiceHook)
	case boil.AfterUpsertHook:
		fileServiceAfterUpsertHooks = append(fileServiceAfterUpsertHooks, fileServiceHook)
	}
}

// One returns a single fileService record from the query.
func (q fileServiceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FileService, error) {
	o := &FileService{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "tp: failed to execute a one query for minio")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FileService records from the query.
func (q fileServiceQuery) All(ctx context.Context, exec boil.ContextExecutor) (FileServiceSlice, error) {
	var o []*FileService

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "tp: failed to assign all query results to FileService slice")
	}

	if len(fileServiceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FileService records in the query.
func (q fileServiceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "tp: failed to count minio rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fileServiceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "tp: failed to check if minio exists")
	}

	return count > 0, nil
}

// FileServices retrieves all the records using an executor.
func FileServices(mods ...qm.QueryMod) fileServiceQuery {
	mods = append(mods, qm.From("\"minio\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"minio\".*"})
	}

	return fileServiceQuery{q}
}

// FindFileService retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFileService(ctx context.Context, exec boil.ContextExecutor, iD uuid.UUID, selectCols ...string) (*FileService, error) {
	fileServiceObj := &FileService{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"minio\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fileServiceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "tp: unable to select from minio")
	}

	if err = fileServiceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return fileServiceObj, err
	}

	return fileServiceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FileService) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("tp: no minio provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileServiceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fileServiceInsertCacheMut.RLock()
	cache, cached := fileServiceInsertCache[key]
	fileServiceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fileServiceAllColumns,
			fileServiceColumnsWithDefault,
			fileServiceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fileServiceType, fileServiceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fileServiceType, fileServiceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"minio\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"minio\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "tp: unable to insert into minio")
	}

	if !cached {
		fileServiceInsertCacheMut.Lock()
		fileServiceInsertCache[key] = cache
		fileServiceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FileService.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FileService) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fileServiceUpdateCacheMut.RLock()
	cache, cached := fileServiceUpdateCache[key]
	fileServiceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fileServiceAllColumns,
			fileServicePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("tp: unable to update minio, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"minio\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, fileServicePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fileServiceType, fileServiceMapping, append(wl, fileServicePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to update minio row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: failed to get rows affected by update for minio")
	}

	if !cached {
		fileServiceUpdateCacheMut.Lock()
		fileServiceUpdateCache[key] = cache
		fileServiceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fileServiceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to update all for minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to retrieve rows affected for minio")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FileServiceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("tp: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"minio\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, fileServicePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to update all in fileService slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to retrieve rows affected all in update all fileService")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FileService) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("tp: no minio provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileServiceColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	fileServiceUpsertCacheMut.RLock()
	cache, cached := fileServiceUpsertCache[key]
	fileServiceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fileServiceAllColumns,
			fileServiceColumnsWithDefault,
			fileServiceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			fileServiceAllColumns,
			fileServicePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("tp: unable to upsert minio, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(fileServicePrimaryKeyColumns))
			copy(conflict, fileServicePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"minio\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(fileServiceType, fileServiceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fileServiceType, fileServiceMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "tp: unable to upsert minio")
	}

	if !cached {
		fileServiceUpsertCacheMut.Lock()
		fileServiceUpsertCache[key] = cache
		fileServiceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FileService record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FileService) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("tp: no FileService provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fileServicePrimaryKeyMapping)
	sql := "DELETE FROM \"minio\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to delete from minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: failed to get rows affected by delete for minio")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fileServiceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("tp: no fileServiceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to delete all from minio")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: failed to get rows affected by deleteall for minio")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FileServiceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fileServiceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"minio\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fileServicePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tp: unable to delete all from fileService slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tp: failed to get rows affected by deleteall for minio")
	}

	if len(fileServiceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FileService) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFileService(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FileServiceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FileServiceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fileServicePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"minio\".* FROM \"minio\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, fileServicePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "tp: unable to reload all in FileServiceSlice")
	}

	*o = slice

	return nil
}

// FileServiceExists checks if the FileService row exists.
func FileServiceExists(ctx context.Context, exec boil.ContextExecutor, iD uuid.UUID) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"minio\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "tp: unable to check if minio exists")
	}

	return exists, nil
}
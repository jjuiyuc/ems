// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package deremsmodels

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AiDatum is an object representing the database table.
type AiDatum struct {
	ID          int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GWUUID      string     `boil:"gw_uuid" json:"gwUUID" toml:"gwUUID" yaml:"gwUUID"`
	LogDate     time.Time  `boil:"log_date" json:"logDate" toml:"logDate" yaml:"logDate"`
	LogType     string     `boil:"log_type" json:"logType" toml:"logType" yaml:"logType"`
	GWID        null.Int64 `boil:"gw_id" json:"gwID,omitempty" toml:"gwID" yaml:"gwID,omitempty"`
	LocationID  null.Int64 `boil:"location_id" json:"locationID,omitempty" toml:"locationID" yaml:"locationID,omitempty"`
	LocalAiData null.JSON  `boil:"local_ai_data" json:"localAiData,omitempty" toml:"localAiData" yaml:"localAiData,omitempty"`
	CreatedAt   time.Time  `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt   time.Time  `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *aiDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L aiDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AiDatumColumns = struct {
	ID          string
	GWUUID      string
	LogDate     string
	LogType     string
	GWID        string
	LocationID  string
	LocalAiData string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	GWUUID:      "gw_uuid",
	LogDate:     "log_date",
	LogType:     "log_type",
	GWID:        "gw_id",
	LocationID:  "location_id",
	LocalAiData: "local_ai_data",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var AiDatumTableColumns = struct {
	ID          string
	GWUUID      string
	LogDate     string
	LogType     string
	GWID        string
	LocationID  string
	LocalAiData string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "ai_data.id",
	GWUUID:      "ai_data.gw_uuid",
	LogDate:     "ai_data.log_date",
	LogType:     "ai_data.log_type",
	GWID:        "ai_data.gw_id",
	LocationID:  "ai_data.location_id",
	LocalAiData: "ai_data.local_ai_data",
	CreatedAt:   "ai_data.created_at",
	UpdatedAt:   "ai_data.updated_at",
}

// Generated where

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

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var AiDatumWhere = struct {
	ID          whereHelperint64
	GWUUID      whereHelperstring
	LogDate     whereHelpertime_Time
	LogType     whereHelperstring
	GWID        whereHelpernull_Int64
	LocationID  whereHelpernull_Int64
	LocalAiData whereHelpernull_JSON
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "`ai_data`.`id`"},
	GWUUID:      whereHelperstring{field: "`ai_data`.`gw_uuid`"},
	LogDate:     whereHelpertime_Time{field: "`ai_data`.`log_date`"},
	LogType:     whereHelperstring{field: "`ai_data`.`log_type`"},
	GWID:        whereHelpernull_Int64{field: "`ai_data`.`gw_id`"},
	LocationID:  whereHelpernull_Int64{field: "`ai_data`.`location_id`"},
	LocalAiData: whereHelpernull_JSON{field: "`ai_data`.`local_ai_data`"},
	CreatedAt:   whereHelpertime_Time{field: "`ai_data`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`ai_data`.`updated_at`"},
}

// AiDatumRels is where relationship names are stored.
var AiDatumRels = struct {
}{}

// aiDatumR is where relationships are stored.
type aiDatumR struct {
}

// NewStruct creates a new relationship struct
func (*aiDatumR) NewStruct() *aiDatumR {
	return &aiDatumR{}
}

// aiDatumL is where Load methods for each relationship are stored.
type aiDatumL struct{}

var (
	aiDatumAllColumns            = []string{"id", "gw_uuid", "log_date", "log_type", "gw_id", "location_id", "local_ai_data", "created_at", "updated_at"}
	aiDatumColumnsWithoutDefault = []string{"gw_uuid", "log_date", "log_type", "gw_id", "location_id", "local_ai_data"}
	aiDatumColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	aiDatumPrimaryKeyColumns     = []string{"id"}
	aiDatumGeneratedColumns      = []string{}
)

type (
	// AiDatumSlice is an alias for a slice of pointers to AiDatum.
	// This should almost always be used instead of []AiDatum.
	AiDatumSlice []*AiDatum

	aiDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	aiDatumType                 = reflect.TypeOf(&AiDatum{})
	aiDatumMapping              = queries.MakeStructMapping(aiDatumType)
	aiDatumPrimaryKeyMapping, _ = queries.BindMapping(aiDatumType, aiDatumMapping, aiDatumPrimaryKeyColumns)
	aiDatumInsertCacheMut       sync.RWMutex
	aiDatumInsertCache          = make(map[string]insertCache)
	aiDatumUpdateCacheMut       sync.RWMutex
	aiDatumUpdateCache          = make(map[string]updateCache)
	aiDatumUpsertCacheMut       sync.RWMutex
	aiDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single aiDatum record from the query.
func (q aiDatumQuery) One(exec boil.Executor) (*AiDatum, error) {
	o := &AiDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for ai_data")
	}

	return o, nil
}

// All returns all AiDatum records from the query.
func (q aiDatumQuery) All(exec boil.Executor) (AiDatumSlice, error) {
	var o []*AiDatum

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to AiDatum slice")
	}

	return o, nil
}

// Count returns the count of all AiDatum records in the query.
func (q aiDatumQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count ai_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q aiDatumQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if ai_data exists")
	}

	return count > 0, nil
}

// AiData retrieves all the records using an executor.
func AiData(mods ...qm.QueryMod) aiDatumQuery {
	mods = append(mods, qm.From("`ai_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`ai_data`.*"})
	}

	return aiDatumQuery{q}
}

// FindAiDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAiDatum(exec boil.Executor, iD int64, selectCols ...string) (*AiDatum, error) {
	aiDatumObj := &AiDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `ai_data` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, aiDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from ai_data")
	}

	return aiDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AiDatum) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no ai_data provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(aiDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	aiDatumInsertCacheMut.RLock()
	cache, cached := aiDatumInsertCache[key]
	aiDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			aiDatumAllColumns,
			aiDatumColumnsWithDefault,
			aiDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(aiDatumType, aiDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(aiDatumType, aiDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `ai_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `ai_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `ai_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, aiDatumPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to insert into ai_data")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == aiDatumMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for ai_data")
	}

CacheNoHooks:
	if !cached {
		aiDatumInsertCacheMut.Lock()
		aiDatumInsertCache[key] = cache
		aiDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AiDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AiDatum) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	aiDatumUpdateCacheMut.RLock()
	cache, cached := aiDatumUpdateCache[key]
	aiDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			aiDatumAllColumns,
			aiDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update ai_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `ai_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, aiDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(aiDatumType, aiDatumMapping, append(wl, aiDatumPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update ai_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for ai_data")
	}

	if !cached {
		aiDatumUpdateCacheMut.Lock()
		aiDatumUpdateCache[key] = cache
		aiDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q aiDatumQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for ai_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for ai_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AiDatumSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("deremsmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aiDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `ai_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aiDatumPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in aiDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all aiDatum")
	}
	return rowsAff, nil
}

var mySQLAiDatumUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AiDatum) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no ai_data provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(aiDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAiDatumUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	aiDatumUpsertCacheMut.RLock()
	cache, cached := aiDatumUpsertCache[key]
	aiDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			aiDatumAllColumns,
			aiDatumColumnsWithDefault,
			aiDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			aiDatumAllColumns,
			aiDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert ai_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`ai_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `ai_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(aiDatumType, aiDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(aiDatumType, aiDatumMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to upsert for ai_data")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == aiDatumMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(aiDatumType, aiDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for ai_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for ai_data")
	}

CacheNoHooks:
	if !cached {
		aiDatumUpsertCacheMut.Lock()
		aiDatumUpsertCache[key] = cache
		aiDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AiDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AiDatum) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no AiDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), aiDatumPrimaryKeyMapping)
	sql := "DELETE FROM `ai_data` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from ai_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for ai_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q aiDatumQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no aiDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from ai_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for ai_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AiDatumSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aiDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `ai_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aiDatumPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from aiDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for ai_data")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AiDatum) Reload(exec boil.Executor) error {
	ret, err := FindAiDatum(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AiDatumSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AiDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aiDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `ai_data`.* FROM `ai_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aiDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in AiDatumSlice")
	}

	*o = slice

	return nil
}

// AiDatumExists checks if the AiDatum row exists.
func AiDatumExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `ai_data` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if ai_data exists")
	}

	return exists, nil
}

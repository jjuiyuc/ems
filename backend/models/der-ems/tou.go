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

// Tou is an object representing the database table.
type Tou struct {
	ID            int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	TOULocationID null.Int     `boil:"tou_location_id" json:"touLocationID,omitempty" toml:"touLocationID" yaml:"touLocationID,omitempty"`
	VoltageType   null.String  `boil:"voltage_type" json:"voltageType,omitempty" toml:"voltageType" yaml:"voltageType,omitempty"`
	TOUType       null.String  `boil:"tou_type" json:"touType,omitempty" toml:"touType" yaml:"touType,omitempty"`
	PeriodType    null.String  `boil:"period_type" json:"periodType,omitempty" toml:"periodType" yaml:"periodType,omitempty"`
	PeakType      null.String  `boil:"peak_type" json:"peakType,omitempty" toml:"peakType" yaml:"peakType,omitempty"`
	IsSummer      null.Bool    `boil:"is_summer" json:"isSummer,omitempty" toml:"isSummer" yaml:"isSummer,omitempty"`
	PeriodStime   null.String  `boil:"period_stime" json:"periodStime,omitempty" toml:"periodStime" yaml:"periodStime,omitempty"`
	PeriodEtime   null.String  `boil:"period_etime" json:"periodEtime,omitempty" toml:"periodEtime" yaml:"periodEtime,omitempty"`
	BasicCharge   null.Float32 `boil:"basic_charge" json:"basicCharge,omitempty" toml:"basicCharge" yaml:"basicCharge,omitempty"`
	BasicRate     null.Float32 `boil:"basic_rate" json:"basicRate,omitempty" toml:"basicRate" yaml:"basicRate,omitempty"`
	FlowRate      null.Float32 `boil:"flow_rate" json:"flowRate,omitempty" toml:"flowRate" yaml:"flowRate,omitempty"`
	EnableAt      null.Time    `boil:"enable_at" json:"enableAt,omitempty" toml:"enableAt" yaml:"enableAt,omitempty"`
	DisableAt     null.Time    `boil:"disable_at" json:"disableAt,omitempty" toml:"disableAt" yaml:"disableAt,omitempty"`
	CreatedAt     time.Time    `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt     null.Time    `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`

	R *touR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L touL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TouColumns = struct {
	ID            string
	TOULocationID string
	VoltageType   string
	TOUType       string
	PeriodType    string
	PeakType      string
	IsSummer      string
	PeriodStime   string
	PeriodEtime   string
	BasicCharge   string
	BasicRate     string
	FlowRate      string
	EnableAt      string
	DisableAt     string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	TOULocationID: "tou_location_id",
	VoltageType:   "voltage_type",
	TOUType:       "tou_type",
	PeriodType:    "period_type",
	PeakType:      "peak_type",
	IsSummer:      "is_summer",
	PeriodStime:   "period_stime",
	PeriodEtime:   "period_etime",
	BasicCharge:   "basic_charge",
	BasicRate:     "basic_rate",
	FlowRate:      "flow_rate",
	EnableAt:      "enable_at",
	DisableAt:     "disable_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var TouTableColumns = struct {
	ID            string
	TOULocationID string
	VoltageType   string
	TOUType       string
	PeriodType    string
	PeakType      string
	IsSummer      string
	PeriodStime   string
	PeriodEtime   string
	BasicCharge   string
	BasicRate     string
	FlowRate      string
	EnableAt      string
	DisableAt     string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "tou.id",
	TOULocationID: "tou.tou_location_id",
	VoltageType:   "tou.voltage_type",
	TOUType:       "tou.tou_type",
	PeriodType:    "tou.period_type",
	PeakType:      "tou.peak_type",
	IsSummer:      "tou.is_summer",
	PeriodStime:   "tou.period_stime",
	PeriodEtime:   "tou.period_etime",
	BasicCharge:   "tou.basic_charge",
	BasicRate:     "tou.basic_rate",
	FlowRate:      "tou.flow_rate",
	EnableAt:      "tou.enable_at",
	DisableAt:     "tou.disable_at",
	CreatedAt:     "tou.created_at",
	UpdatedAt:     "tou.updated_at",
}

// Generated where

var TouWhere = struct {
	ID            whereHelperint
	TOULocationID whereHelpernull_Int
	VoltageType   whereHelpernull_String
	TOUType       whereHelpernull_String
	PeriodType    whereHelpernull_String
	PeakType      whereHelpernull_String
	IsSummer      whereHelpernull_Bool
	PeriodStime   whereHelpernull_String
	PeriodEtime   whereHelpernull_String
	BasicCharge   whereHelpernull_Float32
	BasicRate     whereHelpernull_Float32
	FlowRate      whereHelpernull_Float32
	EnableAt      whereHelpernull_Time
	DisableAt     whereHelpernull_Time
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpernull_Time
}{
	ID:            whereHelperint{field: "`tou`.`id`"},
	TOULocationID: whereHelpernull_Int{field: "`tou`.`tou_location_id`"},
	VoltageType:   whereHelpernull_String{field: "`tou`.`voltage_type`"},
	TOUType:       whereHelpernull_String{field: "`tou`.`tou_type`"},
	PeriodType:    whereHelpernull_String{field: "`tou`.`period_type`"},
	PeakType:      whereHelpernull_String{field: "`tou`.`peak_type`"},
	IsSummer:      whereHelpernull_Bool{field: "`tou`.`is_summer`"},
	PeriodStime:   whereHelpernull_String{field: "`tou`.`period_stime`"},
	PeriodEtime:   whereHelpernull_String{field: "`tou`.`period_etime`"},
	BasicCharge:   whereHelpernull_Float32{field: "`tou`.`basic_charge`"},
	BasicRate:     whereHelpernull_Float32{field: "`tou`.`basic_rate`"},
	FlowRate:      whereHelpernull_Float32{field: "`tou`.`flow_rate`"},
	EnableAt:      whereHelpernull_Time{field: "`tou`.`enable_at`"},
	DisableAt:     whereHelpernull_Time{field: "`tou`.`disable_at`"},
	CreatedAt:     whereHelpertime_Time{field: "`tou`.`created_at`"},
	UpdatedAt:     whereHelpernull_Time{field: "`tou`.`updated_at`"},
}

// TouRels is where relationship names are stored.
var TouRels = struct {
}{}

// touR is where relationships are stored.
type touR struct {
}

// NewStruct creates a new relationship struct
func (*touR) NewStruct() *touR {
	return &touR{}
}

// touL is where Load methods for each relationship are stored.
type touL struct{}

var (
	touAllColumns            = []string{"id", "tou_location_id", "voltage_type", "tou_type", "period_type", "peak_type", "is_summer", "period_stime", "period_etime", "basic_charge", "basic_rate", "flow_rate", "enable_at", "disable_at", "created_at", "updated_at"}
	touColumnsWithoutDefault = []string{"tou_location_id", "voltage_type", "tou_type", "period_type", "peak_type", "is_summer", "period_stime", "period_etime", "basic_charge", "basic_rate", "flow_rate", "enable_at", "disable_at", "updated_at"}
	touColumnsWithDefault    = []string{"id", "created_at"}
	touPrimaryKeyColumns     = []string{"id"}
	touGeneratedColumns      = []string{}
)

type (
	// TouSlice is an alias for a slice of pointers to Tou.
	// This should almost always be used instead of []Tou.
	TouSlice []*Tou

	touQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	touType                 = reflect.TypeOf(&Tou{})
	touMapping              = queries.MakeStructMapping(touType)
	touPrimaryKeyMapping, _ = queries.BindMapping(touType, touMapping, touPrimaryKeyColumns)
	touInsertCacheMut       sync.RWMutex
	touInsertCache          = make(map[string]insertCache)
	touUpdateCacheMut       sync.RWMutex
	touUpdateCache          = make(map[string]updateCache)
	touUpsertCacheMut       sync.RWMutex
	touUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single tou record from the query.
func (q touQuery) One(exec boil.Executor) (*Tou, error) {
	o := &Tou{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for tou")
	}

	return o, nil
}

// All returns all Tou records from the query.
func (q touQuery) All(exec boil.Executor) (TouSlice, error) {
	var o []*Tou

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to Tou slice")
	}

	return o, nil
}

// Count returns the count of all Tou records in the query.
func (q touQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count tou rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q touQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if tou exists")
	}

	return count > 0, nil
}

// Tous retrieves all the records using an executor.
func Tous(mods ...qm.QueryMod) touQuery {
	mods = append(mods, qm.From("`tou`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tou`.*"})
	}

	return touQuery{q}
}

// FindTou retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTou(exec boil.Executor, iD int, selectCols ...string) (*Tou, error) {
	touObj := &Tou{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tou` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, touObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from tou")
	}

	return touObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tou) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no tou provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(touColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	touInsertCacheMut.RLock()
	cache, cached := touInsertCache[key]
	touInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			touAllColumns,
			touColumnsWithDefault,
			touColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(touType, touMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(touType, touMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tou` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tou` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tou` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, touPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into tou")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == touMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for tou")
	}

CacheNoHooks:
	if !cached {
		touInsertCacheMut.Lock()
		touInsertCache[key] = cache
		touInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Tou.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tou) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	key := makeCacheKey(columns, nil)
	touUpdateCacheMut.RLock()
	cache, cached := touUpdateCache[key]
	touUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			touAllColumns,
			touPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update tou, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tou` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, touPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(touType, touMapping, append(wl, touPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update tou row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for tou")
	}

	if !cached {
		touUpdateCacheMut.Lock()
		touUpdateCache[key] = cache
		touUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q touQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for tou")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for tou")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TouSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tou` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in tou slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all tou")
	}
	return rowsAff, nil
}

var mySQLTouUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tou) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no tou provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	nzDefaults := queries.NonZeroDefaultSet(touColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTouUniqueColumns, o)

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

	touUpsertCacheMut.RLock()
	cache, cached := touUpsertCache[key]
	touUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			touAllColumns,
			touColumnsWithDefault,
			touColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			touAllColumns,
			touPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert tou, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`tou`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tou` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(touType, touMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(touType, touMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for tou")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == touMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(touType, touMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for tou")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for tou")
	}

CacheNoHooks:
	if !cached {
		touUpsertCacheMut.Lock()
		touUpsertCache[key] = cache
		touUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Tou record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tou) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no Tou provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), touPrimaryKeyMapping)
	sql := "DELETE FROM `tou` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from tou")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for tou")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q touQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no touQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from tou")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for tou")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TouSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tou` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from tou slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for tou")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tou) Reload(exec boil.Executor) error {
	ret, err := FindTou(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TouSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TouSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tou`.* FROM `tou` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in TouSlice")
	}

	*o = slice

	return nil
}

// TouExists checks if the Tou row exists.
func TouExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tou` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if tou exists")
	}

	return exists, nil
}

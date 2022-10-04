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

// TouHoliday is an object representing the database table.
type TouHoliday struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TOULocationID null.Int    `boil:"tou_location_id" json:"touLocationID,omitempty" toml:"touLocationID" yaml:"touLocationID,omitempty"`
	Year          null.String `boil:"year" json:"year,omitempty" toml:"year" yaml:"year,omitempty"`
	Day           null.Time   `boil:"day" json:"day,omitempty" toml:"day" yaml:"day,omitempty"`
	CreatedAt     time.Time   `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`

	R *touHolidayR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L touHolidayL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TouHolidayColumns = struct {
	ID            string
	TOULocationID string
	Year          string
	Day           string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	TOULocationID: "tou_location_id",
	Year:          "year",
	Day:           "day",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var TouHolidayTableColumns = struct {
	ID            string
	TOULocationID string
	Year          string
	Day           string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "tou_holiday.id",
	TOULocationID: "tou_holiday.tou_location_id",
	Year:          "tou_holiday.year",
	Day:           "tou_holiday.day",
	CreatedAt:     "tou_holiday.created_at",
	UpdatedAt:     "tou_holiday.updated_at",
}

// Generated where

var TouHolidayWhere = struct {
	ID            whereHelperint64
	TOULocationID whereHelpernull_Int
	Year          whereHelpernull_String
	Day           whereHelpernull_Time
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpernull_Time
}{
	ID:            whereHelperint64{field: "`tou_holiday`.`id`"},
	TOULocationID: whereHelpernull_Int{field: "`tou_holiday`.`tou_location_id`"},
	Year:          whereHelpernull_String{field: "`tou_holiday`.`year`"},
	Day:           whereHelpernull_Time{field: "`tou_holiday`.`day`"},
	CreatedAt:     whereHelpertime_Time{field: "`tou_holiday`.`created_at`"},
	UpdatedAt:     whereHelpernull_Time{field: "`tou_holiday`.`updated_at`"},
}

// TouHolidayRels is where relationship names are stored.
var TouHolidayRels = struct {
}{}

// touHolidayR is where relationships are stored.
type touHolidayR struct {
}

// NewStruct creates a new relationship struct
func (*touHolidayR) NewStruct() *touHolidayR {
	return &touHolidayR{}
}

// touHolidayL is where Load methods for each relationship are stored.
type touHolidayL struct{}

var (
	touHolidayAllColumns            = []string{"id", "tou_location_id", "year", "day", "created_at", "updated_at"}
	touHolidayColumnsWithoutDefault = []string{"tou_location_id", "year", "day", "updated_at"}
	touHolidayColumnsWithDefault    = []string{"id", "created_at"}
	touHolidayPrimaryKeyColumns     = []string{"id"}
	touHolidayGeneratedColumns      = []string{}
)

type (
	// TouHolidaySlice is an alias for a slice of pointers to TouHoliday.
	// This should almost always be used instead of []TouHoliday.
	TouHolidaySlice []*TouHoliday

	touHolidayQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	touHolidayType                 = reflect.TypeOf(&TouHoliday{})
	touHolidayMapping              = queries.MakeStructMapping(touHolidayType)
	touHolidayPrimaryKeyMapping, _ = queries.BindMapping(touHolidayType, touHolidayMapping, touHolidayPrimaryKeyColumns)
	touHolidayInsertCacheMut       sync.RWMutex
	touHolidayInsertCache          = make(map[string]insertCache)
	touHolidayUpdateCacheMut       sync.RWMutex
	touHolidayUpdateCache          = make(map[string]updateCache)
	touHolidayUpsertCacheMut       sync.RWMutex
	touHolidayUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single touHoliday record from the query.
func (q touHolidayQuery) One(exec boil.Executor) (*TouHoliday, error) {
	o := &TouHoliday{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for tou_holiday")
	}

	return o, nil
}

// All returns all TouHoliday records from the query.
func (q touHolidayQuery) All(exec boil.Executor) (TouHolidaySlice, error) {
	var o []*TouHoliday

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to TouHoliday slice")
	}

	return o, nil
}

// Count returns the count of all TouHoliday records in the query.
func (q touHolidayQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count tou_holiday rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q touHolidayQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if tou_holiday exists")
	}

	return count > 0, nil
}

// TouHolidays retrieves all the records using an executor.
func TouHolidays(mods ...qm.QueryMod) touHolidayQuery {
	mods = append(mods, qm.From("`tou_holiday`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tou_holiday`.*"})
	}

	return touHolidayQuery{q}
}

// FindTouHoliday retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTouHoliday(exec boil.Executor, iD int64, selectCols ...string) (*TouHoliday, error) {
	touHolidayObj := &TouHoliday{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tou_holiday` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, touHolidayObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from tou_holiday")
	}

	return touHolidayObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TouHoliday) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no tou_holiday provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(touHolidayColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	touHolidayInsertCacheMut.RLock()
	cache, cached := touHolidayInsertCache[key]
	touHolidayInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			touHolidayAllColumns,
			touHolidayColumnsWithDefault,
			touHolidayColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(touHolidayType, touHolidayMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(touHolidayType, touHolidayMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tou_holiday` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tou_holiday` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tou_holiday` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, touHolidayPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into tou_holiday")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == touHolidayMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for tou_holiday")
	}

CacheNoHooks:
	if !cached {
		touHolidayInsertCacheMut.Lock()
		touHolidayInsertCache[key] = cache
		touHolidayInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the TouHoliday.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TouHoliday) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	key := makeCacheKey(columns, nil)
	touHolidayUpdateCacheMut.RLock()
	cache, cached := touHolidayUpdateCache[key]
	touHolidayUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			touHolidayAllColumns,
			touHolidayPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update tou_holiday, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tou_holiday` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, touHolidayPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(touHolidayType, touHolidayMapping, append(wl, touHolidayPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update tou_holiday row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for tou_holiday")
	}

	if !cached {
		touHolidayUpdateCacheMut.Lock()
		touHolidayUpdateCache[key] = cache
		touHolidayUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q touHolidayQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for tou_holiday")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for tou_holiday")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TouHolidaySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touHolidayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tou_holiday` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touHolidayPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in touHoliday slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all touHoliday")
	}
	return rowsAff, nil
}

var mySQLTouHolidayUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TouHoliday) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no tou_holiday provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	nzDefaults := queries.NonZeroDefaultSet(touHolidayColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTouHolidayUniqueColumns, o)

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

	touHolidayUpsertCacheMut.RLock()
	cache, cached := touHolidayUpsertCache[key]
	touHolidayUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			touHolidayAllColumns,
			touHolidayColumnsWithDefault,
			touHolidayColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			touHolidayAllColumns,
			touHolidayPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert tou_holiday, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`tou_holiday`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tou_holiday` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(touHolidayType, touHolidayMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(touHolidayType, touHolidayMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for tou_holiday")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == touHolidayMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(touHolidayType, touHolidayMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for tou_holiday")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for tou_holiday")
	}

CacheNoHooks:
	if !cached {
		touHolidayUpsertCacheMut.Lock()
		touHolidayUpsertCache[key] = cache
		touHolidayUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single TouHoliday record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TouHoliday) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no TouHoliday provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), touHolidayPrimaryKeyMapping)
	sql := "DELETE FROM `tou_holiday` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from tou_holiday")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for tou_holiday")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q touHolidayQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no touHolidayQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from tou_holiday")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for tou_holiday")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TouHolidaySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touHolidayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tou_holiday` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touHolidayPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from touHoliday slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for tou_holiday")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TouHoliday) Reload(exec boil.Executor) error {
	ret, err := FindTouHoliday(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TouHolidaySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TouHolidaySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), touHolidayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tou_holiday`.* FROM `tou_holiday` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, touHolidayPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in TouHolidaySlice")
	}

	*o = slice

	return nil
}

// TouHolidayExists checks if the TouHoliday row exists.
func TouHolidayExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tou_holiday` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if tou_holiday exists")
	}

	return exists, nil
}

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

// GroupGatewayRightLog is an object representing the database table.
type GroupGatewayRightLog struct {
	ID                         int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GroupGatewayRightID        null.Int64 `boil:"group_gateway_right_id" json:"groupGatewayRightID,omitempty" toml:"groupGatewayRightID" yaml:"groupGatewayRightID,omitempty"`
	GroupID                    null.Int64 `boil:"group_id" json:"groupID,omitempty" toml:"groupID" yaml:"groupID,omitempty"`
	GWID                       null.Int64 `boil:"gw_id" json:"gwID,omitempty" toml:"gwID" yaml:"gwID,omitempty"`
	EnabledAt                  null.Time  `boil:"enabled_at" json:"enabledAt,omitempty" toml:"enabledAt" yaml:"enabledAt,omitempty"`
	EnabledBy                  null.Int64 `boil:"enabled_by" json:"enabledBy,omitempty" toml:"enabledBy" yaml:"enabledBy,omitempty"`
	DisabledAt                 null.Time  `boil:"disabled_at" json:"disabledAt,omitempty" toml:"disabledAt" yaml:"disabledAt,omitempty"`
	DisabledBy                 null.Int64 `boil:"disabled_by" json:"disabledBy,omitempty" toml:"disabledBy" yaml:"disabledBy,omitempty"`
	GroupGatewayRightUpdatedAt null.Time  `boil:"group_gateway_right_updated_at" json:"groupGatewayRightUpdatedAt,omitempty" toml:"groupGatewayRightUpdatedAt" yaml:"groupGatewayRightUpdatedAt,omitempty"`
	GroupGatewayRightUpdatedBy null.Int64 `boil:"group_gateway_right_updated_by" json:"groupGatewayRightUpdatedBy,omitempty" toml:"groupGatewayRightUpdatedBy" yaml:"groupGatewayRightUpdatedBy,omitempty"`
	CreatedAt                  time.Time  `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *groupGatewayRightLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupGatewayRightLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupGatewayRightLogColumns = struct {
	ID                         string
	GroupGatewayRightID        string
	GroupID                    string
	GWID                       string
	EnabledAt                  string
	EnabledBy                  string
	DisabledAt                 string
	DisabledBy                 string
	GroupGatewayRightUpdatedAt string
	GroupGatewayRightUpdatedBy string
	CreatedAt                  string
}{
	ID:                         "id",
	GroupGatewayRightID:        "group_gateway_right_id",
	GroupID:                    "group_id",
	GWID:                       "gw_id",
	EnabledAt:                  "enabled_at",
	EnabledBy:                  "enabled_by",
	DisabledAt:                 "disabled_at",
	DisabledBy:                 "disabled_by",
	GroupGatewayRightUpdatedAt: "group_gateway_right_updated_at",
	GroupGatewayRightUpdatedBy: "group_gateway_right_updated_by",
	CreatedAt:                  "created_at",
}

var GroupGatewayRightLogTableColumns = struct {
	ID                         string
	GroupGatewayRightID        string
	GroupID                    string
	GWID                       string
	EnabledAt                  string
	EnabledBy                  string
	DisabledAt                 string
	DisabledBy                 string
	GroupGatewayRightUpdatedAt string
	GroupGatewayRightUpdatedBy string
	CreatedAt                  string
}{
	ID:                         "group_gateway_right_log.id",
	GroupGatewayRightID:        "group_gateway_right_log.group_gateway_right_id",
	GroupID:                    "group_gateway_right_log.group_id",
	GWID:                       "group_gateway_right_log.gw_id",
	EnabledAt:                  "group_gateway_right_log.enabled_at",
	EnabledBy:                  "group_gateway_right_log.enabled_by",
	DisabledAt:                 "group_gateway_right_log.disabled_at",
	DisabledBy:                 "group_gateway_right_log.disabled_by",
	GroupGatewayRightUpdatedAt: "group_gateway_right_log.group_gateway_right_updated_at",
	GroupGatewayRightUpdatedBy: "group_gateway_right_log.group_gateway_right_updated_by",
	CreatedAt:                  "group_gateway_right_log.created_at",
}

// Generated where

var GroupGatewayRightLogWhere = struct {
	ID                         whereHelperint64
	GroupGatewayRightID        whereHelpernull_Int64
	GroupID                    whereHelpernull_Int64
	GWID                       whereHelpernull_Int64
	EnabledAt                  whereHelpernull_Time
	EnabledBy                  whereHelpernull_Int64
	DisabledAt                 whereHelpernull_Time
	DisabledBy                 whereHelpernull_Int64
	GroupGatewayRightUpdatedAt whereHelpernull_Time
	GroupGatewayRightUpdatedBy whereHelpernull_Int64
	CreatedAt                  whereHelpertime_Time
}{
	ID:                         whereHelperint64{field: "`group_gateway_right_log`.`id`"},
	GroupGatewayRightID:        whereHelpernull_Int64{field: "`group_gateway_right_log`.`group_gateway_right_id`"},
	GroupID:                    whereHelpernull_Int64{field: "`group_gateway_right_log`.`group_id`"},
	GWID:                       whereHelpernull_Int64{field: "`group_gateway_right_log`.`gw_id`"},
	EnabledAt:                  whereHelpernull_Time{field: "`group_gateway_right_log`.`enabled_at`"},
	EnabledBy:                  whereHelpernull_Int64{field: "`group_gateway_right_log`.`enabled_by`"},
	DisabledAt:                 whereHelpernull_Time{field: "`group_gateway_right_log`.`disabled_at`"},
	DisabledBy:                 whereHelpernull_Int64{field: "`group_gateway_right_log`.`disabled_by`"},
	GroupGatewayRightUpdatedAt: whereHelpernull_Time{field: "`group_gateway_right_log`.`group_gateway_right_updated_at`"},
	GroupGatewayRightUpdatedBy: whereHelpernull_Int64{field: "`group_gateway_right_log`.`group_gateway_right_updated_by`"},
	CreatedAt:                  whereHelpertime_Time{field: "`group_gateway_right_log`.`created_at`"},
}

// GroupGatewayRightLogRels is where relationship names are stored.
var GroupGatewayRightLogRels = struct {
}{}

// groupGatewayRightLogR is where relationships are stored.
type groupGatewayRightLogR struct {
}

// NewStruct creates a new relationship struct
func (*groupGatewayRightLogR) NewStruct() *groupGatewayRightLogR {
	return &groupGatewayRightLogR{}
}

// groupGatewayRightLogL is where Load methods for each relationship are stored.
type groupGatewayRightLogL struct{}

var (
	groupGatewayRightLogAllColumns            = []string{"id", "group_gateway_right_id", "group_id", "gw_id", "enabled_at", "enabled_by", "disabled_at", "disabled_by", "group_gateway_right_updated_at", "group_gateway_right_updated_by", "created_at"}
	groupGatewayRightLogColumnsWithoutDefault = []string{"group_gateway_right_id", "group_id", "gw_id", "enabled_at", "enabled_by", "disabled_at", "disabled_by", "group_gateway_right_updated_at", "group_gateway_right_updated_by"}
	groupGatewayRightLogColumnsWithDefault    = []string{"id", "created_at"}
	groupGatewayRightLogPrimaryKeyColumns     = []string{"id"}
	groupGatewayRightLogGeneratedColumns      = []string{}
)

type (
	// GroupGatewayRightLogSlice is an alias for a slice of pointers to GroupGatewayRightLog.
	// This should almost always be used instead of []GroupGatewayRightLog.
	GroupGatewayRightLogSlice []*GroupGatewayRightLog

	groupGatewayRightLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupGatewayRightLogType                 = reflect.TypeOf(&GroupGatewayRightLog{})
	groupGatewayRightLogMapping              = queries.MakeStructMapping(groupGatewayRightLogType)
	groupGatewayRightLogPrimaryKeyMapping, _ = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, groupGatewayRightLogPrimaryKeyColumns)
	groupGatewayRightLogInsertCacheMut       sync.RWMutex
	groupGatewayRightLogInsertCache          = make(map[string]insertCache)
	groupGatewayRightLogUpdateCacheMut       sync.RWMutex
	groupGatewayRightLogUpdateCache          = make(map[string]updateCache)
	groupGatewayRightLogUpsertCacheMut       sync.RWMutex
	groupGatewayRightLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single groupGatewayRightLog record from the query.
func (q groupGatewayRightLogQuery) One(exec boil.Executor) (*GroupGatewayRightLog, error) {
	o := &GroupGatewayRightLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for group_gateway_right_log")
	}

	return o, nil
}

// All returns all GroupGatewayRightLog records from the query.
func (q groupGatewayRightLogQuery) All(exec boil.Executor) (GroupGatewayRightLogSlice, error) {
	var o []*GroupGatewayRightLog

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to GroupGatewayRightLog slice")
	}

	return o, nil
}

// Count returns the count of all GroupGatewayRightLog records in the query.
func (q groupGatewayRightLogQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count group_gateway_right_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupGatewayRightLogQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if group_gateway_right_log exists")
	}

	return count > 0, nil
}

// GroupGatewayRightLogs retrieves all the records using an executor.
func GroupGatewayRightLogs(mods ...qm.QueryMod) groupGatewayRightLogQuery {
	mods = append(mods, qm.From("`group_gateway_right_log`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`group_gateway_right_log`.*"})
	}

	return groupGatewayRightLogQuery{q}
}

// FindGroupGatewayRightLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroupGatewayRightLog(exec boil.Executor, iD int64, selectCols ...string) (*GroupGatewayRightLog, error) {
	groupGatewayRightLogObj := &GroupGatewayRightLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `group_gateway_right_log` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, groupGatewayRightLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from group_gateway_right_log")
	}

	return groupGatewayRightLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GroupGatewayRightLog) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no group_gateway_right_log provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(groupGatewayRightLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupGatewayRightLogInsertCacheMut.RLock()
	cache, cached := groupGatewayRightLogInsertCache[key]
	groupGatewayRightLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupGatewayRightLogAllColumns,
			groupGatewayRightLogColumnsWithDefault,
			groupGatewayRightLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `group_gateway_right_log` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `group_gateway_right_log` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `group_gateway_right_log` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, groupGatewayRightLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into group_gateway_right_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupGatewayRightLogMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for group_gateway_right_log")
	}

CacheNoHooks:
	if !cached {
		groupGatewayRightLogInsertCacheMut.Lock()
		groupGatewayRightLogInsertCache[key] = cache
		groupGatewayRightLogInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the GroupGatewayRightLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GroupGatewayRightLog) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	groupGatewayRightLogUpdateCacheMut.RLock()
	cache, cached := groupGatewayRightLogUpdateCache[key]
	groupGatewayRightLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupGatewayRightLogAllColumns,
			groupGatewayRightLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update group_gateway_right_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `group_gateway_right_log` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, groupGatewayRightLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, append(wl, groupGatewayRightLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update group_gateway_right_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for group_gateway_right_log")
	}

	if !cached {
		groupGatewayRightLogUpdateCacheMut.Lock()
		groupGatewayRightLogUpdateCache[key] = cache
		groupGatewayRightLogUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q groupGatewayRightLogQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for group_gateway_right_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for group_gateway_right_log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupGatewayRightLogSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupGatewayRightLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `group_gateway_right_log` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupGatewayRightLogPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in groupGatewayRightLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all groupGatewayRightLog")
	}
	return rowsAff, nil
}

var mySQLGroupGatewayRightLogUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GroupGatewayRightLog) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no group_gateway_right_log provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(groupGatewayRightLogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGroupGatewayRightLogUniqueColumns, o)

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

	groupGatewayRightLogUpsertCacheMut.RLock()
	cache, cached := groupGatewayRightLogUpsertCache[key]
	groupGatewayRightLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupGatewayRightLogAllColumns,
			groupGatewayRightLogColumnsWithDefault,
			groupGatewayRightLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			groupGatewayRightLogAllColumns,
			groupGatewayRightLogPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert group_gateway_right_log, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`group_gateway_right_log`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `group_gateway_right_log` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for group_gateway_right_log")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupGatewayRightLogMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(groupGatewayRightLogType, groupGatewayRightLogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for group_gateway_right_log")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for group_gateway_right_log")
	}

CacheNoHooks:
	if !cached {
		groupGatewayRightLogUpsertCacheMut.Lock()
		groupGatewayRightLogUpsertCache[key] = cache
		groupGatewayRightLogUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single GroupGatewayRightLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GroupGatewayRightLog) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no GroupGatewayRightLog provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupGatewayRightLogPrimaryKeyMapping)
	sql := "DELETE FROM `group_gateway_right_log` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from group_gateway_right_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for group_gateway_right_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupGatewayRightLogQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no groupGatewayRightLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from group_gateway_right_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for group_gateway_right_log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupGatewayRightLogSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupGatewayRightLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `group_gateway_right_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupGatewayRightLogPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from groupGatewayRightLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for group_gateway_right_log")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GroupGatewayRightLog) Reload(exec boil.Executor) error {
	ret, err := FindGroupGatewayRightLog(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupGatewayRightLogSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupGatewayRightLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupGatewayRightLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `group_gateway_right_log`.* FROM `group_gateway_right_log` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupGatewayRightLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in GroupGatewayRightLogSlice")
	}

	*o = slice

	return nil
}

// GroupGatewayRightLogExists checks if the GroupGatewayRightLog row exists.
func GroupGatewayRightLogExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `group_gateway_right_log` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if group_gateway_right_log exists")
	}

	return exists, nil
}

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

// CCDataLogCalculatedMonthly is an object representing the database table.
type CCDataLogCalculatedMonthly struct {
	ID                               int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	GWUUID                           string       `boil:"gw_uuid" json:"gwUUID" toml:"gwUUID" yaml:"gwUUID"`
	LatestLogDate                    time.Time    `boil:"latest_log_date" json:"latestLogDate" toml:"latestLogDate" yaml:"latestLogDate"`
	GWID                             null.Int64   `boil:"gw_id" json:"gwID,omitempty" toml:"gwID" yaml:"gwID,omitempty"`
	CustomerID                       null.Int64   `boil:"customer_id" json:"customerID,omitempty" toml:"customerID" yaml:"customerID,omitempty"`
	PvProducedLifetimeEnergyACDiff   null.Float32 `boil:"pv_produced_lifetime_energy_ac_diff" json:"pvProducedLifetimeEnergyAcDiff,omitempty" toml:"pvProducedLifetimeEnergyAcDiff" yaml:"pvProducedLifetimeEnergyAcDiff,omitempty"`
	LoadConsumedLifetimeEnergyACDiff null.Float32 `boil:"load_consumed_lifetime_energy_ac_diff" json:"loadConsumedLifetimeEnergyAcDiff,omitempty" toml:"loadConsumedLifetimeEnergyAcDiff" yaml:"loadConsumedLifetimeEnergyAcDiff,omitempty"`
	BatteryLifetimeEnergyACDiff      null.Float32 `boil:"battery_lifetime_energy_ac_diff" json:"batteryLifetimeEnergyAcDiff,omitempty" toml:"batteryLifetimeEnergyAcDiff" yaml:"batteryLifetimeEnergyAcDiff,omitempty"`
	GridLifetimeEnergyACDiff         null.Float32 `boil:"grid_lifetime_energy_ac_diff" json:"gridLifetimeEnergyAcDiff,omitempty" toml:"gridLifetimeEnergyAcDiff" yaml:"gridLifetimeEnergyAcDiff,omitempty"`
	LoadSelfConsumedEnergyPercentAC  null.Float32 `boil:"load_self_consumed_energy_percent_ac" json:"loadSelfConsumedEnergyPercentAc,omitempty" toml:"loadSelfConsumedEnergyPercentAc" yaml:"loadSelfConsumedEnergyPercentAc,omitempty"`
	CreatedAt                        time.Time    `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt                        time.Time    `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *ccDataLogCalculatedMonthlyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ccDataLogCalculatedMonthlyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CCDataLogCalculatedMonthlyColumns = struct {
	ID                               string
	GWUUID                           string
	LatestLogDate                    string
	GWID                             string
	CustomerID                       string
	PvProducedLifetimeEnergyACDiff   string
	LoadConsumedLifetimeEnergyACDiff string
	BatteryLifetimeEnergyACDiff      string
	GridLifetimeEnergyACDiff         string
	LoadSelfConsumedEnergyPercentAC  string
	CreatedAt                        string
	UpdatedAt                        string
}{
	ID:                               "id",
	GWUUID:                           "gw_uuid",
	LatestLogDate:                    "latest_log_date",
	GWID:                             "gw_id",
	CustomerID:                       "customer_id",
	PvProducedLifetimeEnergyACDiff:   "pv_produced_lifetime_energy_ac_diff",
	LoadConsumedLifetimeEnergyACDiff: "load_consumed_lifetime_energy_ac_diff",
	BatteryLifetimeEnergyACDiff:      "battery_lifetime_energy_ac_diff",
	GridLifetimeEnergyACDiff:         "grid_lifetime_energy_ac_diff",
	LoadSelfConsumedEnergyPercentAC:  "load_self_consumed_energy_percent_ac",
	CreatedAt:                        "created_at",
	UpdatedAt:                        "updated_at",
}

var CCDataLogCalculatedMonthlyTableColumns = struct {
	ID                               string
	GWUUID                           string
	LatestLogDate                    string
	GWID                             string
	CustomerID                       string
	PvProducedLifetimeEnergyACDiff   string
	LoadConsumedLifetimeEnergyACDiff string
	BatteryLifetimeEnergyACDiff      string
	GridLifetimeEnergyACDiff         string
	LoadSelfConsumedEnergyPercentAC  string
	CreatedAt                        string
	UpdatedAt                        string
}{
	ID:                               "cc_data_log_calculated_monthly.id",
	GWUUID:                           "cc_data_log_calculated_monthly.gw_uuid",
	LatestLogDate:                    "cc_data_log_calculated_monthly.latest_log_date",
	GWID:                             "cc_data_log_calculated_monthly.gw_id",
	CustomerID:                       "cc_data_log_calculated_monthly.customer_id",
	PvProducedLifetimeEnergyACDiff:   "cc_data_log_calculated_monthly.pv_produced_lifetime_energy_ac_diff",
	LoadConsumedLifetimeEnergyACDiff: "cc_data_log_calculated_monthly.load_consumed_lifetime_energy_ac_diff",
	BatteryLifetimeEnergyACDiff:      "cc_data_log_calculated_monthly.battery_lifetime_energy_ac_diff",
	GridLifetimeEnergyACDiff:         "cc_data_log_calculated_monthly.grid_lifetime_energy_ac_diff",
	LoadSelfConsumedEnergyPercentAC:  "cc_data_log_calculated_monthly.load_self_consumed_energy_percent_ac",
	CreatedAt:                        "cc_data_log_calculated_monthly.created_at",
	UpdatedAt:                        "cc_data_log_calculated_monthly.updated_at",
}

// Generated where

var CCDataLogCalculatedMonthlyWhere = struct {
	ID                               whereHelperint64
	GWUUID                           whereHelperstring
	LatestLogDate                    whereHelpertime_Time
	GWID                             whereHelpernull_Int64
	CustomerID                       whereHelpernull_Int64
	PvProducedLifetimeEnergyACDiff   whereHelpernull_Float32
	LoadConsumedLifetimeEnergyACDiff whereHelpernull_Float32
	BatteryLifetimeEnergyACDiff      whereHelpernull_Float32
	GridLifetimeEnergyACDiff         whereHelpernull_Float32
	LoadSelfConsumedEnergyPercentAC  whereHelpernull_Float32
	CreatedAt                        whereHelpertime_Time
	UpdatedAt                        whereHelpertime_Time
}{
	ID:                               whereHelperint64{field: "`cc_data_log_calculated_monthly`.`id`"},
	GWUUID:                           whereHelperstring{field: "`cc_data_log_calculated_monthly`.`gw_uuid`"},
	LatestLogDate:                    whereHelpertime_Time{field: "`cc_data_log_calculated_monthly`.`latest_log_date`"},
	GWID:                             whereHelpernull_Int64{field: "`cc_data_log_calculated_monthly`.`gw_id`"},
	CustomerID:                       whereHelpernull_Int64{field: "`cc_data_log_calculated_monthly`.`customer_id`"},
	PvProducedLifetimeEnergyACDiff:   whereHelpernull_Float32{field: "`cc_data_log_calculated_monthly`.`pv_produced_lifetime_energy_ac_diff`"},
	LoadConsumedLifetimeEnergyACDiff: whereHelpernull_Float32{field: "`cc_data_log_calculated_monthly`.`load_consumed_lifetime_energy_ac_diff`"},
	BatteryLifetimeEnergyACDiff:      whereHelpernull_Float32{field: "`cc_data_log_calculated_monthly`.`battery_lifetime_energy_ac_diff`"},
	GridLifetimeEnergyACDiff:         whereHelpernull_Float32{field: "`cc_data_log_calculated_monthly`.`grid_lifetime_energy_ac_diff`"},
	LoadSelfConsumedEnergyPercentAC:  whereHelpernull_Float32{field: "`cc_data_log_calculated_monthly`.`load_self_consumed_energy_percent_ac`"},
	CreatedAt:                        whereHelpertime_Time{field: "`cc_data_log_calculated_monthly`.`created_at`"},
	UpdatedAt:                        whereHelpertime_Time{field: "`cc_data_log_calculated_monthly`.`updated_at`"},
}

// CCDataLogCalculatedMonthlyRels is where relationship names are stored.
var CCDataLogCalculatedMonthlyRels = struct {
}{}

// ccDataLogCalculatedMonthlyR is where relationships are stored.
type ccDataLogCalculatedMonthlyR struct {
}

// NewStruct creates a new relationship struct
func (*ccDataLogCalculatedMonthlyR) NewStruct() *ccDataLogCalculatedMonthlyR {
	return &ccDataLogCalculatedMonthlyR{}
}

// ccDataLogCalculatedMonthlyL is where Load methods for each relationship are stored.
type ccDataLogCalculatedMonthlyL struct{}

var (
	ccDataLogCalculatedMonthlyAllColumns            = []string{"id", "gw_uuid", "latest_log_date", "gw_id", "customer_id", "pv_produced_lifetime_energy_ac_diff", "load_consumed_lifetime_energy_ac_diff", "battery_lifetime_energy_ac_diff", "grid_lifetime_energy_ac_diff", "load_self_consumed_energy_percent_ac", "created_at", "updated_at"}
	ccDataLogCalculatedMonthlyColumnsWithoutDefault = []string{"gw_uuid", "latest_log_date", "gw_id", "customer_id", "pv_produced_lifetime_energy_ac_diff", "load_consumed_lifetime_energy_ac_diff", "battery_lifetime_energy_ac_diff", "grid_lifetime_energy_ac_diff", "load_self_consumed_energy_percent_ac"}
	ccDataLogCalculatedMonthlyColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	ccDataLogCalculatedMonthlyPrimaryKeyColumns     = []string{"id"}
	ccDataLogCalculatedMonthlyGeneratedColumns      = []string{}
)

type (
	// CCDataLogCalculatedMonthlySlice is an alias for a slice of pointers to CCDataLogCalculatedMonthly.
	// This should almost always be used instead of []CCDataLogCalculatedMonthly.
	CCDataLogCalculatedMonthlySlice []*CCDataLogCalculatedMonthly

	ccDataLogCalculatedMonthlyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ccDataLogCalculatedMonthlyType                 = reflect.TypeOf(&CCDataLogCalculatedMonthly{})
	ccDataLogCalculatedMonthlyMapping              = queries.MakeStructMapping(ccDataLogCalculatedMonthlyType)
	ccDataLogCalculatedMonthlyPrimaryKeyMapping, _ = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, ccDataLogCalculatedMonthlyPrimaryKeyColumns)
	ccDataLogCalculatedMonthlyInsertCacheMut       sync.RWMutex
	ccDataLogCalculatedMonthlyInsertCache          = make(map[string]insertCache)
	ccDataLogCalculatedMonthlyUpdateCacheMut       sync.RWMutex
	ccDataLogCalculatedMonthlyUpdateCache          = make(map[string]updateCache)
	ccDataLogCalculatedMonthlyUpsertCacheMut       sync.RWMutex
	ccDataLogCalculatedMonthlyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single ccDataLogCalculatedMonthly record from the query.
func (q ccDataLogCalculatedMonthlyQuery) One(exec boil.Executor) (*CCDataLogCalculatedMonthly, error) {
	o := &CCDataLogCalculatedMonthly{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for cc_data_log_calculated_monthly")
	}

	return o, nil
}

// All returns all CCDataLogCalculatedMonthly records from the query.
func (q ccDataLogCalculatedMonthlyQuery) All(exec boil.Executor) (CCDataLogCalculatedMonthlySlice, error) {
	var o []*CCDataLogCalculatedMonthly

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to CCDataLogCalculatedMonthly slice")
	}

	return o, nil
}

// Count returns the count of all CCDataLogCalculatedMonthly records in the query.
func (q ccDataLogCalculatedMonthlyQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count cc_data_log_calculated_monthly rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ccDataLogCalculatedMonthlyQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if cc_data_log_calculated_monthly exists")
	}

	return count > 0, nil
}

// CCDataLogCalculatedMonthlies retrieves all the records using an executor.
func CCDataLogCalculatedMonthlies(mods ...qm.QueryMod) ccDataLogCalculatedMonthlyQuery {
	mods = append(mods, qm.From("`cc_data_log_calculated_monthly`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`cc_data_log_calculated_monthly`.*"})
	}

	return ccDataLogCalculatedMonthlyQuery{q}
}

// FindCCDataLogCalculatedMonthly retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCCDataLogCalculatedMonthly(exec boil.Executor, iD int64, selectCols ...string) (*CCDataLogCalculatedMonthly, error) {
	ccDataLogCalculatedMonthlyObj := &CCDataLogCalculatedMonthly{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `cc_data_log_calculated_monthly` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, ccDataLogCalculatedMonthlyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from cc_data_log_calculated_monthly")
	}

	return ccDataLogCalculatedMonthlyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CCDataLogCalculatedMonthly) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no cc_data_log_calculated_monthly provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(ccDataLogCalculatedMonthlyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ccDataLogCalculatedMonthlyInsertCacheMut.RLock()
	cache, cached := ccDataLogCalculatedMonthlyInsertCache[key]
	ccDataLogCalculatedMonthlyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ccDataLogCalculatedMonthlyAllColumns,
			ccDataLogCalculatedMonthlyColumnsWithDefault,
			ccDataLogCalculatedMonthlyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `cc_data_log_calculated_monthly` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `cc_data_log_calculated_monthly` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `cc_data_log_calculated_monthly` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, ccDataLogCalculatedMonthlyPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into cc_data_log_calculated_monthly")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ccDataLogCalculatedMonthlyMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for cc_data_log_calculated_monthly")
	}

CacheNoHooks:
	if !cached {
		ccDataLogCalculatedMonthlyInsertCacheMut.Lock()
		ccDataLogCalculatedMonthlyInsertCache[key] = cache
		ccDataLogCalculatedMonthlyInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CCDataLogCalculatedMonthly.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CCDataLogCalculatedMonthly) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	ccDataLogCalculatedMonthlyUpdateCacheMut.RLock()
	cache, cached := ccDataLogCalculatedMonthlyUpdateCache[key]
	ccDataLogCalculatedMonthlyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ccDataLogCalculatedMonthlyAllColumns,
			ccDataLogCalculatedMonthlyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update cc_data_log_calculated_monthly, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `cc_data_log_calculated_monthly` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, ccDataLogCalculatedMonthlyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, append(wl, ccDataLogCalculatedMonthlyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update cc_data_log_calculated_monthly row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for cc_data_log_calculated_monthly")
	}

	if !cached {
		ccDataLogCalculatedMonthlyUpdateCacheMut.Lock()
		ccDataLogCalculatedMonthlyUpdateCache[key] = cache
		ccDataLogCalculatedMonthlyUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q ccDataLogCalculatedMonthlyQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for cc_data_log_calculated_monthly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for cc_data_log_calculated_monthly")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CCDataLogCalculatedMonthlySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ccDataLogCalculatedMonthlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `cc_data_log_calculated_monthly` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ccDataLogCalculatedMonthlyPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in ccDataLogCalculatedMonthly slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all ccDataLogCalculatedMonthly")
	}
	return rowsAff, nil
}

var mySQLCCDataLogCalculatedMonthlyUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CCDataLogCalculatedMonthly) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no cc_data_log_calculated_monthly provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(ccDataLogCalculatedMonthlyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCCDataLogCalculatedMonthlyUniqueColumns, o)

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

	ccDataLogCalculatedMonthlyUpsertCacheMut.RLock()
	cache, cached := ccDataLogCalculatedMonthlyUpsertCache[key]
	ccDataLogCalculatedMonthlyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ccDataLogCalculatedMonthlyAllColumns,
			ccDataLogCalculatedMonthlyColumnsWithDefault,
			ccDataLogCalculatedMonthlyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			ccDataLogCalculatedMonthlyAllColumns,
			ccDataLogCalculatedMonthlyPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert cc_data_log_calculated_monthly, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`cc_data_log_calculated_monthly`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `cc_data_log_calculated_monthly` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for cc_data_log_calculated_monthly")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ccDataLogCalculatedMonthlyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(ccDataLogCalculatedMonthlyType, ccDataLogCalculatedMonthlyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for cc_data_log_calculated_monthly")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for cc_data_log_calculated_monthly")
	}

CacheNoHooks:
	if !cached {
		ccDataLogCalculatedMonthlyUpsertCacheMut.Lock()
		ccDataLogCalculatedMonthlyUpsertCache[key] = cache
		ccDataLogCalculatedMonthlyUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CCDataLogCalculatedMonthly record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CCDataLogCalculatedMonthly) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no CCDataLogCalculatedMonthly provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ccDataLogCalculatedMonthlyPrimaryKeyMapping)
	sql := "DELETE FROM `cc_data_log_calculated_monthly` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from cc_data_log_calculated_monthly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for cc_data_log_calculated_monthly")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ccDataLogCalculatedMonthlyQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no ccDataLogCalculatedMonthlyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from cc_data_log_calculated_monthly")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for cc_data_log_calculated_monthly")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CCDataLogCalculatedMonthlySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ccDataLogCalculatedMonthlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `cc_data_log_calculated_monthly` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ccDataLogCalculatedMonthlyPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from ccDataLogCalculatedMonthly slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for cc_data_log_calculated_monthly")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CCDataLogCalculatedMonthly) Reload(exec boil.Executor) error {
	ret, err := FindCCDataLogCalculatedMonthly(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CCDataLogCalculatedMonthlySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CCDataLogCalculatedMonthlySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ccDataLogCalculatedMonthlyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `cc_data_log_calculated_monthly`.* FROM `cc_data_log_calculated_monthly` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ccDataLogCalculatedMonthlyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in CCDataLogCalculatedMonthlySlice")
	}

	*o = slice

	return nil
}

// CCDataLogCalculatedMonthlyExists checks if the CCDataLogCalculatedMonthly row exists.
func CCDataLogCalculatedMonthlyExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `cc_data_log_calculated_monthly` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if cc_data_log_calculated_monthly exists")
	}

	return exists, nil
}

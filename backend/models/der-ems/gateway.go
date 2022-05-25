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

// Gateway is an object representing the database table.
type Gateway struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UUID           string    `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	CustomerInfoID int       `boil:"customer_info_id" json:"customer_info_id" toml:"customer_info_id" yaml:"customer_info_id"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *gatewayR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gatewayL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GatewayColumns = struct {
	ID             string
	UUID           string
	CustomerInfoID string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	UUID:           "uuid",
	CustomerInfoID: "customer_info_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var GatewayTableColumns = struct {
	ID             string
	UUID           string
	CustomerInfoID string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "gateway.id",
	UUID:           "gateway.uuid",
	CustomerInfoID: "gateway.customer_info_id",
	CreatedAt:      "gateway.created_at",
	UpdatedAt:      "gateway.updated_at",
}

// Generated where

var GatewayWhere = struct {
	ID             whereHelperint
	UUID           whereHelperstring
	CustomerInfoID whereHelperint
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpernull_Time
}{
	ID:             whereHelperint{field: "`gateway`.`id`"},
	UUID:           whereHelperstring{field: "`gateway`.`uuid`"},
	CustomerInfoID: whereHelperint{field: "`gateway`.`customer_info_id`"},
	CreatedAt:      whereHelpertime_Time{field: "`gateway`.`created_at`"},
	UpdatedAt:      whereHelpernull_Time{field: "`gateway`.`updated_at`"},
}

// GatewayRels is where relationship names are stored.
var GatewayRels = struct {
	CustomerInfo string
	GWDevices    string
}{
	CustomerInfo: "CustomerInfo",
	GWDevices:    "GWDevices",
}

// gatewayR is where relationships are stored.
type gatewayR struct {
	CustomerInfo *CustomerInfo `boil:"CustomerInfo" json:"CustomerInfo" toml:"CustomerInfo" yaml:"CustomerInfo"`
	GWDevices    DeviceSlice   `boil:"GWDevices" json:"GWDevices" toml:"GWDevices" yaml:"GWDevices"`
}

// NewStruct creates a new relationship struct
func (*gatewayR) NewStruct() *gatewayR {
	return &gatewayR{}
}

func (r *gatewayR) GetCustomerInfo() *CustomerInfo {
	if r == nil {
		return nil
	}
	return r.CustomerInfo
}

func (r *gatewayR) GetGWDevices() DeviceSlice {
	if r == nil {
		return nil
	}
	return r.GWDevices
}

// gatewayL is where Load methods for each relationship are stored.
type gatewayL struct{}

var (
	gatewayAllColumns            = []string{"id", "uuid", "customer_info_id", "created_at", "updated_at"}
	gatewayColumnsWithoutDefault = []string{"uuid", "customer_info_id", "updated_at"}
	gatewayColumnsWithDefault    = []string{"id", "created_at"}
	gatewayPrimaryKeyColumns     = []string{"id"}
	gatewayGeneratedColumns      = []string{}
)

type (
	// GatewaySlice is an alias for a slice of pointers to Gateway.
	// This should almost always be used instead of []Gateway.
	GatewaySlice []*Gateway

	gatewayQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gatewayType                 = reflect.TypeOf(&Gateway{})
	gatewayMapping              = queries.MakeStructMapping(gatewayType)
	gatewayPrimaryKeyMapping, _ = queries.BindMapping(gatewayType, gatewayMapping, gatewayPrimaryKeyColumns)
	gatewayInsertCacheMut       sync.RWMutex
	gatewayInsertCache          = make(map[string]insertCache)
	gatewayUpdateCacheMut       sync.RWMutex
	gatewayUpdateCache          = make(map[string]updateCache)
	gatewayUpsertCacheMut       sync.RWMutex
	gatewayUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single gateway record from the query.
func (q gatewayQuery) One(exec boil.Executor) (*Gateway, error) {
	o := &Gateway{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for gateway")
	}

	return o, nil
}

// All returns all Gateway records from the query.
func (q gatewayQuery) All(exec boil.Executor) (GatewaySlice, error) {
	var o []*Gateway

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to Gateway slice")
	}

	return o, nil
}

// Count returns the count of all Gateway records in the query.
func (q gatewayQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count gateway rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gatewayQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if gateway exists")
	}

	return count > 0, nil
}

// CustomerInfo pointed to by the foreign key.
func (o *Gateway) CustomerInfo(mods ...qm.QueryMod) customerInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.CustomerInfoID),
	}

	queryMods = append(queryMods, mods...)

	return CustomerInfos(queryMods...)
}

// GWDevices retrieves all the device's Devices with an executor via gw_uuid column.
func (o *Gateway) GWDevices(mods ...qm.QueryMod) deviceQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`device`.`gw_uuid`=?", o.UUID),
	)

	return Devices(queryMods...)
}

// LoadCustomerInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gatewayL) LoadCustomerInfo(e boil.Executor, singular bool, maybeGateway interface{}, mods queries.Applicator) error {
	var slice []*Gateway
	var object *Gateway

	if singular {
		object = maybeGateway.(*Gateway)
	} else {
		slice = *maybeGateway.(*[]*Gateway)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gatewayR{}
		}
		args = append(args, object.CustomerInfoID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gatewayR{}
			}

			for _, a := range args {
				if a == obj.CustomerInfoID {
					continue Outer
				}
			}

			args = append(args, obj.CustomerInfoID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`customer_info`),
		qm.WhereIn(`customer_info.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CustomerInfo")
	}

	var resultSlice []*CustomerInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CustomerInfo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for customer_info")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for customer_info")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.CustomerInfo = foreign
		if foreign.R == nil {
			foreign.R = &customerInfoR{}
		}
		foreign.R.Gateways = append(foreign.R.Gateways, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CustomerInfoID == foreign.ID {
				local.R.CustomerInfo = foreign
				if foreign.R == nil {
					foreign.R = &customerInfoR{}
				}
				foreign.R.Gateways = append(foreign.R.Gateways, local)
				break
			}
		}
	}

	return nil
}

// LoadGWDevices allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (gatewayL) LoadGWDevices(e boil.Executor, singular bool, maybeGateway interface{}, mods queries.Applicator) error {
	var slice []*Gateway
	var object *Gateway

	if singular {
		object = maybeGateway.(*Gateway)
	} else {
		slice = *maybeGateway.(*[]*Gateway)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gatewayR{}
		}
		args = append(args, object.UUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gatewayR{}
			}

			for _, a := range args {
				if a == obj.UUID {
					continue Outer
				}
			}

			args = append(args, obj.UUID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`device`),
		qm.WhereIn(`device.gw_uuid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load device")
	}

	var resultSlice []*Device
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice device")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on device")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for device")
	}

	if singular {
		object.R.GWDevices = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &deviceR{}
			}
			foreign.R.GW = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UUID == foreign.GWUUID {
				local.R.GWDevices = append(local.R.GWDevices, foreign)
				if foreign.R == nil {
					foreign.R = &deviceR{}
				}
				foreign.R.GW = local
				break
			}
		}
	}

	return nil
}

// SetCustomerInfo of the gateway to the related item.
// Sets o.R.CustomerInfo to related.
// Adds o to related.R.Gateways.
func (o *Gateway) SetCustomerInfo(exec boil.Executor, insert bool, related *CustomerInfo) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `gateway` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"customer_info_id"}),
		strmangle.WhereClause("`", "`", 0, gatewayPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CustomerInfoID = related.ID
	if o.R == nil {
		o.R = &gatewayR{
			CustomerInfo: related,
		}
	} else {
		o.R.CustomerInfo = related
	}

	if related.R == nil {
		related.R = &customerInfoR{
			Gateways: GatewaySlice{o},
		}
	} else {
		related.R.Gateways = append(related.R.Gateways, o)
	}

	return nil
}

// AddGWDevices adds the given related objects to the existing relationships
// of the gateway, optionally inserting them as new records.
// Appends related to o.R.GWDevices.
// Sets related.R.GW appropriately.
func (o *Gateway) AddGWDevices(exec boil.Executor, insert bool, related ...*Device) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GWUUID = o.UUID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `device` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"gw_uuid"}),
				strmangle.WhereClause("`", "`", 0, devicePrimaryKeyColumns),
			)
			values := []interface{}{o.UUID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GWUUID = o.UUID
		}
	}

	if o.R == nil {
		o.R = &gatewayR{
			GWDevices: related,
		}
	} else {
		o.R.GWDevices = append(o.R.GWDevices, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &deviceR{
				GW: o,
			}
		} else {
			rel.R.GW = o
		}
	}
	return nil
}

// Gateways retrieves all the records using an executor.
func Gateways(mods ...qm.QueryMod) gatewayQuery {
	mods = append(mods, qm.From("`gateway`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`gateway`.*"})
	}

	return gatewayQuery{q}
}

// FindGateway retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGateway(exec boil.Executor, iD int, selectCols ...string) (*Gateway, error) {
	gatewayObj := &Gateway{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `gateway` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, gatewayObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from gateway")
	}

	return gatewayObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Gateway) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no gateway provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(gatewayColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gatewayInsertCacheMut.RLock()
	cache, cached := gatewayInsertCache[key]
	gatewayInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gatewayAllColumns,
			gatewayColumnsWithDefault,
			gatewayColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gatewayType, gatewayMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gatewayType, gatewayMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `gateway` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `gateway` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `gateway` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, gatewayPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into gateway")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == gatewayMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for gateway")
	}

CacheNoHooks:
	if !cached {
		gatewayInsertCacheMut.Lock()
		gatewayInsertCache[key] = cache
		gatewayInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Gateway.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Gateway) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	key := makeCacheKey(columns, nil)
	gatewayUpdateCacheMut.RLock()
	cache, cached := gatewayUpdateCache[key]
	gatewayUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gatewayAllColumns,
			gatewayPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update gateway, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `gateway` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, gatewayPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gatewayType, gatewayMapping, append(wl, gatewayPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update gateway row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for gateway")
	}

	if !cached {
		gatewayUpdateCacheMut.Lock()
		gatewayUpdateCache[key] = cache
		gatewayUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q gatewayQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for gateway")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for gateway")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GatewaySlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gatewayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `gateway` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gatewayPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in gateway slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all gateway")
	}
	return rowsAff, nil
}

var mySQLGatewayUniqueColumns = []string{
	"id",
	"uuid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Gateway) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no gateway provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	nzDefaults := queries.NonZeroDefaultSet(gatewayColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGatewayUniqueColumns, o)

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

	gatewayUpsertCacheMut.RLock()
	cache, cached := gatewayUpsertCache[key]
	gatewayUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gatewayAllColumns,
			gatewayColumnsWithDefault,
			gatewayColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gatewayAllColumns,
			gatewayPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert gateway, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`gateway`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `gateway` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(gatewayType, gatewayMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gatewayType, gatewayMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for gateway")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == gatewayMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(gatewayType, gatewayMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for gateway")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for gateway")
	}

CacheNoHooks:
	if !cached {
		gatewayUpsertCacheMut.Lock()
		gatewayUpsertCache[key] = cache
		gatewayUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Gateway record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Gateway) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no Gateway provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gatewayPrimaryKeyMapping)
	sql := "DELETE FROM `gateway` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from gateway")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for gateway")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gatewayQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no gatewayQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from gateway")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for gateway")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GatewaySlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gatewayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `gateway` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gatewayPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from gateway slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for gateway")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Gateway) Reload(exec boil.Executor) error {
	ret, err := FindGateway(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GatewaySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GatewaySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gatewayPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `gateway`.* FROM `gateway` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gatewayPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in GatewaySlice")
	}

	*o = slice

	return nil
}

// GatewayExists checks if the Gateway row exists.
func GatewayExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `gateway` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if gateway exists")
	}

	return exists, nil
}

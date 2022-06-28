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

// Customer is an object representing the database table.
type Customer struct {
	ID             int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	CustomerNumber string       `boil:"customer_number" json:"customerNumber" toml:"customerNumber" yaml:"customerNumber"`
	FieldNumber    string       `boil:"field_number" json:"fieldNumber" toml:"fieldNumber" yaml:"fieldNumber"`
	Address        null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Lat            null.Float64 `boil:"lat" json:"lat,omitempty" toml:"lat" yaml:"lat,omitempty"`
	LNG            null.Float64 `boil:"lng" json:"lng,omitempty" toml:"lng" yaml:"lng,omitempty"`
	WeatherLat     null.Float32 `boil:"weather_lat" json:"weatherLat,omitempty" toml:"weatherLat" yaml:"weatherLat,omitempty"`
	WeatherLNG     null.Float32 `boil:"weather_lng" json:"weatherLNG,omitempty" toml:"weatherLNG" yaml:"weatherLNG,omitempty"`
	Timezone       null.String  `boil:"timezone" json:"timezone,omitempty" toml:"timezone" yaml:"timezone,omitempty"`
	PowerCompany   null.String  `boil:"power_company" json:"powerCompany,omitempty" toml:"powerCompany" yaml:"powerCompany,omitempty"`
	VoltageType    null.String  `boil:"voltage_type" json:"voltageType,omitempty" toml:"voltageType" yaml:"voltageType,omitempty"`
	TouType        null.String  `boil:"tou_type" json:"touType,omitempty" toml:"touType" yaml:"touType,omitempty"`
	CreatedAt      time.Time    `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt      null.Time    `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`

	R *customerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomerColumns = struct {
	ID             string
	CustomerNumber string
	FieldNumber    string
	Address        string
	Lat            string
	LNG            string
	WeatherLat     string
	WeatherLNG     string
	Timezone       string
	PowerCompany   string
	VoltageType    string
	TouType        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	CustomerNumber: "customer_number",
	FieldNumber:    "field_number",
	Address:        "address",
	Lat:            "lat",
	LNG:            "lng",
	WeatherLat:     "weather_lat",
	WeatherLNG:     "weather_lng",
	Timezone:       "timezone",
	PowerCompany:   "power_company",
	VoltageType:    "voltage_type",
	TouType:        "tou_type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var CustomerTableColumns = struct {
	ID             string
	CustomerNumber string
	FieldNumber    string
	Address        string
	Lat            string
	LNG            string
	WeatherLat     string
	WeatherLNG     string
	Timezone       string
	PowerCompany   string
	VoltageType    string
	TouType        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "customer.id",
	CustomerNumber: "customer.customer_number",
	FieldNumber:    "customer.field_number",
	Address:        "customer.address",
	Lat:            "customer.lat",
	LNG:            "customer.lng",
	WeatherLat:     "customer.weather_lat",
	WeatherLNG:     "customer.weather_lng",
	Timezone:       "customer.timezone",
	PowerCompany:   "customer.power_company",
	VoltageType:    "customer.voltage_type",
	TouType:        "customer.tou_type",
	CreatedAt:      "customer.created_at",
	UpdatedAt:      "customer.updated_at",
}

// Generated where

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

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Float32 struct{ field string }

func (w whereHelpernull_Float32) EQ(x null.Float32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float32) NEQ(x null.Float32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float32) LT(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float32) LTE(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float32) GT(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float32) GTE(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Float32) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float32) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var CustomerWhere = struct {
	ID             whereHelperint
	CustomerNumber whereHelperstring
	FieldNumber    whereHelperstring
	Address        whereHelpernull_String
	Lat            whereHelpernull_Float64
	LNG            whereHelpernull_Float64
	WeatherLat     whereHelpernull_Float32
	WeatherLNG     whereHelpernull_Float32
	Timezone       whereHelpernull_String
	PowerCompany   whereHelpernull_String
	VoltageType    whereHelpernull_String
	TouType        whereHelpernull_String
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpernull_Time
}{
	ID:             whereHelperint{field: "`customer`.`id`"},
	CustomerNumber: whereHelperstring{field: "`customer`.`customer_number`"},
	FieldNumber:    whereHelperstring{field: "`customer`.`field_number`"},
	Address:        whereHelpernull_String{field: "`customer`.`address`"},
	Lat:            whereHelpernull_Float64{field: "`customer`.`lat`"},
	LNG:            whereHelpernull_Float64{field: "`customer`.`lng`"},
	WeatherLat:     whereHelpernull_Float32{field: "`customer`.`weather_lat`"},
	WeatherLNG:     whereHelpernull_Float32{field: "`customer`.`weather_lng`"},
	Timezone:       whereHelpernull_String{field: "`customer`.`timezone`"},
	PowerCompany:   whereHelpernull_String{field: "`customer`.`power_company`"},
	VoltageType:    whereHelpernull_String{field: "`customer`.`voltage_type`"},
	TouType:        whereHelpernull_String{field: "`customer`.`tou_type`"},
	CreatedAt:      whereHelpertime_Time{field: "`customer`.`created_at`"},
	UpdatedAt:      whereHelpernull_Time{field: "`customer`.`updated_at`"},
}

// CustomerRels is where relationship names are stored.
var CustomerRels = struct {
	Gateways string
}{
	Gateways: "Gateways",
}

// customerR is where relationships are stored.
type customerR struct {
	Gateways GatewaySlice `boil:"Gateways" json:"Gateways" toml:"Gateways" yaml:"Gateways"`
}

// NewStruct creates a new relationship struct
func (*customerR) NewStruct() *customerR {
	return &customerR{}
}

func (r *customerR) GetGateways() GatewaySlice {
	if r == nil {
		return nil
	}
	return r.Gateways
}

// customerL is where Load methods for each relationship are stored.
type customerL struct{}

var (
	customerAllColumns            = []string{"id", "customer_number", "field_number", "address", "lat", "lng", "weather_lat", "weather_lng", "timezone", "power_company", "voltage_type", "tou_type", "created_at", "updated_at"}
	customerColumnsWithoutDefault = []string{"customer_number", "field_number", "address", "lat", "lng", "weather_lat", "weather_lng", "timezone", "power_company", "voltage_type", "tou_type", "updated_at"}
	customerColumnsWithDefault    = []string{"id", "created_at"}
	customerPrimaryKeyColumns     = []string{"id"}
	customerGeneratedColumns      = []string{}
)

type (
	// CustomerSlice is an alias for a slice of pointers to Customer.
	// This should almost always be used instead of []Customer.
	CustomerSlice []*Customer

	customerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customerType                 = reflect.TypeOf(&Customer{})
	customerMapping              = queries.MakeStructMapping(customerType)
	customerPrimaryKeyMapping, _ = queries.BindMapping(customerType, customerMapping, customerPrimaryKeyColumns)
	customerInsertCacheMut       sync.RWMutex
	customerInsertCache          = make(map[string]insertCache)
	customerUpdateCacheMut       sync.RWMutex
	customerUpdateCache          = make(map[string]updateCache)
	customerUpsertCacheMut       sync.RWMutex
	customerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single customer record from the query.
func (q customerQuery) One(exec boil.Executor) (*Customer, error) {
	o := &Customer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for customer")
	}

	return o, nil
}

// All returns all Customer records from the query.
func (q customerQuery) All(exec boil.Executor) (CustomerSlice, error) {
	var o []*Customer

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to Customer slice")
	}

	return o, nil
}

// Count returns the count of all Customer records in the query.
func (q customerQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count customer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q customerQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if customer exists")
	}

	return count > 0, nil
}

// Gateways retrieves all the gateway's Gateways with an executor.
func (o *Customer) Gateways(mods ...qm.QueryMod) gatewayQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`gateway`.`customer_id`=?", o.ID),
	)

	return Gateways(queryMods...)
}

// LoadGateways allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (customerL) LoadGateways(e boil.Executor, singular bool, maybeCustomer interface{}, mods queries.Applicator) error {
	var slice []*Customer
	var object *Customer

	if singular {
		object = maybeCustomer.(*Customer)
	} else {
		slice = *maybeCustomer.(*[]*Customer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &customerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &customerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`gateway`),
		qm.WhereIn(`gateway.customer_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load gateway")
	}

	var resultSlice []*Gateway
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice gateway")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on gateway")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for gateway")
	}

	if singular {
		object.R.Gateways = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &gatewayR{}
			}
			foreign.R.Customer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CustomerID {
				local.R.Gateways = append(local.R.Gateways, foreign)
				if foreign.R == nil {
					foreign.R = &gatewayR{}
				}
				foreign.R.Customer = local
				break
			}
		}
	}

	return nil
}

// AddGateways adds the given related objects to the existing relationships
// of the customer, optionally inserting them as new records.
// Appends related to o.R.Gateways.
// Sets related.R.Customer appropriately.
func (o *Customer) AddGateways(exec boil.Executor, insert bool, related ...*Gateway) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CustomerID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `gateway` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"customer_id"}),
				strmangle.WhereClause("`", "`", 0, gatewayPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CustomerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &customerR{
			Gateways: related,
		}
	} else {
		o.R.Gateways = append(o.R.Gateways, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &gatewayR{
				Customer: o,
			}
		} else {
			rel.R.Customer = o
		}
	}
	return nil
}

// Customers retrieves all the records using an executor.
func Customers(mods ...qm.QueryMod) customerQuery {
	mods = append(mods, qm.From("`customer`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`customer`.*"})
	}

	return customerQuery{q}
}

// FindCustomer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomer(exec boil.Executor, iD int, selectCols ...string) (*Customer, error) {
	customerObj := &Customer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `customer` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, customerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from customer")
	}

	return customerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Customer) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no customer provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customerInsertCacheMut.RLock()
	cache, cached := customerInsertCache[key]
	customerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customerType, customerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `customer` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `customer` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `customer` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, customerPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into customer")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == customerMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for customer")
	}

CacheNoHooks:
	if !cached {
		customerInsertCacheMut.Lock()
		customerInsertCache[key] = cache
		customerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Customer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Customer) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	key := makeCacheKey(columns, nil)
	customerUpdateCacheMut.RLock()
	cache, cached := customerUpdateCache[key]
	customerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update customer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `customer` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, customerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, append(wl, customerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update customer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for customer")
	}

	if !cached {
		customerUpdateCacheMut.Lock()
		customerUpdateCache[key] = cache
		customerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q customerQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for customer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for customer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomerSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `customer` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all customer")
	}
	return rowsAff, nil
}

var mySQLCustomerUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Customer) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no customer provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	nzDefaults := queries.NonZeroDefaultSet(customerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCustomerUniqueColumns, o)

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

	customerUpsertCacheMut.RLock()
	cache, cached := customerUpsertCache[key]
	customerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			customerAllColumns,
			customerColumnsWithDefault,
			customerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			customerAllColumns,
			customerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert customer, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`customer`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `customer` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(customerType, customerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customerType, customerMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for customer")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == customerMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(customerType, customerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for customer")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for customer")
	}

CacheNoHooks:
	if !cached {
		customerUpsertCacheMut.Lock()
		customerUpsertCache[key] = cache
		customerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Customer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Customer) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no Customer provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customerPrimaryKeyMapping)
	sql := "DELETE FROM `customer` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from customer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for customer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q customerQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no customerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from customer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for customer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomerSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `customer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from customer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for customer")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Customer) Reload(exec boil.Executor) error {
	ret, err := FindCustomer(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomerSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `customer`.* FROM `customer` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in CustomerSlice")
	}

	*o = slice

	return nil
}

// CustomerExists checks if the Customer row exists.
func CustomerExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `customer` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if customer exists")
	}

	return exists, nil
}

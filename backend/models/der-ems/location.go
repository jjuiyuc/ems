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

// Location is an object representing the database table.
type Location struct {
	ID             int64        `boil:"id" json:"id" toml:"id" yaml:"id"`
	CustomerNumber string       `boil:"customer_number" json:"customerNumber" toml:"customerNumber" yaml:"customerNumber"`
	FieldNumber    string       `boil:"field_number" json:"fieldNumber" toml:"fieldNumber" yaml:"fieldNumber"`
	Address        null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Lat            null.Float64 `boil:"lat" json:"lat,omitempty" toml:"lat" yaml:"lat,omitempty"`
	Lng            null.Float64 `boil:"lng" json:"lng,omitempty" toml:"lng" yaml:"lng,omitempty"`
	WeatherLat     null.Float32 `boil:"weather_lat" json:"weatherLat,omitempty" toml:"weatherLat" yaml:"weatherLat,omitempty"`
	WeatherLng     null.Float32 `boil:"weather_lng" json:"weatherLNG,omitempty" toml:"weatherLNG" yaml:"weatherLNG,omitempty"`
	Timezone       null.String  `boil:"timezone" json:"timezone,omitempty" toml:"timezone" yaml:"timezone,omitempty"`
	TOULocationID  null.Int64   `boil:"tou_location_id" json:"touLocationID,omitempty" toml:"touLocationID" yaml:"touLocationID,omitempty"`
	VoltageType    null.String  `boil:"voltage_type" json:"voltageType,omitempty" toml:"voltageType" yaml:"voltageType,omitempty"`
	TOUType        null.String  `boil:"tou_type" json:"touType,omitempty" toml:"touType" yaml:"touType,omitempty"`
	CreatedAt      time.Time    `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt      time.Time    `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *locationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L locationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LocationColumns = struct {
	ID             string
	CustomerNumber string
	FieldNumber    string
	Address        string
	Lat            string
	Lng            string
	WeatherLat     string
	WeatherLng     string
	Timezone       string
	TOULocationID  string
	VoltageType    string
	TOUType        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	CustomerNumber: "customer_number",
	FieldNumber:    "field_number",
	Address:        "address",
	Lat:            "lat",
	Lng:            "lng",
	WeatherLat:     "weather_lat",
	WeatherLng:     "weather_lng",
	Timezone:       "timezone",
	TOULocationID:  "tou_location_id",
	VoltageType:    "voltage_type",
	TOUType:        "tou_type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

var LocationTableColumns = struct {
	ID             string
	CustomerNumber string
	FieldNumber    string
	Address        string
	Lat            string
	Lng            string
	WeatherLat     string
	WeatherLng     string
	Timezone       string
	TOULocationID  string
	VoltageType    string
	TOUType        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "location.id",
	CustomerNumber: "location.customer_number",
	FieldNumber:    "location.field_number",
	Address:        "location.address",
	Lat:            "location.lat",
	Lng:            "location.lng",
	WeatherLat:     "location.weather_lat",
	WeatherLng:     "location.weather_lng",
	Timezone:       "location.timezone",
	TOULocationID:  "location.tou_location_id",
	VoltageType:    "location.voltage_type",
	TOUType:        "location.tou_type",
	CreatedAt:      "location.created_at",
	UpdatedAt:      "location.updated_at",
}

// Generated where

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

var LocationWhere = struct {
	ID             whereHelperint64
	CustomerNumber whereHelperstring
	FieldNumber    whereHelperstring
	Address        whereHelpernull_String
	Lat            whereHelpernull_Float64
	Lng            whereHelpernull_Float64
	WeatherLat     whereHelpernull_Float32
	WeatherLng     whereHelpernull_Float32
	Timezone       whereHelpernull_String
	TOULocationID  whereHelpernull_Int64
	VoltageType    whereHelpernull_String
	TOUType        whereHelpernull_String
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpertime_Time
}{
	ID:             whereHelperint64{field: "`location`.`id`"},
	CustomerNumber: whereHelperstring{field: "`location`.`customer_number`"},
	FieldNumber:    whereHelperstring{field: "`location`.`field_number`"},
	Address:        whereHelpernull_String{field: "`location`.`address`"},
	Lat:            whereHelpernull_Float64{field: "`location`.`lat`"},
	Lng:            whereHelpernull_Float64{field: "`location`.`lng`"},
	WeatherLat:     whereHelpernull_Float32{field: "`location`.`weather_lat`"},
	WeatherLng:     whereHelpernull_Float32{field: "`location`.`weather_lng`"},
	Timezone:       whereHelpernull_String{field: "`location`.`timezone`"},
	TOULocationID:  whereHelpernull_Int64{field: "`location`.`tou_location_id`"},
	VoltageType:    whereHelpernull_String{field: "`location`.`voltage_type`"},
	TOUType:        whereHelpernull_String{field: "`location`.`tou_type`"},
	CreatedAt:      whereHelpertime_Time{field: "`location`.`created_at`"},
	UpdatedAt:      whereHelpertime_Time{field: "`location`.`updated_at`"},
}

// LocationRels is where relationship names are stored.
var LocationRels = struct {
	Gateways string
}{
	Gateways: "Gateways",
}

// locationR is where relationships are stored.
type locationR struct {
	Gateways GatewaySlice `boil:"Gateways" json:"Gateways" toml:"Gateways" yaml:"Gateways"`
}

// NewStruct creates a new relationship struct
func (*locationR) NewStruct() *locationR {
	return &locationR{}
}

func (r *locationR) GetGateways() GatewaySlice {
	if r == nil {
		return nil
	}
	return r.Gateways
}

// locationL is where Load methods for each relationship are stored.
type locationL struct{}

var (
	locationAllColumns            = []string{"id", "customer_number", "field_number", "address", "lat", "lng", "weather_lat", "weather_lng", "timezone", "tou_location_id", "voltage_type", "tou_type", "created_at", "updated_at"}
	locationColumnsWithoutDefault = []string{"customer_number", "field_number", "address", "lat", "lng", "weather_lat", "weather_lng", "timezone", "tou_location_id", "voltage_type", "tou_type"}
	locationColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	locationPrimaryKeyColumns     = []string{"id"}
	locationGeneratedColumns      = []string{}
)

type (
	// LocationSlice is an alias for a slice of pointers to Location.
	// This should almost always be used instead of []Location.
	LocationSlice []*Location

	locationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	locationType                 = reflect.TypeOf(&Location{})
	locationMapping              = queries.MakeStructMapping(locationType)
	locationPrimaryKeyMapping, _ = queries.BindMapping(locationType, locationMapping, locationPrimaryKeyColumns)
	locationInsertCacheMut       sync.RWMutex
	locationInsertCache          = make(map[string]insertCache)
	locationUpdateCacheMut       sync.RWMutex
	locationUpdateCache          = make(map[string]updateCache)
	locationUpsertCacheMut       sync.RWMutex
	locationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single location record from the query.
func (q locationQuery) One(exec boil.Executor) (*Location, error) {
	o := &Location{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for location")
	}

	return o, nil
}

// All returns all Location records from the query.
func (q locationQuery) All(exec boil.Executor) (LocationSlice, error) {
	var o []*Location

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to Location slice")
	}

	return o, nil
}

// Count returns the count of all Location records in the query.
func (q locationQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count location rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q locationQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if location exists")
	}

	return count > 0, nil
}

// Gateways retrieves all the gateway's Gateways with an executor.
func (o *Location) Gateways(mods ...qm.QueryMod) gatewayQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`gateway`.`location_id`=?", o.ID),
	)

	return Gateways(queryMods...)
}

// LoadGateways allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (locationL) LoadGateways(e boil.Executor, singular bool, maybeLocation interface{}, mods queries.Applicator) error {
	var slice []*Location
	var object *Location

	if singular {
		object = maybeLocation.(*Location)
	} else {
		slice = *maybeLocation.(*[]*Location)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &locationR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &locationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
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
		qm.WhereIn(`gateway.location_id in ?`, args...),
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
			foreign.R.Location = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.LocationID) {
				local.R.Gateways = append(local.R.Gateways, foreign)
				if foreign.R == nil {
					foreign.R = &gatewayR{}
				}
				foreign.R.Location = local
				break
			}
		}
	}

	return nil
}

// AddGateways adds the given related objects to the existing relationships
// of the location, optionally inserting them as new records.
// Appends related to o.R.Gateways.
// Sets related.R.Location appropriately.
func (o *Location) AddGateways(exec boil.Executor, insert bool, related ...*Gateway) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.LocationID, o.ID)
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `gateway` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"location_id"}),
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

			queries.Assign(&rel.LocationID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &locationR{
			Gateways: related,
		}
	} else {
		o.R.Gateways = append(o.R.Gateways, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &gatewayR{
				Location: o,
			}
		} else {
			rel.R.Location = o
		}
	}
	return nil
}

// SetGateways removes all previously related items of the
// location replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Location's Gateways accordingly.
// Replaces o.R.Gateways with related.
// Sets related.R.Location's Gateways accordingly.
func (o *Location) SetGateways(exec boil.Executor, insert bool, related ...*Gateway) error {
	query := "update `gateway` set `location_id` = null where `location_id` = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Gateways {
			queries.SetScanner(&rel.LocationID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Location = nil
		}
		o.R.Gateways = nil
	}

	return o.AddGateways(exec, insert, related...)
}

// RemoveGateways relationships from objects passed in.
// Removes related items from R.Gateways (uses pointer comparison, removal does not keep order)
// Sets related.R.Location.
func (o *Location) RemoveGateways(exec boil.Executor, related ...*Gateway) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.LocationID, nil)
		if rel.R != nil {
			rel.R.Location = nil
		}
		if _, err = rel.Update(exec, boil.Whitelist("location_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Gateways {
			if rel != ri {
				continue
			}

			ln := len(o.R.Gateways)
			if ln > 1 && i < ln-1 {
				o.R.Gateways[i] = o.R.Gateways[ln-1]
			}
			o.R.Gateways = o.R.Gateways[:ln-1]
			break
		}
	}

	return nil
}

// Locations retrieves all the records using an executor.
func Locations(mods ...qm.QueryMod) locationQuery {
	mods = append(mods, qm.From("`location`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`location`.*"})
	}

	return locationQuery{q}
}

// FindLocation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLocation(exec boil.Executor, iD int64, selectCols ...string) (*Location, error) {
	locationObj := &Location{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `location` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, locationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from location")
	}

	return locationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Location) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no location provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(locationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	locationInsertCacheMut.RLock()
	cache, cached := locationInsertCache[key]
	locationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			locationAllColumns,
			locationColumnsWithDefault,
			locationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(locationType, locationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(locationType, locationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `location` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `location` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `location` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, locationPrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into location")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == locationMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for location")
	}

CacheNoHooks:
	if !cached {
		locationInsertCacheMut.Lock()
		locationInsertCache[key] = cache
		locationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Location.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Location) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	locationUpdateCacheMut.RLock()
	cache, cached := locationUpdateCache[key]
	locationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			locationAllColumns,
			locationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update location, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `location` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, locationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(locationType, locationMapping, append(wl, locationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update location row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for location")
	}

	if !cached {
		locationUpdateCacheMut.Lock()
		locationUpdateCache[key] = cache
		locationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q locationQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for location")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for location")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LocationSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), locationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `location` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, locationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in location slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all location")
	}
	return rowsAff, nil
}

var mySQLLocationUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Location) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no location provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(locationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLocationUniqueColumns, o)

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

	locationUpsertCacheMut.RLock()
	cache, cached := locationUpsertCache[key]
	locationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			locationAllColumns,
			locationColumnsWithDefault,
			locationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			locationAllColumns,
			locationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert location, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`location`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `location` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(locationType, locationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(locationType, locationMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for location")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == locationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(locationType, locationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for location")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for location")
	}

CacheNoHooks:
	if !cached {
		locationUpsertCacheMut.Lock()
		locationUpsertCache[key] = cache
		locationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Location record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Location) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no Location provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), locationPrimaryKeyMapping)
	sql := "DELETE FROM `location` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from location")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for location")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q locationQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no locationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from location")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for location")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LocationSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), locationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `location` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, locationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from location slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for location")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Location) Reload(exec boil.Executor) error {
	ret, err := FindLocation(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LocationSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LocationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), locationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `location`.* FROM `location` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, locationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in LocationSlice")
	}

	*o = slice

	return nil
}

// LocationExists checks if the Location row exists.
func LocationExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `location` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if location exists")
	}

	return exists, nil
}

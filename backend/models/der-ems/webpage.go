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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Webpage is an object representing the database table.
type Webpage struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *webpageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L webpageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WebpageColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var WebpageTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "webpage.id",
	Name:      "webpage.name",
	CreatedAt: "webpage.created_at",
	UpdatedAt: "webpage.updated_at",
}

// Generated where

var WebpageWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`webpage`.`id`"},
	Name:      whereHelperstring{field: "`webpage`.`name`"},
	CreatedAt: whereHelpertime_Time{field: "`webpage`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`webpage`.`updated_at`"},
}

// WebpageRels is where relationship names are stored.
var WebpageRels = struct {
	GroupTypeWebpageRights string
}{
	GroupTypeWebpageRights: "GroupTypeWebpageRights",
}

// webpageR is where relationships are stored.
type webpageR struct {
	GroupTypeWebpageRights GroupTypeWebpageRightSlice `boil:"GroupTypeWebpageRights" json:"GroupTypeWebpageRights" toml:"GroupTypeWebpageRights" yaml:"GroupTypeWebpageRights"`
}

// NewStruct creates a new relationship struct
func (*webpageR) NewStruct() *webpageR {
	return &webpageR{}
}

func (r *webpageR) GetGroupTypeWebpageRights() GroupTypeWebpageRightSlice {
	if r == nil {
		return nil
	}
	return r.GroupTypeWebpageRights
}

// webpageL is where Load methods for each relationship are stored.
type webpageL struct{}

var (
	webpageAllColumns            = []string{"id", "name", "created_at", "updated_at"}
	webpageColumnsWithoutDefault = []string{"name"}
	webpageColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	webpagePrimaryKeyColumns     = []string{"id"}
	webpageGeneratedColumns      = []string{}
)

type (
	// WebpageSlice is an alias for a slice of pointers to Webpage.
	// This should almost always be used instead of []Webpage.
	WebpageSlice []*Webpage

	webpageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	webpageType                 = reflect.TypeOf(&Webpage{})
	webpageMapping              = queries.MakeStructMapping(webpageType)
	webpagePrimaryKeyMapping, _ = queries.BindMapping(webpageType, webpageMapping, webpagePrimaryKeyColumns)
	webpageInsertCacheMut       sync.RWMutex
	webpageInsertCache          = make(map[string]insertCache)
	webpageUpdateCacheMut       sync.RWMutex
	webpageUpdateCache          = make(map[string]updateCache)
	webpageUpsertCacheMut       sync.RWMutex
	webpageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single webpage record from the query.
func (q webpageQuery) One(exec boil.Executor) (*Webpage, error) {
	o := &Webpage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for webpage")
	}

	return o, nil
}

// All returns all Webpage records from the query.
func (q webpageQuery) All(exec boil.Executor) (WebpageSlice, error) {
	var o []*Webpage

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to Webpage slice")
	}

	return o, nil
}

// Count returns the count of all Webpage records in the query.
func (q webpageQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count webpage rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q webpageQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if webpage exists")
	}

	return count > 0, nil
}

// GroupTypeWebpageRights retrieves all the group_type_webpage_right's GroupTypeWebpageRights with an executor.
func (o *Webpage) GroupTypeWebpageRights(mods ...qm.QueryMod) groupTypeWebpageRightQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`group_type_webpage_right`.`webpage_id`=?", o.ID),
	)

	return GroupTypeWebpageRights(queryMods...)
}

// LoadGroupTypeWebpageRights allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (webpageL) LoadGroupTypeWebpageRights(e boil.Executor, singular bool, maybeWebpage interface{}, mods queries.Applicator) error {
	var slice []*Webpage
	var object *Webpage

	if singular {
		object = maybeWebpage.(*Webpage)
	} else {
		slice = *maybeWebpage.(*[]*Webpage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &webpageR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &webpageR{}
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
		qm.From(`group_type_webpage_right`),
		qm.WhereIn(`group_type_webpage_right.webpage_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load group_type_webpage_right")
	}

	var resultSlice []*GroupTypeWebpageRight
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice group_type_webpage_right")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on group_type_webpage_right")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for group_type_webpage_right")
	}

	if singular {
		object.R.GroupTypeWebpageRights = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupTypeWebpageRightR{}
			}
			foreign.R.Webpage = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.WebpageID {
				local.R.GroupTypeWebpageRights = append(local.R.GroupTypeWebpageRights, foreign)
				if foreign.R == nil {
					foreign.R = &groupTypeWebpageRightR{}
				}
				foreign.R.Webpage = local
				break
			}
		}
	}

	return nil
}

// AddGroupTypeWebpageRights adds the given related objects to the existing relationships
// of the webpage, optionally inserting them as new records.
// Appends related to o.R.GroupTypeWebpageRights.
// Sets related.R.Webpage appropriately.
func (o *Webpage) AddGroupTypeWebpageRights(exec boil.Executor, insert bool, related ...*GroupTypeWebpageRight) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.WebpageID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `group_type_webpage_right` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"webpage_id"}),
				strmangle.WhereClause("`", "`", 0, groupTypeWebpageRightPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.WebpageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &webpageR{
			GroupTypeWebpageRights: related,
		}
	} else {
		o.R.GroupTypeWebpageRights = append(o.R.GroupTypeWebpageRights, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupTypeWebpageRightR{
				Webpage: o,
			}
		} else {
			rel.R.Webpage = o
		}
	}
	return nil
}

// Webpages retrieves all the records using an executor.
func Webpages(mods ...qm.QueryMod) webpageQuery {
	mods = append(mods, qm.From("`webpage`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`webpage`.*"})
	}

	return webpageQuery{q}
}

// FindWebpage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWebpage(exec boil.Executor, iD int64, selectCols ...string) (*Webpage, error) {
	webpageObj := &Webpage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `webpage` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, webpageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from webpage")
	}

	return webpageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Webpage) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no webpage provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(webpageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	webpageInsertCacheMut.RLock()
	cache, cached := webpageInsertCache[key]
	webpageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			webpageAllColumns,
			webpageColumnsWithDefault,
			webpageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(webpageType, webpageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(webpageType, webpageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `webpage` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `webpage` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `webpage` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, webpagePrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into webpage")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webpageMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for webpage")
	}

CacheNoHooks:
	if !cached {
		webpageInsertCacheMut.Lock()
		webpageInsertCache[key] = cache
		webpageInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Webpage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Webpage) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	webpageUpdateCacheMut.RLock()
	cache, cached := webpageUpdateCache[key]
	webpageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			webpageAllColumns,
			webpagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update webpage, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `webpage` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, webpagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(webpageType, webpageMapping, append(wl, webpagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update webpage row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for webpage")
	}

	if !cached {
		webpageUpdateCacheMut.Lock()
		webpageUpdateCache[key] = cache
		webpageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q webpageQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for webpage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for webpage")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WebpageSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webpagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `webpage` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webpagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in webpage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all webpage")
	}
	return rowsAff, nil
}

var mySQLWebpageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Webpage) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no webpage provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(webpageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLWebpageUniqueColumns, o)

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

	webpageUpsertCacheMut.RLock()
	cache, cached := webpageUpsertCache[key]
	webpageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			webpageAllColumns,
			webpageColumnsWithDefault,
			webpageColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			webpageAllColumns,
			webpagePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert webpage, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`webpage`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `webpage` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(webpageType, webpageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(webpageType, webpageMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for webpage")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webpageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(webpageType, webpageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for webpage")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for webpage")
	}

CacheNoHooks:
	if !cached {
		webpageUpsertCacheMut.Lock()
		webpageUpsertCache[key] = cache
		webpageUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Webpage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Webpage) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no Webpage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), webpagePrimaryKeyMapping)
	sql := "DELETE FROM `webpage` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from webpage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for webpage")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q webpageQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no webpageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from webpage")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for webpage")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WebpageSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webpagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `webpage` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webpagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from webpage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for webpage")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Webpage) Reload(exec boil.Executor) error {
	ret, err := FindWebpage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WebpageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WebpageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webpagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `webpage`.* FROM `webpage` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webpagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in WebpageSlice")
	}

	*o = slice

	return nil
}

// WebpageExists checks if the Webpage row exists.
func WebpageExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `webpage` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if webpage exists")
	}

	return exists, nil
}

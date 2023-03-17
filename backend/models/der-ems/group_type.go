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

// GroupType is an object representing the database table.
type GroupType struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	R *groupTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupTypeColumns = struct {
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

var GroupTypeTableColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "group_type.id",
	Name:      "group_type.name",
	CreatedAt: "group_type.created_at",
	UpdatedAt: "group_type.updated_at",
}

// Generated where

var GroupTypeWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`group_type`.`id`"},
	Name:      whereHelperstring{field: "`group_type`.`name`"},
	CreatedAt: whereHelpertime_Time{field: "`group_type`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`group_type`.`updated_at`"},
}

// GroupTypeRels is where relationship names are stored.
var GroupTypeRels = struct {
	TypeGroups                 string
	TypeGroupTypeWebpageRights string
}{
	TypeGroups:                 "TypeGroups",
	TypeGroupTypeWebpageRights: "TypeGroupTypeWebpageRights",
}

// groupTypeR is where relationships are stored.
type groupTypeR struct {
	TypeGroups                 GroupSlice                 `boil:"TypeGroups" json:"TypeGroups" toml:"TypeGroups" yaml:"TypeGroups"`
	TypeGroupTypeWebpageRights GroupTypeWebpageRightSlice `boil:"TypeGroupTypeWebpageRights" json:"TypeGroupTypeWebpageRights" toml:"TypeGroupTypeWebpageRights" yaml:"TypeGroupTypeWebpageRights"`
}

// NewStruct creates a new relationship struct
func (*groupTypeR) NewStruct() *groupTypeR {
	return &groupTypeR{}
}

func (r *groupTypeR) GetTypeGroups() GroupSlice {
	if r == nil {
		return nil
	}
	return r.TypeGroups
}

func (r *groupTypeR) GetTypeGroupTypeWebpageRights() GroupTypeWebpageRightSlice {
	if r == nil {
		return nil
	}
	return r.TypeGroupTypeWebpageRights
}

// groupTypeL is where Load methods for each relationship are stored.
type groupTypeL struct{}

var (
	groupTypeAllColumns            = []string{"id", "name", "created_at", "updated_at"}
	groupTypeColumnsWithoutDefault = []string{"name"}
	groupTypeColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	groupTypePrimaryKeyColumns     = []string{"id"}
	groupTypeGeneratedColumns      = []string{}
)

type (
	// GroupTypeSlice is an alias for a slice of pointers to GroupType.
	// This should almost always be used instead of []GroupType.
	GroupTypeSlice []*GroupType

	groupTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupTypeType                 = reflect.TypeOf(&GroupType{})
	groupTypeMapping              = queries.MakeStructMapping(groupTypeType)
	groupTypePrimaryKeyMapping, _ = queries.BindMapping(groupTypeType, groupTypeMapping, groupTypePrimaryKeyColumns)
	groupTypeInsertCacheMut       sync.RWMutex
	groupTypeInsertCache          = make(map[string]insertCache)
	groupTypeUpdateCacheMut       sync.RWMutex
	groupTypeUpdateCache          = make(map[string]updateCache)
	groupTypeUpsertCacheMut       sync.RWMutex
	groupTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single groupType record from the query.
func (q groupTypeQuery) One(exec boil.Executor) (*GroupType, error) {
	o := &GroupType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: failed to execute a one query for group_type")
	}

	return o, nil
}

// All returns all GroupType records from the query.
func (q groupTypeQuery) All(exec boil.Executor) (GroupTypeSlice, error) {
	var o []*GroupType

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "deremsmodels: failed to assign all query results to GroupType slice")
	}

	return o, nil
}

// Count returns the count of all GroupType records in the query.
func (q groupTypeQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to count group_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupTypeQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: failed to check if group_type exists")
	}

	return count > 0, nil
}

// TypeGroups retrieves all the group's Groups with an executor via type_id column.
func (o *GroupType) TypeGroups(mods ...qm.QueryMod) groupQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`group`.`type_id`=?", o.ID),
	)

	return Groups(queryMods...)
}

// TypeGroupTypeWebpageRights retrieves all the group_type_webpage_right's GroupTypeWebpageRights with an executor via type_id column.
func (o *GroupType) TypeGroupTypeWebpageRights(mods ...qm.QueryMod) groupTypeWebpageRightQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`group_type_webpage_right`.`type_id`=?", o.ID),
	)

	return GroupTypeWebpageRights(queryMods...)
}

// LoadTypeGroups allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupTypeL) LoadTypeGroups(e boil.Executor, singular bool, maybeGroupType interface{}, mods queries.Applicator) error {
	var slice []*GroupType
	var object *GroupType

	if singular {
		object = maybeGroupType.(*GroupType)
	} else {
		slice = *maybeGroupType.(*[]*GroupType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupTypeR{}
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
		qm.From(`group`),
		qm.WhereIn(`group.type_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on group")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for group")
	}

	if singular {
		object.R.TypeGroups = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupR{}
			}
			foreign.R.Type = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TypeID {
				local.R.TypeGroups = append(local.R.TypeGroups, foreign)
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.Type = local
				break
			}
		}
	}

	return nil
}

// LoadTypeGroupTypeWebpageRights allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupTypeL) LoadTypeGroupTypeWebpageRights(e boil.Executor, singular bool, maybeGroupType interface{}, mods queries.Applicator) error {
	var slice []*GroupType
	var object *GroupType

	if singular {
		object = maybeGroupType.(*GroupType)
	} else {
		slice = *maybeGroupType.(*[]*GroupType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupTypeR{}
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
		qm.WhereIn(`group_type_webpage_right.type_id in ?`, args...),
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
		object.R.TypeGroupTypeWebpageRights = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &groupTypeWebpageRightR{}
			}
			foreign.R.Type = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TypeID {
				local.R.TypeGroupTypeWebpageRights = append(local.R.TypeGroupTypeWebpageRights, foreign)
				if foreign.R == nil {
					foreign.R = &groupTypeWebpageRightR{}
				}
				foreign.R.Type = local
				break
			}
		}
	}

	return nil
}

// AddTypeGroups adds the given related objects to the existing relationships
// of the group_type, optionally inserting them as new records.
// Appends related to o.R.TypeGroups.
// Sets related.R.Type appropriately.
func (o *GroupType) AddTypeGroups(exec boil.Executor, insert bool, related ...*Group) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TypeID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `group` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"type_id"}),
				strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &groupTypeR{
			TypeGroups: related,
		}
	} else {
		o.R.TypeGroups = append(o.R.TypeGroups, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// AddTypeGroupTypeWebpageRights adds the given related objects to the existing relationships
// of the group_type, optionally inserting them as new records.
// Appends related to o.R.TypeGroupTypeWebpageRights.
// Sets related.R.Type appropriately.
func (o *GroupType) AddTypeGroupTypeWebpageRights(exec boil.Executor, insert bool, related ...*GroupTypeWebpageRight) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TypeID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `group_type_webpage_right` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"type_id"}),
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

			rel.TypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &groupTypeR{
			TypeGroupTypeWebpageRights: related,
		}
	} else {
		o.R.TypeGroupTypeWebpageRights = append(o.R.TypeGroupTypeWebpageRights, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &groupTypeWebpageRightR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// GroupTypes retrieves all the records using an executor.
func GroupTypes(mods ...qm.QueryMod) groupTypeQuery {
	mods = append(mods, qm.From("`group_type`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`group_type`.*"})
	}

	return groupTypeQuery{q}
}

// FindGroupType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroupType(exec boil.Executor, iD int64, selectCols ...string) (*GroupType, error) {
	groupTypeObj := &GroupType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `group_type` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, groupTypeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "deremsmodels: unable to select from group_type")
	}

	return groupTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GroupType) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no group_type provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(groupTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupTypeInsertCacheMut.RLock()
	cache, cached := groupTypeInsertCache[key]
	groupTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupTypeAllColumns,
			groupTypeColumnsWithDefault,
			groupTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupTypeType, groupTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupTypeType, groupTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `group_type` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `group_type` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `group_type` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, groupTypePrimaryKeyColumns))
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
		return errors.Wrap(err, "deremsmodels: unable to insert into group_type")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupTypeMapping["id"] {
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
		return errors.Wrap(err, "deremsmodels: unable to populate default values for group_type")
	}

CacheNoHooks:
	if !cached {
		groupTypeInsertCacheMut.Lock()
		groupTypeInsertCache[key] = cache
		groupTypeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the GroupType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GroupType) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	groupTypeUpdateCacheMut.RLock()
	cache, cached := groupTypeUpdateCache[key]
	groupTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupTypeAllColumns,
			groupTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("deremsmodels: unable to update group_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `group_type` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, groupTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupTypeType, groupTypeMapping, append(wl, groupTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "deremsmodels: unable to update group_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by update for group_type")
	}

	if !cached {
		groupTypeUpdateCacheMut.Lock()
		groupTypeUpdateCache[key] = cache
		groupTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q groupTypeQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all for group_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected for group_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupTypeSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `group_type` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to update all in groupType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to retrieve rows affected all in update all groupType")
	}
	return rowsAff, nil
}

var mySQLGroupTypeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GroupType) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("deremsmodels: no group_type provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(groupTypeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGroupTypeUniqueColumns, o)

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

	groupTypeUpsertCacheMut.RLock()
	cache, cached := groupTypeUpsertCache[key]
	groupTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupTypeAllColumns,
			groupTypeColumnsWithDefault,
			groupTypeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			groupTypeAllColumns,
			groupTypePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("deremsmodels: unable to upsert group_type, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`group_type`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `group_type` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(groupTypeType, groupTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupTypeType, groupTypeMapping, ret)
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
		return errors.Wrap(err, "deremsmodels: unable to upsert for group_type")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupTypeMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(groupTypeType, groupTypeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to retrieve unique values for group_type")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to populate default values for group_type")
	}

CacheNoHooks:
	if !cached {
		groupTypeUpsertCacheMut.Lock()
		groupTypeUpsertCache[key] = cache
		groupTypeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single GroupType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GroupType) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("deremsmodels: no GroupType provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupTypePrimaryKeyMapping)
	sql := "DELETE FROM `group_type` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete from group_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by delete for group_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupTypeQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("deremsmodels: no groupTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from group_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for group_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupTypeSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `group_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: unable to delete all from groupType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "deremsmodels: failed to get rows affected by deleteall for group_type")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GroupType) Reload(exec boil.Executor) error {
	ret, err := FindGroupType(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupTypeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `group_type`.* FROM `group_type` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "deremsmodels: unable to reload all in GroupTypeSlice")
	}

	*o = slice

	return nil
}

// GroupTypeExists checks if the GroupType row exists.
func GroupTypeExists(exec boil.Executor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `group_type` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "deremsmodels: unable to check if group_type exists")
	}

	return exists, nil
}

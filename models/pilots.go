// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Pilot is an object representing the database table.
type Pilot struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *pilotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pilotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PilotColumns = struct {
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

// Generated where

var PilotWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"pilots\".\"id\""},
	Name:      whereHelperstring{field: "\"pilots\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"pilots\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"pilots\".\"updated_at\""},
}

// PilotRels is where relationship names are stored.
var PilotRels = struct {
	Jets           string
	PilotLanguages string
}{
	Jets:           "Jets",
	PilotLanguages: "PilotLanguages",
}

// pilotR is where relationships are stored.
type pilotR struct {
	Jets           JetSlice
	PilotLanguages PilotLanguageSlice
}

// NewStruct creates a new relationship struct
func (*pilotR) NewStruct() *pilotR {
	return &pilotR{}
}

// pilotL is where Load methods for each relationship are stored.
type pilotL struct{}

var (
	pilotAllColumns            = []string{"id", "name", "created_at", "updated_at"}
	pilotColumnsWithoutDefault = []string{"name", "created_at", "updated_at"}
	pilotColumnsWithDefault    = []string{"id"}
	pilotPrimaryKeyColumns     = []string{"id"}
)

type (
	// PilotSlice is an alias for a slice of pointers to Pilot.
	// This should generally be used opposed to []Pilot.
	PilotSlice []*Pilot

	pilotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pilotType                 = reflect.TypeOf(&Pilot{})
	pilotMapping              = queries.MakeStructMapping(pilotType)
	pilotPrimaryKeyMapping, _ = queries.BindMapping(pilotType, pilotMapping, pilotPrimaryKeyColumns)
	pilotInsertCacheMut       sync.RWMutex
	pilotInsertCache          = make(map[string]insertCache)
	pilotUpdateCacheMut       sync.RWMutex
	pilotUpdateCache          = make(map[string]updateCache)
	pilotUpsertCacheMut       sync.RWMutex
	pilotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneP returns a single pilot record from the query, and panics on error.
func (q pilotQuery) OneP(ctx context.Context, exec boil.ContextExecutor) *Pilot {
	o, err := q.One(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single pilot record from the query.
func (q pilotQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pilot, error) {
	o := &Pilot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pilots")
	}

	return o, nil
}

// AllP returns all Pilot records from the query, and panics on error.
func (q pilotQuery) AllP(ctx context.Context, exec boil.ContextExecutor) PilotSlice {
	o, err := q.All(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Pilot records from the query.
func (q pilotQuery) All(ctx context.Context, exec boil.ContextExecutor) (PilotSlice, error) {
	var o []*Pilot

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pilot slice")
	}

	return o, nil
}

// CountP returns the count of all Pilot records in the query, and panics on error.
func (q pilotQuery) CountP(ctx context.Context, exec boil.ContextExecutor) int64 {
	c, err := q.Count(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Pilot records in the query.
func (q pilotQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pilots rows")
	}

	return count, nil
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q pilotQuery) ExistsP(ctx context.Context, exec boil.ContextExecutor) bool {
	e, err := q.Exists(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q pilotQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pilots exists")
	}

	return count > 0, nil
}

// Jets retrieves all the jet's Jets with an executor.
func (o *Pilot) Jets(mods ...qm.QueryMod) jetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"jets\".\"pilot_id\"=?", o.ID),
	)

	query := Jets(queryMods...)
	queries.SetFrom(query.Query, "\"jets\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"jets\".*"})
	}

	return query
}

// PilotLanguages retrieves all the pilot_language's PilotLanguages with an executor.
func (o *Pilot) PilotLanguages(mods ...qm.QueryMod) pilotLanguageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"pilot_languages\".\"pilot_id\"=?", o.ID),
	)

	query := PilotLanguages(queryMods...)
	queries.SetFrom(query.Query, "\"pilot_languages\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"pilot_languages\".*"})
	}

	return query
}

// LoadJets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadJets(ctx context.Context, e boil.ContextExecutor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
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

	query := NewQuery(qm.From(`jets`), qm.WhereIn(`jets.pilot_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load jets")
	}

	var resultSlice []*Jet
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice jets")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on jets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for jets")
	}

	if singular {
		object.R.Jets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &jetR{}
			}
			foreign.R.Pilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PilotID {
				local.R.Jets = append(local.R.Jets, foreign)
				if foreign.R == nil {
					foreign.R = &jetR{}
				}
				foreign.R.Pilot = local
				break
			}
		}
	}

	return nil
}

// LoadPilotLanguages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadPilotLanguages(ctx context.Context, e boil.ContextExecutor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
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

	query := NewQuery(qm.From(`pilot_languages`), qm.WhereIn(`pilot_languages.pilot_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load pilot_languages")
	}

	var resultSlice []*PilotLanguage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice pilot_languages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on pilot_languages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for pilot_languages")
	}

	if singular {
		object.R.PilotLanguages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &pilotLanguageR{}
			}
			foreign.R.Pilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PilotID {
				local.R.PilotLanguages = append(local.R.PilotLanguages, foreign)
				if foreign.R == nil {
					foreign.R = &pilotLanguageR{}
				}
				foreign.R.Pilot = local
				break
			}
		}
	}

	return nil
}

// AddJetsP adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Jets.
// Sets related.R.Pilot appropriately.
// Panics on error.
func (o *Pilot) AddJetsP(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Jet) {
	if err := o.AddJets(ctx, exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddJets adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Jets.
// Sets related.R.Pilot appropriately.
func (o *Pilot) AddJets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Jet) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PilotID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"jets\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pilot_id"}),
				strmangle.WhereClause("\"", "\"", 2, jetPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PilotID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pilotR{
			Jets: related,
		}
	} else {
		o.R.Jets = append(o.R.Jets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jetR{
				Pilot: o,
			}
		} else {
			rel.R.Pilot = o
		}
	}
	return nil
}

// AddPilotLanguagesP adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.PilotLanguages.
// Sets related.R.Pilot appropriately.
// Panics on error.
func (o *Pilot) AddPilotLanguagesP(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PilotLanguage) {
	if err := o.AddPilotLanguages(ctx, exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddPilotLanguages adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.PilotLanguages.
// Sets related.R.Pilot appropriately.
func (o *Pilot) AddPilotLanguages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PilotLanguage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PilotID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"pilot_languages\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pilot_id"}),
				strmangle.WhereClause("\"", "\"", 2, pilotLanguagePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.PilotID, rel.LanguageID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PilotID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pilotR{
			PilotLanguages: related,
		}
	} else {
		o.R.PilotLanguages = append(o.R.PilotLanguages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pilotLanguageR{
				Pilot: o,
			}
		} else {
			rel.R.Pilot = o
		}
	}
	return nil
}

// Pilots retrieves all the records using an executor.
func Pilots(mods ...qm.QueryMod) pilotQuery {
	mods = append(mods, qm.From("\"pilots\""))
	return pilotQuery{NewQuery(mods...)}
}

// FindPilotP retrieves a single record by ID with an executor, and panics on error.
func FindPilotP(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) *Pilot {
	retobj, err := FindPilot(ctx, exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPilot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPilot(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Pilot, error) {
	pilotObj := &Pilot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pilots\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pilotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pilots")
	}

	return pilotObj, nil
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Pilot) InsertP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) {
	if err := o.Insert(ctx, exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pilot) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pilotInsertCacheMut.RLock()
	cache, cached := pilotInsertCache[key]
	pilotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pilots\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pilots\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pilots")
	}

	if !cached {
		pilotInsertCacheMut.Lock()
		pilotInsertCache[key] = cache
		pilotInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateP uses an executor to update the Pilot, and panics on error.
// See Update for more documentation.
func (o *Pilot) UpdateP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Pilot.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pilot) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	pilotUpdateCacheMut.RLock()
	cache, cached := pilotUpdateCache[key]
	pilotUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pilots, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pilots\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pilotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, append(wl, pilotPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update pilots row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pilots")
	}

	if !cached {
		pilotUpdateCacheMut.Lock()
		pilotUpdateCache[key] = cache
		pilotUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q pilotQuery) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := q.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values.
func (q pilotQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pilots")
	}

	return rowsAff, nil
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PilotSlice) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PilotSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pilots\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pilotPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pilot")
	}
	return rowsAff, nil
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Pilot) UpsertP(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pilot) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	pilotUpsertCacheMut.RLock()
	cache, cached := pilotUpsertCache[key]
	pilotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pilots, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pilotPrimaryKeyColumns))
			copy(conflict, pilotPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pilots\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert pilots")
	}

	if !cached {
		pilotUpsertCacheMut.Lock()
		pilotUpsertCache[key] = cache
		pilotUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Pilot record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pilot) DeleteP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.Delete(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Pilot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pilot) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pilot provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pilotPrimaryKeyMapping)
	sql := "DELETE FROM \"pilots\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pilots")
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q pilotQuery) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := q.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q pilotQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pilotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PilotSlice) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PilotSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pilots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	return rowsAff, nil
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Pilot) ReloadP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.Reload(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pilot) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPilot(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PilotSlice) ReloadAllP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.ReloadAll(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PilotSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PilotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pilots\".* FROM \"pilots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PilotSlice")
	}

	*o = slice

	return nil
}

// PilotExistsP checks if the Pilot row exists. Panics on error.
func PilotExistsP(ctx context.Context, exec boil.ContextExecutor, iD string) bool {
	e, err := PilotExists(ctx, exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PilotExists checks if the Pilot row exists.
func PilotExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pilots\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pilots exists")
	}

	return exists, nil
}

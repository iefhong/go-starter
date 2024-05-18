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

// Language is an object representing the database table.
type Language struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Language  string    `boil:"language" json:"language" toml:"language" yaml:"language"`

	R *languageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L languageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LanguageColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	Language  string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Language:  "language",
}

// Generated where

var LanguageWhere = struct {
	ID        whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
	Language  whereHelperstring
}{
	ID:        whereHelperstring{field: "\"languages\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"languages\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"languages\".\"updated_at\""},
	Language:  whereHelperstring{field: "\"languages\".\"language\""},
}

// LanguageRels is where relationship names are stored.
var LanguageRels = struct {
	PilotLanguages string
}{
	PilotLanguages: "PilotLanguages",
}

// languageR is where relationships are stored.
type languageR struct {
	PilotLanguages PilotLanguageSlice
}

// NewStruct creates a new relationship struct
func (*languageR) NewStruct() *languageR {
	return &languageR{}
}

// languageL is where Load methods for each relationship are stored.
type languageL struct{}

var (
	languageAllColumns            = []string{"id", "created_at", "updated_at", "language"}
	languageColumnsWithoutDefault = []string{"created_at", "updated_at", "language"}
	languageColumnsWithDefault    = []string{"id"}
	languagePrimaryKeyColumns     = []string{"id"}
)

type (
	// LanguageSlice is an alias for a slice of pointers to Language.
	// This should generally be used opposed to []Language.
	LanguageSlice []*Language

	languageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	languageType                 = reflect.TypeOf(&Language{})
	languageMapping              = queries.MakeStructMapping(languageType)
	languagePrimaryKeyMapping, _ = queries.BindMapping(languageType, languageMapping, languagePrimaryKeyColumns)
	languageInsertCacheMut       sync.RWMutex
	languageInsertCache          = make(map[string]insertCache)
	languageUpdateCacheMut       sync.RWMutex
	languageUpdateCache          = make(map[string]updateCache)
	languageUpsertCacheMut       sync.RWMutex
	languageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneP returns a single language record from the query, and panics on error.
func (q languageQuery) OneP(ctx context.Context, exec boil.ContextExecutor) *Language {
	o, err := q.One(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single language record from the query.
func (q languageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Language, error) {
	o := &Language{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for languages")
	}

	return o, nil
}

// AllP returns all Language records from the query, and panics on error.
func (q languageQuery) AllP(ctx context.Context, exec boil.ContextExecutor) LanguageSlice {
	o, err := q.All(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Language records from the query.
func (q languageQuery) All(ctx context.Context, exec boil.ContextExecutor) (LanguageSlice, error) {
	var o []*Language

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Language slice")
	}

	return o, nil
}

// CountP returns the count of all Language records in the query, and panics on error.
func (q languageQuery) CountP(ctx context.Context, exec boil.ContextExecutor) int64 {
	c, err := q.Count(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Language records in the query.
func (q languageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count languages rows")
	}

	return count, nil
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q languageQuery) ExistsP(ctx context.Context, exec boil.ContextExecutor) bool {
	e, err := q.Exists(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q languageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if languages exists")
	}

	return count > 0, nil
}

// PilotLanguages retrieves all the pilot_language's PilotLanguages with an executor.
func (o *Language) PilotLanguages(mods ...qm.QueryMod) pilotLanguageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"pilot_languages\".\"language_id\"=?", o.ID),
	)

	query := PilotLanguages(queryMods...)
	queries.SetFrom(query.Query, "\"pilot_languages\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"pilot_languages\".*"})
	}

	return query
}

// LoadPilotLanguages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (languageL) LoadPilotLanguages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLanguage interface{}, mods queries.Applicator) error {
	var slice []*Language
	var object *Language

	if singular {
		object = maybeLanguage.(*Language)
	} else {
		slice = *maybeLanguage.(*[]*Language)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &languageR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &languageR{}
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

	query := NewQuery(qm.From(`pilot_languages`), qm.WhereIn(`pilot_languages.language_id in ?`, args...))
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
			foreign.R.Language = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.LanguageID {
				local.R.PilotLanguages = append(local.R.PilotLanguages, foreign)
				if foreign.R == nil {
					foreign.R = &pilotLanguageR{}
				}
				foreign.R.Language = local
				break
			}
		}
	}

	return nil
}

// AddPilotLanguagesP adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.PilotLanguages.
// Sets related.R.Language appropriately.
// Panics on error.
func (o *Language) AddPilotLanguagesP(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PilotLanguage) {
	if err := o.AddPilotLanguages(ctx, exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddPilotLanguages adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.PilotLanguages.
// Sets related.R.Language appropriately.
func (o *Language) AddPilotLanguages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PilotLanguage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.LanguageID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"pilot_languages\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"language_id"}),
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

			rel.LanguageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &languageR{
			PilotLanguages: related,
		}
	} else {
		o.R.PilotLanguages = append(o.R.PilotLanguages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pilotLanguageR{
				Language: o,
			}
		} else {
			rel.R.Language = o
		}
	}
	return nil
}

// Languages retrieves all the records using an executor.
func Languages(mods ...qm.QueryMod) languageQuery {
	mods = append(mods, qm.From("\"languages\""))
	return languageQuery{NewQuery(mods...)}
}

// FindLanguageP retrieves a single record by ID with an executor, and panics on error.
func FindLanguageP(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) *Language {
	retobj, err := FindLanguage(ctx, exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLanguage(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Language, error) {
	languageObj := &Language{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"languages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, languageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from languages")
	}

	return languageObj, nil
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Language) InsertP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) {
	if err := o.Insert(ctx, exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Language) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no languages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	languageInsertCacheMut.RLock()
	cache, cached := languageInsertCache[key]
	languageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			languageAllColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(languageType, languageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"languages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"languages\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into languages")
	}

	if !cached {
		languageInsertCacheMut.Lock()
		languageInsertCache[key] = cache
		languageInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateP uses an executor to update the Language, and panics on error.
// See Update for more documentation.
func (o *Language) UpdateP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Language.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Language) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	languageUpdateCacheMut.RLock()
	cache, cached := languageUpdateCache[key]
	languageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			languageAllColumns,
			languagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update languages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"languages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, languagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, append(wl, languagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update languages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for languages")
	}

	if !cached {
		languageUpdateCacheMut.Lock()
		languageUpdateCache[key] = cache
		languageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q languageQuery) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := q.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values.
func (q languageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for languages")
	}

	return rowsAff, nil
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LanguageSlice) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LanguageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"languages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, languagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in language slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all language")
	}
	return rowsAff, nil
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Language) UpsertP(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Language) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no languages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

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

	languageUpsertCacheMut.RLock()
	cache, cached := languageUpsertCache[key]
	languageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			languageAllColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			languageAllColumns,
			languagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert languages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(languagePrimaryKeyColumns))
			copy(conflict, languagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"languages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(languageType, languageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert languages")
	}

	if !cached {
		languageUpsertCacheMut.Lock()
		languageUpsertCache[key] = cache
		languageUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Language record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Language) DeleteP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.Delete(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Language record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Language) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Language provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), languagePrimaryKeyMapping)
	sql := "DELETE FROM \"languages\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for languages")
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q languageQuery) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := q.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q languageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no languageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for languages")
	}

	return rowsAff, nil
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LanguageSlice) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LanguageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, languagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from language slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for languages")
	}

	return rowsAff, nil
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Language) ReloadP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.Reload(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Language) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLanguage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LanguageSlice) ReloadAllP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.ReloadAll(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LanguageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"languages\".* FROM \"languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, languagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LanguageSlice")
	}

	*o = slice

	return nil
}

// LanguageExistsP checks if the Language row exists. Panics on error.
func LanguageExistsP(ctx context.Context, exec boil.ContextExecutor, iD string) bool {
	e, err := LanguageExists(ctx, exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LanguageExists checks if the Language row exists.
func LanguageExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"languages\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if languages exists")
	}

	return exists, nil
}

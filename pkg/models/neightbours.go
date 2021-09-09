// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Neightbour is an object representing the database table.
type Neightbour struct {
	PeerID           string    `boil:"peer_id" json:"peer_id" toml:"peer_id" yaml:"peer_id"`
	NeightbourPeerID string    `boil:"neightbour_peer_id" json:"neightbour_peer_id" toml:"neightbour_peer_id" yaml:"neightbour_peer_id"`
	CreatedAt        null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *neightbourR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L neightbourL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NeightbourColumns = struct {
	PeerID           string
	NeightbourPeerID string
	CreatedAt        string
}{
	PeerID:           "peer_id",
	NeightbourPeerID: "neightbour_peer_id",
	CreatedAt:        "created_at",
}

var NeightbourTableColumns = struct {
	PeerID           string
	NeightbourPeerID string
	CreatedAt        string
}{
	PeerID:           "neightbours.peer_id",
	NeightbourPeerID: "neightbours.neightbour_peer_id",
	CreatedAt:        "neightbours.created_at",
}

// Generated where

var NeightbourWhere = struct {
	PeerID           whereHelperstring
	NeightbourPeerID whereHelperstring
	CreatedAt        whereHelpernull_Time
}{
	PeerID:           whereHelperstring{field: "\"neightbours\".\"peer_id\""},
	NeightbourPeerID: whereHelperstring{field: "\"neightbours\".\"neightbour_peer_id\""},
	CreatedAt:        whereHelpernull_Time{field: "\"neightbours\".\"created_at\""},
}

// NeightbourRels is where relationship names are stored.
var NeightbourRels = struct {
}{}

// neightbourR is where relationships are stored.
type neightbourR struct {
}

// NewStruct creates a new relationship struct
func (*neightbourR) NewStruct() *neightbourR {
	return &neightbourR{}
}

// neightbourL is where Load methods for each relationship are stored.
type neightbourL struct{}

var (
	neightbourAllColumns            = []string{"peer_id", "neightbour_peer_id", "created_at"}
	neightbourColumnsWithoutDefault = []string{"peer_id", "neightbour_peer_id", "created_at"}
	neightbourColumnsWithDefault    = []string{}
	neightbourPrimaryKeyColumns     = []string{"peer_id", "neightbour_peer_id"}
)

type (
	// NeightbourSlice is an alias for a slice of pointers to Neightbour.
	// This should almost always be used instead of []Neightbour.
	NeightbourSlice []*Neightbour
	// NeightbourHook is the signature for custom Neightbour hook methods
	NeightbourHook func(context.Context, boil.ContextExecutor, *Neightbour) error

	neightbourQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	neightbourType                 = reflect.TypeOf(&Neightbour{})
	neightbourMapping              = queries.MakeStructMapping(neightbourType)
	neightbourPrimaryKeyMapping, _ = queries.BindMapping(neightbourType, neightbourMapping, neightbourPrimaryKeyColumns)
	neightbourInsertCacheMut       sync.RWMutex
	neightbourInsertCache          = make(map[string]insertCache)
	neightbourUpdateCacheMut       sync.RWMutex
	neightbourUpdateCache          = make(map[string]updateCache)
	neightbourUpsertCacheMut       sync.RWMutex
	neightbourUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var neightbourBeforeInsertHooks []NeightbourHook
var neightbourBeforeUpdateHooks []NeightbourHook
var neightbourBeforeDeleteHooks []NeightbourHook
var neightbourBeforeUpsertHooks []NeightbourHook

var neightbourAfterInsertHooks []NeightbourHook
var neightbourAfterSelectHooks []NeightbourHook
var neightbourAfterUpdateHooks []NeightbourHook
var neightbourAfterDeleteHooks []NeightbourHook
var neightbourAfterUpsertHooks []NeightbourHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Neightbour) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Neightbour) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Neightbour) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Neightbour) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Neightbour) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Neightbour) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Neightbour) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Neightbour) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Neightbour) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range neightbourAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNeightbourHook registers your hook function for all future operations.
func AddNeightbourHook(hookPoint boil.HookPoint, neightbourHook NeightbourHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		neightbourBeforeInsertHooks = append(neightbourBeforeInsertHooks, neightbourHook)
	case boil.BeforeUpdateHook:
		neightbourBeforeUpdateHooks = append(neightbourBeforeUpdateHooks, neightbourHook)
	case boil.BeforeDeleteHook:
		neightbourBeforeDeleteHooks = append(neightbourBeforeDeleteHooks, neightbourHook)
	case boil.BeforeUpsertHook:
		neightbourBeforeUpsertHooks = append(neightbourBeforeUpsertHooks, neightbourHook)
	case boil.AfterInsertHook:
		neightbourAfterInsertHooks = append(neightbourAfterInsertHooks, neightbourHook)
	case boil.AfterSelectHook:
		neightbourAfterSelectHooks = append(neightbourAfterSelectHooks, neightbourHook)
	case boil.AfterUpdateHook:
		neightbourAfterUpdateHooks = append(neightbourAfterUpdateHooks, neightbourHook)
	case boil.AfterDeleteHook:
		neightbourAfterDeleteHooks = append(neightbourAfterDeleteHooks, neightbourHook)
	case boil.AfterUpsertHook:
		neightbourAfterUpsertHooks = append(neightbourAfterUpsertHooks, neightbourHook)
	}
}

// One returns a single neightbour record from the query.
func (q neightbourQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Neightbour, error) {
	o := &Neightbour{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for neightbours")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Neightbour records from the query.
func (q neightbourQuery) All(ctx context.Context, exec boil.ContextExecutor) (NeightbourSlice, error) {
	var o []*Neightbour

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Neightbour slice")
	}

	if len(neightbourAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Neightbour records in the query.
func (q neightbourQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count neightbours rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q neightbourQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if neightbours exists")
	}

	return count > 0, nil
}

// Neightbours retrieves all the records using an executor.
func Neightbours(mods ...qm.QueryMod) neightbourQuery {
	mods = append(mods, qm.From("\"neightbours\""))
	return neightbourQuery{NewQuery(mods...)}
}

// FindNeightbour retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNeightbour(ctx context.Context, exec boil.ContextExecutor, peerID string, neightbourPeerID string, selectCols ...string) (*Neightbour, error) {
	neightbourObj := &Neightbour{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"neightbours\" where \"peer_id\"=$1 AND \"neightbour_peer_id\"=$2", sel,
	)

	q := queries.Raw(query, peerID, neightbourPeerID)

	err := q.Bind(ctx, exec, neightbourObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from neightbours")
	}

	if err = neightbourObj.doAfterSelectHooks(ctx, exec); err != nil {
		return neightbourObj, err
	}

	return neightbourObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Neightbour) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no neightbours provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(neightbourColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	neightbourInsertCacheMut.RLock()
	cache, cached := neightbourInsertCache[key]
	neightbourInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			neightbourAllColumns,
			neightbourColumnsWithDefault,
			neightbourColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(neightbourType, neightbourMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(neightbourType, neightbourMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"neightbours\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"neightbours\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into neightbours")
	}

	if !cached {
		neightbourInsertCacheMut.Lock()
		neightbourInsertCache[key] = cache
		neightbourInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Neightbour.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Neightbour) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	neightbourUpdateCacheMut.RLock()
	cache, cached := neightbourUpdateCache[key]
	neightbourUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			neightbourAllColumns,
			neightbourPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update neightbours, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"neightbours\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, neightbourPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(neightbourType, neightbourMapping, append(wl, neightbourPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update neightbours row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for neightbours")
	}

	if !cached {
		neightbourUpdateCacheMut.Lock()
		neightbourUpdateCache[key] = cache
		neightbourUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q neightbourQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for neightbours")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for neightbours")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NeightbourSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), neightbourPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"neightbours\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, neightbourPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in neightbour slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all neightbour")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Neightbour) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no neightbours provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(neightbourColumnsWithDefault, o)

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

	neightbourUpsertCacheMut.RLock()
	cache, cached := neightbourUpsertCache[key]
	neightbourUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			neightbourAllColumns,
			neightbourColumnsWithDefault,
			neightbourColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			neightbourAllColumns,
			neightbourPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert neightbours, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(neightbourPrimaryKeyColumns))
			copy(conflict, neightbourPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"neightbours\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(neightbourType, neightbourMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(neightbourType, neightbourMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert neightbours")
	}

	if !cached {
		neightbourUpsertCacheMut.Lock()
		neightbourUpsertCache[key] = cache
		neightbourUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Neightbour record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Neightbour) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Neightbour provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), neightbourPrimaryKeyMapping)
	sql := "DELETE FROM \"neightbours\" WHERE \"peer_id\"=$1 AND \"neightbour_peer_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from neightbours")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for neightbours")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q neightbourQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no neightbourQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from neightbours")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for neightbours")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NeightbourSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(neightbourBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), neightbourPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"neightbours\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, neightbourPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from neightbour slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for neightbours")
	}

	if len(neightbourAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Neightbour) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNeightbour(ctx, exec, o.PeerID, o.NeightbourPeerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NeightbourSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NeightbourSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), neightbourPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"neightbours\".* FROM \"neightbours\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, neightbourPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NeightbourSlice")
	}

	*o = slice

	return nil
}

// NeightbourExists checks if the Neightbour row exists.
func NeightbourExists(ctx context.Context, exec boil.ContextExecutor, peerID string, neightbourPeerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"neightbours\" where \"peer_id\"=$1 AND \"neightbour_peer_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, peerID, neightbourPeerID)
	}
	row := exec.QueryRowContext(ctx, sql, peerID, neightbourPeerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if neightbours exists")
	}

	return exists, nil
}

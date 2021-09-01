// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMonitoringPeers(t *testing.T) {
	t.Parallel()

	query := MonitoringPeers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMonitoringPeersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringPeersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MonitoringPeers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringPeersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MonitoringPeerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMonitoringPeersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MonitoringPeerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MonitoringPeer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MonitoringPeerExists to return true, but got false.")
	}
}

func testMonitoringPeersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	monitoringPeerFound, err := FindMonitoringPeer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if monitoringPeerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMonitoringPeersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MonitoringPeers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMonitoringPeersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MonitoringPeers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMonitoringPeersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	monitoringPeerOne := &MonitoringPeer{}
	monitoringPeerTwo := &MonitoringPeer{}
	if err = randomize.Struct(seed, monitoringPeerOne, monitoringPeerDBTypes, false, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, monitoringPeerTwo, monitoringPeerDBTypes, false, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = monitoringPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = monitoringPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MonitoringPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMonitoringPeersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	monitoringPeerOne := &MonitoringPeer{}
	monitoringPeerTwo := &MonitoringPeer{}
	if err = randomize.Struct(seed, monitoringPeerOne, monitoringPeerDBTypes, false, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, monitoringPeerTwo, monitoringPeerDBTypes, false, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = monitoringPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = monitoringPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func monitoringPeerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func monitoringPeerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MonitoringPeer) error {
	*o = MonitoringPeer{}
	return nil
}

func testMonitoringPeersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MonitoringPeer{}
	o := &MonitoringPeer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer object: %s", err)
	}

	AddMonitoringPeerHook(boil.BeforeInsertHook, monitoringPeerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	monitoringPeerBeforeInsertHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.AfterInsertHook, monitoringPeerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	monitoringPeerAfterInsertHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.AfterSelectHook, monitoringPeerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	monitoringPeerAfterSelectHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.BeforeUpdateHook, monitoringPeerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	monitoringPeerBeforeUpdateHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.AfterUpdateHook, monitoringPeerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	monitoringPeerAfterUpdateHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.BeforeDeleteHook, monitoringPeerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	monitoringPeerBeforeDeleteHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.AfterDeleteHook, monitoringPeerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	monitoringPeerAfterDeleteHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.BeforeUpsertHook, monitoringPeerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	monitoringPeerBeforeUpsertHooks = []MonitoringPeerHook{}

	AddMonitoringPeerHook(boil.AfterUpsertHook, monitoringPeerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	monitoringPeerAfterUpsertHooks = []MonitoringPeerHook{}
}

func testMonitoringPeersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMonitoringPeersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(monitoringPeerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMonitoringPeersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMonitoringPeersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MonitoringPeerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMonitoringPeersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MonitoringPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	monitoringPeerDBTypes = map[string]string{`ID`: `integer`, `PeerID`: `character varying`, `IPAddress`: `character varying`, `GeoLocation`: `character varying`, `NextDialAttempt`: `timestamp with time zone`, `MinDuration`: `interval`, `MaxDuration`: `interval`, `SuccessfulDials`: `integer`, `LastSuccessfulAt`: `timestamp with time zone`, `LastFailedAt`: `timestamp with time zone`}
	_                     = bytes.MinRead
)

func testMonitoringPeersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(monitoringPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(monitoringPeerAllColumns) == len(monitoringPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMonitoringPeersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(monitoringPeerAllColumns) == len(monitoringPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MonitoringPeer{}
	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, monitoringPeerDBTypes, true, monitoringPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(monitoringPeerAllColumns, monitoringPeerPrimaryKeyColumns) {
		fields = monitoringPeerAllColumns
	} else {
		fields = strmangle.SetComplement(
			monitoringPeerAllColumns,
			monitoringPeerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MonitoringPeerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMonitoringPeersUpsert(t *testing.T) {
	t.Parallel()

	if len(monitoringPeerAllColumns) == len(monitoringPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MonitoringPeer{}
	if err = randomize.Struct(seed, &o, monitoringPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MonitoringPeer: %s", err)
	}

	count, err := MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, monitoringPeerDBTypes, false, monitoringPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MonitoringPeer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MonitoringPeer: %s", err)
	}

	count, err = MonitoringPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

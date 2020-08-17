// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOutputAchievements(t *testing.T) {
	t.Parallel()

	query := OutputAchievements()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOutputAchievementsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
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

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OutputAchievements().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OutputAchievementSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputAchievementsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OutputAchievementExists(ctx, tx, o.OutputAchievementID)
	if err != nil {
		t.Errorf("Unable to check if OutputAchievement exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OutputAchievementExists to return true, but got false.")
	}
}

func testOutputAchievementsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	outputAchievementFound, err := FindOutputAchievement(ctx, tx, o.OutputAchievementID)
	if err != nil {
		t.Error(err)
	}

	if outputAchievementFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOutputAchievementsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OutputAchievements().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOutputAchievementsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OutputAchievements().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOutputAchievementsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	outputAchievementOne := &OutputAchievement{}
	outputAchievementTwo := &OutputAchievement{}
	if err = randomize.Struct(seed, outputAchievementOne, outputAchievementDBTypes, false, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}
	if err = randomize.Struct(seed, outputAchievementTwo, outputAchievementDBTypes, false, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = outputAchievementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = outputAchievementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OutputAchievements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOutputAchievementsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	outputAchievementOne := &OutputAchievement{}
	outputAchievementTwo := &OutputAchievement{}
	if err = randomize.Struct(seed, outputAchievementOne, outputAchievementDBTypes, false, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}
	if err = randomize.Struct(seed, outputAchievementTwo, outputAchievementDBTypes, false, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = outputAchievementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = outputAchievementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func outputAchievementBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func outputAchievementAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OutputAchievement) error {
	*o = OutputAchievement{}
	return nil
}

func testOutputAchievementsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OutputAchievement{}
	o := &OutputAchievement{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OutputAchievement object: %s", err)
	}

	AddOutputAchievementHook(boil.BeforeInsertHook, outputAchievementBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementBeforeInsertHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.AfterInsertHook, outputAchievementAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementAfterInsertHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.AfterSelectHook, outputAchievementAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	outputAchievementAfterSelectHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.BeforeUpdateHook, outputAchievementBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	outputAchievementBeforeUpdateHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.AfterUpdateHook, outputAchievementAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	outputAchievementAfterUpdateHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.BeforeDeleteHook, outputAchievementBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	outputAchievementBeforeDeleteHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.AfterDeleteHook, outputAchievementAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	outputAchievementAfterDeleteHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.BeforeUpsertHook, outputAchievementBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementBeforeUpsertHooks = []OutputAchievementHook{}

	AddOutputAchievementHook(boil.AfterUpsertHook, outputAchievementAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	outputAchievementAfterUpsertHooks = []OutputAchievementHook{}
}

func testOutputAchievementsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputAchievementsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(outputAchievementColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputAchievementToManyOutputAchievementTags(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievement
	var b, c OutputAchievementTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, outputAchievementTagDBTypes, false, outputAchievementTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, outputAchievementTagDBTypes, false, outputAchievementTagColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OutputAchievementID = a.OutputAchievementID
	c.OutputAchievementID = a.OutputAchievementID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OutputAchievementTags().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OutputAchievementID == b.OutputAchievementID {
			bFound = true
		}
		if v.OutputAchievementID == c.OutputAchievementID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OutputAchievementSlice{&a}
	if err = a.L.LoadOutputAchievementTags(ctx, tx, false, (*[]*OutputAchievement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OutputAchievementTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OutputAchievementTags = nil
	if err = a.L.LoadOutputAchievementTags(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OutputAchievementTags); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOutputAchievementToManyAddOpOutputAchievementTags(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievement
	var b, c, d, e OutputAchievementTag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementDBTypes, false, strmangle.SetComplement(outputAchievementPrimaryKeyColumns, outputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OutputAchievementTag{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputAchievementTagDBTypes, false, strmangle.SetComplement(outputAchievementTagPrimaryKeyColumns, outputAchievementTagColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OutputAchievementTag{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOutputAchievementTags(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.OutputAchievementID != first.OutputAchievementID {
			t.Error("foreign key was wrong value", a.OutputAchievementID, first.OutputAchievementID)
		}
		if a.OutputAchievementID != second.OutputAchievementID {
			t.Error("foreign key was wrong value", a.OutputAchievementID, second.OutputAchievementID)
		}

		if first.R.OutputAchievement != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.OutputAchievement != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OutputAchievementTags[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OutputAchievementTags[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OutputAchievementTags().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOutputAchievementToOneTodoUsingTodo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OutputAchievement
	var foreign Todo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, todoDBTypes, false, todoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Todo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.TodoID, foreign.TodoID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Todo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.TodoID, foreign.TodoID) {
		t.Errorf("want: %v, got %v", foreign.TodoID, check.TodoID)
	}

	slice := OutputAchievementSlice{&local}
	if err = local.L.LoadTodo(ctx, tx, false, (*[]*OutputAchievement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Todo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Todo = nil
	if err = local.L.LoadTodo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Todo == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOutputAchievementToOneSetOpTodoUsingTodo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievement
	var b, c Todo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementDBTypes, false, strmangle.SetComplement(outputAchievementPrimaryKeyColumns, outputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, todoDBTypes, false, strmangle.SetComplement(todoPrimaryKeyColumns, todoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, todoDBTypes, false, strmangle.SetComplement(todoPrimaryKeyColumns, todoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Todo{&b, &c} {
		err = a.SetTodo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Todo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OutputAchievements[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.TodoID, x.TodoID) {
			t.Error("foreign key was wrong value", a.TodoID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TodoID))
		reflect.Indirect(reflect.ValueOf(&a.TodoID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.TodoID, x.TodoID) {
			t.Error("foreign key was wrong value", a.TodoID, x.TodoID)
		}
	}
}

func testOutputAchievementToOneRemoveOpTodoUsingTodo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OutputAchievement
	var b Todo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputAchievementDBTypes, false, strmangle.SetComplement(outputAchievementPrimaryKeyColumns, outputAchievementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, todoDBTypes, false, strmangle.SetComplement(todoPrimaryKeyColumns, todoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetTodo(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveTodo(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Todo().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Todo != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.TodoID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OutputAchievements) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOutputAchievementsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
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

func testOutputAchievementsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OutputAchievementSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOutputAchievementsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OutputAchievements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	outputAchievementDBTypes = map[string]string{`OutputAchievementID`: `int`, `Title`: `varchar`, `TodoID`: `int`, `ReferenceURL`: `varchar`, `Summary`: `varchar`, `CreatedBy`: `int`, `CreatedAt`: `timestamp`, `ModifiedBy`: `int`, `ModifiedAt`: `timestamp`}
	_                        = bytes.MinRead
)

func testOutputAchievementsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(outputAchievementPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(outputAchievementAllColumns) == len(outputAchievementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOutputAchievementsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(outputAchievementAllColumns) == len(outputAchievementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OutputAchievement{}
	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, outputAchievementDBTypes, true, outputAchievementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(outputAchievementAllColumns, outputAchievementPrimaryKeyColumns) {
		fields = outputAchievementAllColumns
	} else {
		fields = strmangle.SetComplement(
			outputAchievementAllColumns,
			outputAchievementPrimaryKeyColumns,
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

	slice := OutputAchievementSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOutputAchievementsUpsert(t *testing.T) {
	t.Parallel()

	if len(outputAchievementAllColumns) == len(outputAchievementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOutputAchievementUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OutputAchievement{}
	if err = randomize.Struct(seed, &o, outputAchievementDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OutputAchievement: %s", err)
	}

	count, err := OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, outputAchievementDBTypes, false, outputAchievementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OutputAchievement struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OutputAchievement: %s", err)
	}

	count, err = OutputAchievements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
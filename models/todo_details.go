// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// TodoDetail is an object representing the database table.
type TodoDetail struct {
	TodoDetailID int       `boil:"todo_detail_id" json:"todo_detail_id" toml:"todo_detail_id" yaml:"todo_detail_id"`
	Content      string    `boil:"content" json:"content" toml:"content" yaml:"content"`
	CreatedBy    int       `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Checked      null.Int8 `boil:"checked" json:"checked,omitempty" toml:"checked" yaml:"checked,omitempty"`
	TodoID       int       `boil:"todo_id" json:"todo_id" toml:"todo_id" yaml:"todo_id"`
	ModifiedBy   null.Int  `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedAt   null.Time `boil:"modified_at" json:"modified_at,omitempty" toml:"modified_at" yaml:"modified_at,omitempty"`

	R *todoDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L todoDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

func (t TodoDetail) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d]", t.Content, t.TodoID, t.CreatedBy, t.CreatedAt)
}

var TodoDetailColumns = struct {
	TodoDetailID string
	Content      string
	CreatedBy    string
	CreatedAt    string
	Checked      string
	TodoID       string
	ModifiedBy   string
	ModifiedAt   string
}{
	TodoDetailID: "todo_detail_id",
	Content:      "content",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	Checked:      "checked",
	TodoID:       "todo_id",
	ModifiedBy:   "modified_by",
	ModifiedAt:   "modified_at",
}

// Generated where

type whereHelpernull_Int8 struct{ field string }

func (w whereHelpernull_Int8) EQ(x null.Int8) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int8) NEQ(x null.Int8) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int8) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int8) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int8) LT(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int8) LTE(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int8) GT(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int8) GTE(x null.Int8) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TodoDetailWhere = struct {
	TodoDetailID whereHelperint
	Content      whereHelperstring
	CreatedBy    whereHelperint
	CreatedAt    whereHelpertime_Time
	Checked      whereHelpernull_Int8
	TodoID       whereHelperint
	ModifiedBy   whereHelpernull_Int
	ModifiedAt   whereHelpernull_Time
}{
	TodoDetailID: whereHelperint{field: "`todo_details`.`todo_detail_id`"},
	Content:      whereHelperstring{field: "`todo_details`.`content`"},
	CreatedBy:    whereHelperint{field: "`todo_details`.`created_by`"},
	CreatedAt:    whereHelpertime_Time{field: "`todo_details`.`created_at`"},
	Checked:      whereHelpernull_Int8{field: "`todo_details`.`checked`"},
	TodoID:       whereHelperint{field: "`todo_details`.`todo_id`"},
	ModifiedBy:   whereHelpernull_Int{field: "`todo_details`.`modified_by`"},
	ModifiedAt:   whereHelpernull_Time{field: "`todo_details`.`modified_at`"},
}

// TodoDetailRels is where relationship names are stored.
var TodoDetailRels = struct {
	Todo string
}{
	Todo: "Todo",
}

// todoDetailR is where relationships are stored.
type todoDetailR struct {
	Todo *Todo
}

// NewStruct creates a new relationship struct
func (*todoDetailR) NewStruct() *todoDetailR {
	return &todoDetailR{}
}

// todoDetailL is where Load methods for each relationship are stored.
type todoDetailL struct{}

var (
	todoDetailAllColumns            = []string{"todo_detail_id", "content", "created_by", "created_at", "checked", "todo_id", "modified_by", "modified_at"}
	todoDetailColumnsWithoutDefault = []string{"content", "created_by", "created_at", "checked", "todo_id", "modified_by", "modified_at"}
	todoDetailColumnsWithDefault    = []string{"todo_detail_id"}
	todoDetailPrimaryKeyColumns     = []string{"todo_detail_id"}
)

type (
	// TodoDetailSlice is an alias for a slice of pointers to TodoDetail.
	// This should generally be used opposed to []TodoDetail.
	TodoDetailSlice []*TodoDetail
	// TodoDetailHook is the signature for custom TodoDetail hook methods
	TodoDetailHook func(context.Context, boil.ContextExecutor, *TodoDetail) error

	todoDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	todoDetailType                 = reflect.TypeOf(&TodoDetail{})
	todoDetailMapping              = queries.MakeStructMapping(todoDetailType)
	todoDetailPrimaryKeyMapping, _ = queries.BindMapping(todoDetailType, todoDetailMapping, todoDetailPrimaryKeyColumns)
	todoDetailInsertCacheMut       sync.RWMutex
	todoDetailInsertCache          = make(map[string]insertCache)
	todoDetailUpdateCacheMut       sync.RWMutex
	todoDetailUpdateCache          = make(map[string]updateCache)
	todoDetailUpsertCacheMut       sync.RWMutex
	todoDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var todoDetailBeforeInsertHooks []TodoDetailHook
var todoDetailBeforeUpdateHooks []TodoDetailHook
var todoDetailBeforeDeleteHooks []TodoDetailHook
var todoDetailBeforeUpsertHooks []TodoDetailHook

var todoDetailAfterInsertHooks []TodoDetailHook
var todoDetailAfterSelectHooks []TodoDetailHook
var todoDetailAfterUpdateHooks []TodoDetailHook
var todoDetailAfterDeleteHooks []TodoDetailHook
var todoDetailAfterUpsertHooks []TodoDetailHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TodoDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TodoDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TodoDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TodoDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TodoDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TodoDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TodoDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TodoDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TodoDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range todoDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTodoDetailHook registers your hook function for all future operations.
func AddTodoDetailHook(hookPoint boil.HookPoint, todoDetailHook TodoDetailHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		todoDetailBeforeInsertHooks = append(todoDetailBeforeInsertHooks, todoDetailHook)
	case boil.BeforeUpdateHook:
		todoDetailBeforeUpdateHooks = append(todoDetailBeforeUpdateHooks, todoDetailHook)
	case boil.BeforeDeleteHook:
		todoDetailBeforeDeleteHooks = append(todoDetailBeforeDeleteHooks, todoDetailHook)
	case boil.BeforeUpsertHook:
		todoDetailBeforeUpsertHooks = append(todoDetailBeforeUpsertHooks, todoDetailHook)
	case boil.AfterInsertHook:
		todoDetailAfterInsertHooks = append(todoDetailAfterInsertHooks, todoDetailHook)
	case boil.AfterSelectHook:
		todoDetailAfterSelectHooks = append(todoDetailAfterSelectHooks, todoDetailHook)
	case boil.AfterUpdateHook:
		todoDetailAfterUpdateHooks = append(todoDetailAfterUpdateHooks, todoDetailHook)
	case boil.AfterDeleteHook:
		todoDetailAfterDeleteHooks = append(todoDetailAfterDeleteHooks, todoDetailHook)
	case boil.AfterUpsertHook:
		todoDetailAfterUpsertHooks = append(todoDetailAfterUpsertHooks, todoDetailHook)
	}
}

// One returns a single todoDetail record from the query.
func (q todoDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TodoDetail, error) {
	o := &TodoDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for todo_details")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TodoDetail records from the query.
func (q todoDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (TodoDetailSlice, error) {
	var o []*TodoDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TodoDetail slice")
	}

	if len(todoDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TodoDetail records in the query.
func (q todoDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count todo_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q todoDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if todo_details exists")
	}

	return count > 0, nil
}

// Todo pointed to by the foreign key.
func (o *TodoDetail) Todo(mods ...qm.QueryMod) todoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`todo_id` = ?", o.TodoID),
	}

	queryMods = append(queryMods, mods...)

	query := Todos(queryMods...)
	queries.SetFrom(query.Query, "`todos`")

	return query
}

// LoadTodo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (todoDetailL) LoadTodo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTodoDetail interface{}, mods queries.Applicator) error {
	var slice []*TodoDetail
	var object *TodoDetail

	if singular {
		object = maybeTodoDetail.(*TodoDetail)
	} else {
		slice = *maybeTodoDetail.(*[]*TodoDetail)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &todoDetailR{}
		}
		args = append(args, object.TodoID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &todoDetailR{}
			}

			for _, a := range args {
				if a == obj.TodoID {
					continue Outer
				}
			}

			args = append(args, obj.TodoID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`todos`), qm.WhereIn(`todos.todo_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Todo")
	}

	var resultSlice []*Todo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Todo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for todos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for todos")
	}

	if len(todoDetailAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Todo = foreign
		if foreign.R == nil {
			foreign.R = &todoR{}
		}
		foreign.R.TodoDetails = append(foreign.R.TodoDetails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TodoID == foreign.TodoID {
				local.R.Todo = foreign
				if foreign.R == nil {
					foreign.R = &todoR{}
				}
				foreign.R.TodoDetails = append(foreign.R.TodoDetails, local)
				break
			}
		}
	}

	return nil
}

// SetTodo of the todoDetail to the related item.
// Sets o.R.Todo to related.
// Adds o to related.R.TodoDetails.
func (o *TodoDetail) SetTodo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Todo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `todo_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"todo_id"}),
		strmangle.WhereClause("`", "`", 0, todoDetailPrimaryKeyColumns),
	)
	values := []interface{}{related.TodoID, o.TodoDetailID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TodoID = related.TodoID
	if o.R == nil {
		o.R = &todoDetailR{
			Todo: related,
		}
	} else {
		o.R.Todo = related
	}

	if related.R == nil {
		related.R = &todoR{
			TodoDetails: TodoDetailSlice{o},
		}
	} else {
		related.R.TodoDetails = append(related.R.TodoDetails, o)
	}

	return nil
}

// TodoDetails retrieves all the records using an executor.
func TodoDetails(mods ...qm.QueryMod) todoDetailQuery {
	mods = append(mods, qm.From("`todo_details`"))
	return todoDetailQuery{NewQuery(mods...)}
}

// FindTodoDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTodoDetail(ctx context.Context, exec boil.ContextExecutor, todoDetailID int, selectCols ...string) (*TodoDetail, error) {
	todoDetailObj := &TodoDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `todo_details` where `todo_detail_id`=?", sel,
	)

	q := queries.Raw(query, todoDetailID)

	err := q.Bind(ctx, exec, todoDetailObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from todo_details")
	}

	return todoDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TodoDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no todo_details provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(todoDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	todoDetailInsertCacheMut.RLock()
	cache, cached := todoDetailInsertCache[key]
	todoDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			todoDetailAllColumns,
			todoDetailColumnsWithDefault,
			todoDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(todoDetailType, todoDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(todoDetailType, todoDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `todo_details` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `todo_details` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `todo_details` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, todoDetailPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into todo_details")
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

	o.TodoDetailID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == todoDetailMapping["todo_detail_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TodoDetailID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for todo_details")
	}

CacheNoHooks:
	if !cached {
		todoDetailInsertCacheMut.Lock()
		todoDetailInsertCache[key] = cache
		todoDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TodoDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TodoDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	todoDetailUpdateCacheMut.RLock()
	cache, cached := todoDetailUpdateCache[key]
	todoDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			todoDetailAllColumns,
			todoDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update todo_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `todo_details` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, todoDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(todoDetailType, todoDetailMapping, append(wl, todoDetailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update todo_details row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for todo_details")
	}

	if !cached {
		todoDetailUpdateCacheMut.Lock()
		todoDetailUpdateCache[key] = cache
		todoDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q todoDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for todo_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for todo_details")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TodoDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todoDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `todo_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todoDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in todoDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all todoDetail")
	}
	return rowsAff, nil
}

var mySQLTodoDetailUniqueColumns = []string{
	"todo_detail_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TodoDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no todo_details provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(todoDetailColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTodoDetailUniqueColumns, o)

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

	todoDetailUpsertCacheMut.RLock()
	cache, cached := todoDetailUpsertCache[key]
	todoDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			todoDetailAllColumns,
			todoDetailColumnsWithDefault,
			todoDetailColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			todoDetailAllColumns,
			todoDetailPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert todo_details, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "todo_details", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `todo_details` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(todoDetailType, todoDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(todoDetailType, todoDetailMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for todo_details")
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

	o.TodoDetailID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == todoDetailMapping["todo_detail_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(todoDetailType, todoDetailMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for todo_details")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for todo_details")
	}

CacheNoHooks:
	if !cached {
		todoDetailUpsertCacheMut.Lock()
		todoDetailUpsertCache[key] = cache
		todoDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TodoDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TodoDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TodoDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), todoDetailPrimaryKeyMapping)
	sql := "DELETE FROM `todo_details` WHERE `todo_detail_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from todo_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for todo_details")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q todoDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no todoDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from todo_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for todo_details")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TodoDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(todoDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todoDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `todo_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todoDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from todoDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for todo_details")
	}

	if len(todoDetailAfterDeleteHooks) != 0 {
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
func (o *TodoDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTodoDetail(ctx, exec, o.TodoDetailID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TodoDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TodoDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), todoDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `todo_details`.* FROM `todo_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, todoDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TodoDetailSlice")
	}

	*o = slice

	return nil
}

// TodoDetailExists checks if the TodoDetail row exists.
func TodoDetailExists(ctx context.Context, exec boil.ContextExecutor, todoDetailID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `todo_details` where `todo_detail_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, todoDetailID)
	}
	row := exec.QueryRowContext(ctx, sql, todoDetailID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if todo_details exists")
	}

	return exists, nil
}

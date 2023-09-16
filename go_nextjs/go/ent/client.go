// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/0maru/x_lang_todo_list/go_nextjs/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/0maru/x_lang_todo_list/go_nextjs/ent/todo"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Todo is the client for interacting with the Todo builders.
	Todo *TodoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Todo = NewTodoClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Todo:   NewTodoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Todo:   NewTodoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Todo.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Todo.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Todo.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *TodoMutation:
		return c.Todo.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// TodoClient is a client for the Todo schema.
type TodoClient struct {
	config
}

// NewTodoClient returns a client for the Todo from the given config.
func NewTodoClient(c config) *TodoClient {
	return &TodoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todo.Hooks(f(g(h())))`.
func (c *TodoClient) Use(hooks ...Hook) {
	c.hooks.Todo = append(c.hooks.Todo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `todo.Intercept(f(g(h())))`.
func (c *TodoClient) Intercept(interceptors ...Interceptor) {
	c.inters.Todo = append(c.inters.Todo, interceptors...)
}

// Create returns a builder for creating a Todo entity.
func (c *TodoClient) Create() *TodoCreate {
	mutation := newTodoMutation(c.config, OpCreate)
	return &TodoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Todo entities.
func (c *TodoClient) CreateBulk(builders ...*TodoCreate) *TodoCreateBulk {
	return &TodoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Todo.
func (c *TodoClient) Update() *TodoUpdate {
	mutation := newTodoMutation(c.config, OpUpdate)
	return &TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoClient) UpdateOne(t *Todo) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodo(t))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoClient) UpdateOneID(id int) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodoID(id))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Todo.
func (c *TodoClient) Delete() *TodoDelete {
	mutation := newTodoMutation(c.config, OpDelete)
	return &TodoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoClient) DeleteOne(t *Todo) *TodoDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TodoClient) DeleteOneID(id int) *TodoDeleteOne {
	builder := c.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoDeleteOne{builder}
}

// Query returns a query builder for Todo.
func (c *TodoClient) Query() *TodoQuery {
	return &TodoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTodo},
		inters: c.Interceptors(),
	}
}

// Get returns a Todo entity by its id.
func (c *TodoClient) Get(ctx context.Context, id int) (*Todo, error) {
	return c.Query().Where(todo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoClient) GetX(ctx context.Context, id int) *Todo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryChildren queries the children edge of a Todo.
func (c *TodoClient) QueryChildren(t *Todo) *TodoQuery {
	query := (&TodoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, todo.ChildrenTable, todo.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Todo.
func (c *TodoClient) QueryParent(t *Todo) *TodoQuery {
	query := (&TodoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, todo.ParentTable, todo.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoClient) Hooks() []Hook {
	return c.hooks.Todo
}

// Interceptors returns the client interceptors.
func (c *TodoClient) Interceptors() []Interceptor {
	return c.inters.Todo
}

func (c *TodoClient) mutate(ctx context.Context, m *TodoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TodoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TodoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Todo mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Todo []ent.Hook
	}
	inters struct {
		Todo []ent.Interceptor
	}
)

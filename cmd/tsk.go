package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
	"tsk/internal/database"
	db "tsk/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	DB *db.Queries
}

type Params struct {
	Limit     int64
	Priority  string
	Tag       string
	All       bool
	Completed bool
	Pending   bool
}

func convertToSQLCParams(opts Params) interface{} {
	return opts
}

func NewTodo() *apiConfig {

	connection, err := sql.Open("sqlite3", "./tsk.db")

	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		priority TEXT NOT NULL CHECK(priority IN ('low', 'medium', 'high')) DEFAULT 'medium',
		tag TEXT,
		completed_at DATETIME,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	`
	_, err = connection.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		log.Fatal(err)
	}

	apiConfig := apiConfig{
		DB: db.New(connection),
	}
	return &apiConfig
}

func (apiConfig *apiConfig) AddTodo(name string, priority string, tag string) (database.Todo, error) {
	newTodo := db.CreateTodoParams{
		Name:        name,
		Completed:   false,
		Priority:    priority,
		Tag:         sql.NullString{String: tag, Valid: true},
		CompletedAt: sql.NullTime{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx := context.Background()

	return apiConfig.DB.CreateTodo(ctx, newTodo)
}

func (apiConfig *apiConfig) ListTodo(completed bool, pending bool, all bool, priority string, tag string, sortBy string, limit int64, orderType string) error {

	var tsk []database.Todo
	var err error

	priorityParams := db.GetTodosByFiltersByPriorityParams{
		All:       all,
		Completed: completed,
		Pending:   pending,
		Priority:  priority,
		Tag:       tag,
		Limit:     limit,
		Offset:    0,
		IsDesc:    1,
	}

	createdParams := db.GetTodosByFiltersByCreatedAtParams{
		All:       all,
		Completed: completed,
		Pending:   pending,
		Priority:  priority,
		Tag:       tag,
		Limit:     limit,
		Offset:    0,
		IsDesc:    1,
	}

	defaultParams := db.GetTodosByFiltersByIdParams{
		All:       all,
		Completed: completed,
		Pending:   pending,
		Priority:  priority,
		Tag:       tag,
		Limit:     limit,
		Offset:    0,
		IsDesc:    1,
	}

	if orderType == "ASC" {
		priorityParams.IsDesc = 0
		createdParams.IsDesc = 0
		defaultParams.IsDesc = 0
	} else if orderType != "DESC" {
		return errors.New("invalid order type: please use ASC or DESC")
	}

	switch sortBy {
	case "priority":
		tsk, err = apiConfig.DB.GetTodosByFiltersByPriority(context.Background(), priorityParams)
	case "created_at":
		tsk, err = apiConfig.DB.GetTodosByFiltersByCreatedAt(context.Background(), createdParams)
	default:
		tsk, err = apiConfig.DB.GetTodosByFiltersById(context.Background(), defaultParams)
	}

	if err != nil {
		return err
	}

	if len(tsk) == 0 {
		return errors.New("you have not added any tasks yet")
	}

	headers := []string{"Id", "Completed", "Name", "Priority", "Tag"}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, h := range headers {
		fmt.Fprintf(w, "%s\t", h)
	}
	fmt.Fprintf(w, "\n")
	for _, t := range tsk {
		fmt.Fprintf(w, "%d\t", t.ID)
		fmt.Fprintf(w, "%t\t", t.Completed)
		fmt.Fprintf(w, "%s\t", strings.Split(t.Name, " ")[0])
		fmt.Fprintf(w, "%s\t", t.Priority)
		if t.Tag.Valid {
			fmt.Fprintf(w, "%s\t", t.Tag.String)
		} else {
			fmt.Fprintf(w, " \t")
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()

	return nil
}

func (apiConfig *apiConfig) CompleteTodo(ID int64) error {
	completedTime := sql.NullTime{Time: time.Time{}, Valid: true}
	completedTime.Time = time.Now()

	completeTodoParams := db.CompleteTodoParams{
		ID:          ID,
		CompletedAt: completedTime,
		UpdatedAt:   time.Now(),
	}
	rows, error := apiConfig.DB.CompleteTodo(context.Background(), completeTodoParams)
	if rows == 0 {
		return fmt.Errorf("data does not exist with this ID: %d", ID)
	}
	return error
}

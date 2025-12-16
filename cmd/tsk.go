package cmd

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"
	db "tsk/internal/database"
)

type apiConfig struct {
	DB *db.Queries
}

func NewTodo() *apiConfig {

	connection, err := sql.Open("sqlite3", "../tsk.db")

	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
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

func (apiConfig *apiConfig) AddTodo(name string) {
	newTodo := db.CreateTodoParams{
		Name:        name,
		Completed:   false,
		CompletedAt: sql.NullTime{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx := context.Background()

	apiConfig.DB.CreateTodo(ctx, newTodo)
}

func (apiConfig *apiConfig) ListTodo() error {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
	fmt.Println("in the function")
	headers := []string{"ID", "Name", "Created At", "Updated At", "Completed", "Completed At"}

	formatRow := func(cells []string) string {
		return "| " + strings.Join(cells, " \t| ") + " \t|"
	}

	tsk, err := apiConfig.DB.GetAllTodos(context.Background())

	if err != nil {
		fmt.Println("error")
		return err
	}
	if len(tsk) == 0 {
		fmt.Println("You don't have any task yet. Please add any task first")
		return nil
	}
	fmt.Fprintln(w, formatRow(headers))

	for _, todo := range tsk {
		cells := []string{
			fmt.Sprintf("%d", todo.ID),
			todo.Name,
			todo.CreatedAt.Format("2006-01-02 15:04:05"),
			todo.UpdatedAt.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%v", todo.Completed),
			fmt.Sprintf("%v", todo.CompletedAt),
		}

		fmt.Fprintln(w, formatRow(cells))
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

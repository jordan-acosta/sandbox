package gqlgen

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/vektah/gqlgen/handler"
)

type API struct {
	todos []Todo
}

func NewAPI() {
	port := ":8080"
	api := &API{
		todos: []Todo{}, // Would normally be a reference to the db.
	}
	http.Handle("/", handler.Playground("Dataloader", "/query"))
	http.Handle("/query", handler.GraphQL(MakeExecutableSchema(api)))
	log.Println("Listening on ", port)
	log.Fatal(http.ListenAndServe(port, nil))
	return
}

func (api *API) Mutation_createTodo(ctx context.Context, text string) (todo Todo, err error) {
	todo = Todo{
		Text:   text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: fmt.Sprintf("U%d", rand.Int()),
	}
	api.todos = append(api.todos, todo)
	return
}

func (api *API) Query_todos(ctx context.Context) (todos []Todo, err error) {
	todos = api.todos
	return
}

func (api *API) Todo_user(ctx context.Context, todo *Todo) (user User, err error) {
	user = User{
		ID:   todo.UserID,
		Name: "user " + todo.UserID,
	}
	return
}

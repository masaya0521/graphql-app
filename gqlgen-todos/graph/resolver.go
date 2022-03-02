package graph

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
	"github.com/masaya0521/gqlgen-todos/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *sql.DB
}

// 追加
func (r *Resolver) createTodo(input model.NewTodo) *model.Todo {
	// dbにuserを保存する処理
	rand_id := fmt.Sprintf("T%d", rand.Int())
	_, err := r.DB.Exec(
		"INSERT INTO todo(id, t_text, user_id) VALUES($1,$2,$3)",
		rand_id,
		input.Text,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 省略しますが、dbに保存されたデータをここで取得して、idとnameという変数に保存
	var todo model.Todo = model.Todo{
		ID:   rand_id,
		Text: input.Text,
	}

	return &todo
}

func (r *Resolver) createUser(input model.NewUser) *model.User {
	// dbにuserを保存する処理
	rand_id := fmt.Sprintf("T%d", rand.Int())
	_, err := r.DB.Exec(
		"INSERT INTO users(id, t_name) VALUES($1,$2)",
		rand_id,
		input.Name,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 省略しますが、dbに保存されたデータをここで取得して、idとnameという変数に保存
	var user model.User = model.User{
		ID:   rand_id,
		Name: input.Name,
	}

	return &user
}

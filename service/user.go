package service

import (
	"time"

	// "github.com/kzpolicy/amb-todo/models"
	"github.com/volatiletech/sqlboiler/boil"
	"local.packages/models"

	"context"
	"log"
)

type UserService struct{}

// func (UserService) GetTodoList() models.UserSlice {
// 	ctx := context.Background()
// 	db := boil.GetContextDB()

// 	// データ取得
// 	todos, err := models.Todos().All(ctx, db)

// 	if err != nil {
// 		log.Fatalf("error %v", err)
// 		return nil
// 	}

// 	fmt.Print(json.Marshal(todos))

// 	return todos
// }

func (UserService) AddUser(user models.User) (err error) {
	ctx := context.Background()
	db := boil.GetContextDB()
	now := time.Now()

	// User登録
	var u models.User
	u.Email = user.Email
	u.Password = user.Password
	u.CreatedBy = 1
	u.CreatedAt = now
	err1 := t.Insert(ctx, db, boil.Infer())

	if err1 != nil {
		log.Fatalf("error %v", err1)
		return err1
	}

	return nil
}

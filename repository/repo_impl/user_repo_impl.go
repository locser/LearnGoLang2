package repo_impl

import (
	"LearnGoLang2/banana"
	"LearnGoLang2/db"
	"LearnGoLang2/model"
	"LearnGoLang2/repository"
	"context"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

/*
UserId    string    `db:"user_id, omitempty"`
FullName  string    `db:"full_name, omitempty"`
Email     string    `db:"email, omitempty"`
Password  string    `db:"password, omitempty"`
Role      string    `db:"role, omitempty"`
CreatedAt time.Time `db:"created_at, omitempty"`
UpdatedAt time.Time `db:"updated_at, omitempty"`
Token     string
*/
func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statment := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)		
		VALUES (:user_id, :email, :password, :role, :full_name, :created_at,  :updated_at )
	`

	_, err := u.sql.Db.NamedExecContext(context, statment, user)

	if err != nil {
		//log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	return user, nil
}

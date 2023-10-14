package users_repo

import (
	"apipost/model/user"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

const TableName = "users"

type Repo struct {
	ctx  context.Context
	conn *pgx.Conn
}

func New(ctx context.Context, conn *pgx.Conn) *Repo {
	return &Repo{
		ctx:  ctx,
		conn: conn,
	}
}

func (r *Repo) Select(query string) user.List {
	rows, err := r.conn.Query(r.ctx, fmt.Sprintf("select id, user_name, user_sur_name from users"))
	if err != nil {
		return user.List{}
	}
	defer rows.Close()

	res := user.List{}

	for rows.Next() {
		usr := user.Data{}
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Surname); err != nil {
			return user.List{}
		}
		res = append(res, usr)
	}
	return res
}

func (r *Repo) Get(query string) user.Data {
	return user.NewUser(0, "", "")
}

func (r *Repo) GET(uid int) user.Data {
	rows, err := r.conn.Query(r.ctx, fmt.Sprintf("select id, user_name, user_sur_name from users where id = %d limit 1", uid))
	if err != nil {
		return user.Data{}
	}
	defer rows.Close()

	res := user.Data{}
	for rows.Next() {
		if err := rows.Scan(&res.Id, &res.Name, &res.Surname); err != nil {
			return user.Data{}
		}
	}
	return res
}

/*
	rows, err := conn.Query(ctx, "select user_name, user_sur_name, id from users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.name, &user.surName, &user.id); err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Printf("id: %d, name: %s, sur_name: %s\n", user.id, user.name, user.surName)
	}

*/

package main

import (
	"context"
	"fmt"
	"log"
	"training-go/go-session6-db-pgx/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := "postgresql://postgres:P4ssw0rd@localhost:5432/training_golang"
	ctx := context.Background()
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	//query untuk mengambil row
	var u entity.User
	err = pool.QueryRow(ctx, "SELECT id, name, email FROM users order by id desc limit 1", 1).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user retrieved", u)

	//exec untuk menjalankan task insert/update/delete
	_, err = pool.Exec(ctx, "INSERT INTO users(name, email, password,created_at,updated_at) VALUES('test', 'yk@gmail.com', '123', now(), now())")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user inserted", u)

	//query untuk mengambil banyak row
	var users []entity.User
	rows, err := pool.Query(ctx, "SELECT id, name, email FROM users order by id desc")
	if err != nil {
		log.Panicln(err)
	}
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println("all users retrieved", users)
}

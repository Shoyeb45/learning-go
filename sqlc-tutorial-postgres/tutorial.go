package main

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"tutorial.sqlc.dev/app/tutorial"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// list all authors
	authors, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateUsers(ctx, tutorial.CreateUsersParams{
		Name: "Shoyeb Ansari",
		Dob:  pgtype.Date{
			Time: time.Date(2004, 6, 17, 0, 0, 0, 0, time.UTC),
			Valid: true,
		},
	});

	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetUsers(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

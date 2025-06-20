package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aserafim/golang_sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	queries := db.New(conn)
	// queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend with Go",
	// 	Description: sql.NullString{String: "An introductory course on backend with Golang", Valid: true},
	// })

	//queries.DeleteCategory(ctx, "fa65ee18-a807-4caa-98ab-06d27101c654")

	// queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Complete Frontend",
	// 	Description: sql.NullString{String: "A complete course in HTML, CSS and JS", Valid: true},
	// })

	// queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "UX/UI",
	// 	Description: sql.NullString{String: "All that you need to know about User Experience", Valid: true},
	// })

	// queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Data Science Course",
	// 	Description: sql.NullString{String: "Learn about analise all your data-sets", Valid: true},
	// })

	// queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Blockchain",
	// 	Description: sql.NullString{String: "Understand the misterious world of Crypto Concurrencies", Valid: true},
	// })

	// queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "e38e2f77-cf4e-48f9-a45d-e469e2dc5c3a",
	// 	Name:        "Backend + Frontend",
	// 	Description: sql.NullString{String: "A complete course on HTML, CSS, JS and NodeJS", Valid: true},
	// })

	// cat, err := queries.GetCategory(ctx, "e38e2f77-cf4e-48f9-a45d-e469e2dc5c3a")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("ID: %s\nName: %s\nDescription: %s\n", cat.ID, cat.Name, cat.Description.String)

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, cat := range categories {
		fmt.Printf("ID: %s\nName: %s\nDescription: %s\n\n", cat.ID, cat.Name, cat.Description.String)
	}
}

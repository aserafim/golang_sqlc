package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aserafim/golang_sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CourseParams struct {
	ID    string
	Name  string
	Price float64
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {

	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		_, err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:         argsCourse.ID,
			Name:       argsCourse.Name,
			Price:      argsCourse.Price,
			CategoryID: argsCategory.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil
	}
	return nil
}

func main() {
	ctx := context.Background()

	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	queries := db.New(conn)

	courses, err := queries.ListCourses(ctx)

	if err != nil {
		panic(err)
	}

	for _, c := range courses {
		fmt.Printf("Category: %s\nCourse ID: %s\nCourse name: %s\nCourse Price: %f\n\n", c.CategoryName, c.ID, c.Name, c.Price)
	}

	// courseArgs := CourseParams{
	// 	ID:    uuid.New().String(),
	// 	Name:  "Python to Data Science",
	// 	Price: 685,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Data-science",
	// 	Description: sql.NullString{String: "Data-science Course", Valid: true},
	// }

	// courseDB := NewCourseDB(conn)
	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	// if err != nil {
	// 	panic(err)
	// }

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

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, cat := range categories {
	// 	fmt.Printf("ID: %s\nName: %s\nDescription: %s\n\n", cat.ID, cat.Name, cat.Description.String)
	// }

}

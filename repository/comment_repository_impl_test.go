package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comment := entity.Comment{
		Email:   "admin1@mail.com",
		Comment: "Halo Guys",
	}
	ctx := context.Background()
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

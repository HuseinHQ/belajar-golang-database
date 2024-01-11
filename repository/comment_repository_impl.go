package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES($1, $2) RETURNING id"
	var id int
	err := repository.DB.QueryRowContext(ctx, script, comment.Email, comment.Comment).Scan(&id)
	if err != nil {
		return comment, err
	}

	comment.Id = id
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = $1"
	var comment entity.Comment

	rows, err := repository.DB.QueryContext(ctx, script, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(id) + " not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	var comments []entity.Comment

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}

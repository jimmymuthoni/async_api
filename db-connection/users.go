package dbconnection

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	db *sqlx.DB
}

type User struct {
	Id 						uuid.UUID `db:"id"`
	Email 					string `db:"email"`
	HashedPasswordBase64 	string `db:"hashed_password"`
	CreatedAt 				time.Time `db:"created_at"`
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: sqlx.NewDb(db, "postgres"),
	}

}

func (u *User) ComparePassword(password string) error {
	hashedPassword, err := base64.StdEncoding.DecodeString(u.HashedPasswordBase64)
	if err != nil {
		return fmt.Errorf("failed to decode password hash: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password")
	}
	return nil
}

//creating user
func (s *UserStore) CreateUser(ctx context.Context, email, password string) (*User, error){
	const dml = `INSERT INTO users (email, hashed_password) VALUES ($1, $2) RETURNING *;`
	var user User

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	 }
	 hashedPasswordbase64 := base64.StdEncoding.EncodeToString(bytes)

	if err := s.db.GetContext(ctx, &user, dml, email, hashedPasswordbase64); err != nil {
		return nil, fmt.Errorf("Failed to insert user: %w", err)
	}
	return &user, nil
}

//fetching users by user email and user id
func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (*User, error){
	const query = `SELECT * FROM users WHERE email = $1;`
	var user User
	if err := s.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, fmt.Errorf("Failed to fetch user: %w", err)
	}
	return &user, nil
}

func (s *UserStore) GetUserById(ctx context.Context, userId uuid.UUID) (*User, error){
	const query = `SELECT * FROM users WHERE id = $1;`
	var user User
	if err := s.db.GetContext(ctx, &user, query, userId); err != nil {
		return nil, fmt.Errorf("Failed to fetch user by id %s: %w", userId, err)
	}
	return &user, nil
}
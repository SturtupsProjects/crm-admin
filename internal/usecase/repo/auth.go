package repo

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) usecase.UsersRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) AddAdmin(in entity.AdminPass) (entity.Message, error) {
	res := entity.Message{}

	_, err := u.db.Exec(`insert into users(first_name, last_name, email, phone_number, password, role)
values ($1, $2, $3, $4, $5, $6)`, "admin", "admin", "admin", in.Login, in.Password, "admin")
	if err != nil {
		return res, err
	}

	res.Message = "Owner added"

	return res, nil
}

func (u *UserRepo) CreateUser(in entity.User) (entity.UserRequest, error) {
	var user entity.UserRequest
	query := `
		INSERT INTO users (first_name, last_name, email, phone_number, password, role)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING user_id, first_name, last_name, email, phone_number, role, created_at
	`
	err := u.db.QueryRowx(query, in.FirstName, in.LastName, in.Email, in.PhoneNumber, in.Password, in.Role).
		Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return entity.UserRequest{}, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil
}

// GetUser retrieves a user by their ID.
func (u *UserRepo) GetUser(in entity.UserID) (entity.UserRequest, error) {
	var user entity.UserRequest
	query := `SELECT user_id, first_name, last_name, email, phone_number, role, created_at FROM users WHERE user_id = $1`
	err := u.db.Get(&user, query, in.ID)
	if err != nil {
		return entity.UserRequest{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// GetListUser retrieves a list of users based on filter criteria.
func (u *UserRepo) GetListUser(in entity.FilterUser) (entity.UserList, error) {
	var users []entity.UserRequest
	query := `
		SELECT user_id, first_name, last_name, email, phone_number, role, created_at
		FROM users
		WHERE ($1::VARCHAR IS NULL OR first_name ILIKE '%' || $1 || '%')
		  AND ($2::VARCHAR IS NULL OR last_name ILIKE '%' || $2 || '%')
		  AND ($3::VARCHAR IS NULL OR role = $3)
		ORDER BY created_at DESC
	`
	err := u.db.Select(&users, query, in.FirstName, in.LastName, in.Role)
	if err != nil {
		return entity.UserList{}, fmt.Errorf("failed to list users: %w", err)
	}
	return entity.UserList{Users: users}, nil
}

// DeleteUser removes a user by their ID.
func (u *UserRepo) DeleteUser(in entity.UserID) (entity.Message, error) {
	query := `DELETE FROM users WHERE user_id = $1`
	res, err := u.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete user: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d user(s)", rows)}, nil
}

// UpdateUser modifies the fields of a user based on the fields provided in UserRequest.
func (u *UserRepo) UpdateUser(in entity.UserRequest) (entity.UserRequest, error) {
	var user entity.UserRequest
	query := `UPDATE users SET `
	var args []interface{}
	argCounter := 1

	// Dynamically build the query based on non-empty fields
	if in.FirstName != "" {
		query += fmt.Sprintf("first_name = $%d, ", argCounter)
		args = append(args, in.FirstName)
		argCounter++
	}
	if in.LastName != "" {
		query += fmt.Sprintf("last_name = $%d, ", argCounter)
		args = append(args, in.LastName)
		argCounter++
	}
	if in.Email != "" {
		query += fmt.Sprintf("email = $%d, ", argCounter)
		args = append(args, in.Email)
		argCounter++
	}
	if in.PhoneNumber != "" {
		query += fmt.Sprintf("phone_number = $%d, ", argCounter)
		args = append(args, in.PhoneNumber)
		argCounter++
	}
	if in.Role != "" {
		query += fmt.Sprintf("role = $%d, ", argCounter)
		args = append(args, in.Role)
		argCounter++
	}

	// Remove trailing comma and add WHERE clause
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE user_id = $%d RETURNING user_id, first_name, last_name, email, phone_number, role, created_at", argCounter)
	args = append(args, in.UserID)

	// Execute the query
	err := u.db.QueryRowx(query, args...).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err != nil {
		return entity.UserRequest{}, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (u *UserRepo) LogIn(in entity.PhoneNumber) (entity.LogInReq, error) {
	res := entity.LogInReq{}

	err := u.db.Get(&res, `select user_id, first_name, phone_number, role from users
	where phone_number = $1`, in.PhoneNumber)

	if err != nil {
		return entity.LogInReq{}, err
	}

	return res, nil
}

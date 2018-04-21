package models

import (
	"errors"
	"time"
)

// User user
type User struct {
	ID        string
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Users users
type Users []User

// AllUsers allUsers
func AllUsers() Users {
	rows, err := db.Query("SELECT * FROM sistema.user")
	defer rows.Close()

	checkErr(err)

	users := Users{}
	user := User{}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		checkErr(err)

		users = append(users, user)
	}

	return users
}

// FindUser findUser
func FindUser(id string) (*User, error) {
	if validate.Var(id, "required") != nil {
		return nil, errors.New("Código é obrigatório")
	}

	stmt, err := db.Prepare("SELECT * FROM sistema.user WHERE id=$1")
	defer stmt.Close()

	checkErr(err)

	user := User{}
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, nil
	}

	return &user, nil
}

// CreateUser createUser
func CreateUser(user User) (*User, error) {
	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("INSERT INTO sistema.user (name, email) VALUES ($1,$2)")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		return nil, nil
	}

	row, err := db.Query("SELECT * FROM sistema.user ORDER BY created_at DESC LIMIT 1")
	defer row.Close()

	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		checkErr(err)
	}

	return &user, nil
}

// UpdateUser updateUser
func UpdateUser(id string, user User) (*User, error) {
	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("UPDATE sistema.user SET name=$1, email=$2 WHERE id=$3")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(user.Name, user.Email, id)
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		// erro - nenhum registro atualizado
		return nil, nil
	}

	u, _ := FindUser(id)
	return u, nil
}

// RemoveUser removeUser
func RemoveUser(id string) error {

	if validate.Var(id, "required") != nil {
		return errors.New("Código é obrigatório")
	}

	stmt, err := db.Prepare("DELETE FROM sistema.user WHERE id =$1")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	checkErr(err)

	if rows == 1 {
		return nil
	}

	return errors.New("Algo de errado aconteceu")
}

// GetUserTags getUserTags
func GetUserTags(id string) Tags {
	rows, err := db.Query("SELECT * FROM sistema.tag WHERE id_user=$1", id)
	defer rows.Close()

	checkErr(err)

	tags := Tags{}
	tag := Tag{}

	for rows.Next() {
		err = rows.Scan(&tag.ID, &tag.IDUser, &tag.Tag, &tag.Description, &tag.State, &tag.CreatedAt, &tag.UpdatedAt)
		checkErr(err)

		tags = append(tags, tag)
	}

	return tags
}

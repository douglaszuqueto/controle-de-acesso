package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Tag tag
type Tag struct {
	ID          string
	IDUser      string `validate:"required"`
	Tag         string `validate:"required"`
	Description string
	State       int
	User        *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Tags tags
type Tags []Tag

// AllTags allTags
func AllTags() *Tags {
	rows, err := Query("SELECT * FROM sistema.tag")
	defer rows.Close()

	checkErr(err)

	tags := Tags{}
	tag := Tag{}

	for rows.Next() {
		err = rows.Scan(&tag.ID, &tag.IDUser, &tag.Tag, &tag.Description, &tag.State, &tag.CreatedAt, &tag.UpdatedAt)
		checkErr(err)

		user, _ := FindUser(tag.IDUser)
		tag.User = user

		tags = append(tags, tag)
	}

	return &tags
}

// FindTag findTag
func FindTag(id string) (*Tag, error) {
	if validate.Var(id, "required") != nil {
		return nil, errors.New("Código é obrigatório")
	}

	stmt, err := db.Prepare("SELECT * FROM sistema.tag WHERE id=$1")
	defer stmt.Close()

	checkErr(err)

	tag := Tag{}
	err = stmt.QueryRow(id).Scan(&tag.ID, &tag.IDUser, &tag.Tag, &tag.Description, &tag.State, &tag.CreatedAt, &tag.UpdatedAt)
	if err != nil {
		return nil, nil
	}

	user, _ := FindUser(tag.IDUser)
	tag.User = user

	return &tag, nil
}

// CreateTag createTag
func CreateTag(tag Tag) (*Tag, error) {
	err := validate.Struct(tag)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("INSERT INTO sistema.tag (id_user, tag, description, state) VALUES ($1,$2,$3,$4)")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(tag.IDUser, strings.ToUpper(tag.Tag), tag.Description, tag.State)
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		return nil, nil
	}

	row, err := db.Query("SELECT * FROM sistema.tag ORDER BY created_at DESC LIMIT 1")
	defer row.Close()

	checkErr(err)

	for row.Next() {
		err = row.Scan(&tag.ID, &tag.IDUser, &tag.Tag, &tag.Description, &tag.State, &tag.CreatedAt, &tag.UpdatedAt)
		checkErr(err)

		user, _ := FindUser(tag.IDUser)
		tag.User = user
	}

	return &tag, nil
}

// UpdateTag updateTag
func UpdateTag(id string, tag Tag) (*Tag, error) {
	err := validate.Struct(tag)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("UPDATE sistema.tag SET id_user=$1, description=$2, state=$3 WHERE id=$4")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(tag.IDUser, tag.Description, tag.State, id)
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		// erro - nenhum registro atualizado
		return nil, nil
	}
	t, _ := FindTag(id)

	return t, nil
}

// RemoveTag removeTag
func RemoveTag(id string) error {
	if validate.Var(id, "required") != nil {
		return errors.New("Código é obrigatório")
	}

	stmt, err := db.Prepare("DELETE FROM sistema.tag WHERE id =$1")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	rows, err := res.RowsAffected()
	checkErr(err)

	if rows == 1 {
		return nil
	}

	return errors.New("Algo de errado aconteceu")
}

// Ping struct
type Ping struct {
	Device   string
	Tag      string
	User     string
	Email    string
	Gravatar string
	State    int
}

// VerifyTag verifyTag
func VerifyTag(tag, idChip string) *Ping {
	query := `
        SELECT 
            d.name as device, t.tag, u.name as user, u.email, td.state, t.state as tag_state
        FROM 
            sistema.tag_device td 
        INNER JOIN sistema.device d ON td.id_device = d.id 
        INNER JOIN sistema.tag t on t.id = td.id_tag
        INNER JOIN sistema.user u on u.id = t.id_user
        WHERE d.state = 1 and t.tag =$1 and d.chip_id=$2`

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	checkErr(err)

	ping := Ping{}

	var tdState, tState int

	err = stmt.QueryRow(tag, idChip).Scan(&ping.Device, &ping.Tag, &ping.User, &ping.Email, &tdState, &tState)

	if err == sql.ErrNoRows {
		ping.State = 0
		return &ping
	}

	checkErr(err)

	hash := md5.New()
	hash.Write([]byte(ping.Email))

	ping.Gravatar = fmt.Sprintf("https://secure.gravatar.com/avatar/%s", hex.EncodeToString(hash.Sum(nil)))

	if tdState == 0 {
		ping.State = 2
		return &ping
	}

	if tState == 0 {
		ping.State = 3
		return &ping
	}

	ping.State = 1
	return &ping
}

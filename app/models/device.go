package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Device struct
type Device struct {
	ID          string
	Name        string `validate:"required"`
	Description string `validate:"required"`
	ChipID      sql.NullString
	State       int
	CreatedAt   time.Time
	Created     string
	UpdatedAt   time.Time
}

// Devices devices
type Devices []Device

// AllDevices allDevices
func AllDevices() *Devices {
	rows, err := Query("SELECT * FROM sistema.device")
	defer rows.Close()

	checkErr(err)

	devices := Devices{}
	device := Device{}

	for rows.Next() {
		err = rows.Scan(&device.ID, &device.Name, &device.Description, &device.ChipID, &device.State, &device.CreatedAt, &device.UpdatedAt)
		checkErr(err)

		devices = append(devices, device)
	}

	return &devices
}

// FindDevice findDevice
func FindDevice(id string) *Device {
	stmt, err := db.Prepare("SELECT * FROM sistema.device WHERE id=$1")
	defer stmt.Close()

	checkErr(err)

	device := Device{}
	err = stmt.QueryRow(id).Scan(&device.ID, &device.Name, &device.Description, &device.ChipID, &device.State, &device.CreatedAt, &device.UpdatedAt)

	if err != nil {
		return nil
	}

	return &device
}

// DeviceTags struct
type DeviceTags struct {
	Device *Device
	Tags   Tags
}

// FindTags findTags
func FindTags(id string) DeviceTags {
	m := DeviceTags{}
	m.Device = FindDevice(id)

	query := `
        SELECT 
            t.id, t.id_user, t.tag, t.description, td.state, td.created_at, td.updated_at
        FROM 
            sistema.tag_device td 
        INNER JOIN sistema.tag t on t.id = td.id_tag
        WHERE td.id_device=$1`

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	checkErr(err)

	rows, err := stmt.Query(id)
	checkErr(err)

	tags := Tags{}
	tag := Tag{}

	for rows.Next() {
		err = rows.Scan(&tag.ID, &tag.IDUser, &tag.Tag, &tag.Description, &tag.State, &tag.CreatedAt, &tag.UpdatedAt)
		checkErr(err)

		tags = append(tags, tag)
	}

	m.Tags = tags

	return m
}

// CreateDevice createDevice
func CreateDevice(device Device) (*Device, error) {
	err := validate.Struct(device)

	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("INSERT INTO sistema.device (name, description, state, chip_id) VALUES ($1,$2,$3,$4)")
	defer stmt.Close()

	checkErr(err)

	res, err := stmt.Exec(device.Name, device.Description, device.State, strings.ToUpper(device.ChipID.String))
	if err != nil {
		return nil, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		return nil, nil
	}

	row, err := db.Query("SELECT * FROM sistema.device ORDER BY created_at DESC LIMIT 1")
	defer row.Close()

	checkErr(err)

	for row.Next() {
		err = row.Scan(&device.ID, &device.Name, &device.Description, &device.ChipID, &device.State, &device.CreatedAt, &device.UpdatedAt)
		checkErr(err)
	}

	return &device, nil
}

// UpdateDevice updateDevice
func UpdateDevice(id string, device Device) (*Device, error) {
	err := validate.Struct(device)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("UPDATE sistema.device SET name=$1, description=$2, state=$3, chip_id=$4 WHERE id=$5")
	defer stmt.Close()

	checkErr(err)

	_, err = stmt.Exec(device.Name, device.Description, device.State, device.ChipID.String, id)
	if err != nil {
		return nil, err
	}

	return FindDevice(id), nil
}

// RemoveDevice removeDevice
func RemoveDevice(id string) error {
	if validate.Var(id, "required") != nil {
		return errors.New("Código é obrigatório")
	}

	stmt, err := db.Prepare("DELETE FROM sistema.device WHERE id =$1")
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

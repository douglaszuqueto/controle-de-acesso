package models

import (
	"errors"
)

// TagDevice struct
type TagDevice struct {
	IDDevice string `validate:"required"`
	IDTag    string `validate:"required"`
}

func findTagDevice(idTag string, idDevice string) int64 {
	stmt, err := db.Prepare("SELECT * FROM sistema.tag_device WHERE id_tag=$1 and id_device=$2")
	defer stmt.Close()

	checkErr(err)

	rows, err := stmt.Exec(idTag, idDevice)
	if err != nil {
		return 0
	}

	count, _ := rows.RowsAffected()

	return count
}

// AttachTagDevice attachTagDevice
func AttachTagDevice(tagDevice TagDevice) (bool, error) {
	err := validate.Struct(tagDevice)

	if err != nil {
		return false, err
	}

	count := findTagDevice(tagDevice.IDTag, tagDevice.IDDevice)

	if count >= 1 {
		return false, nil
	}

	stmt, err := db.Prepare("INSERT INTO sistema.tag_device (id_device, id_tag) VALUES ($1,$2)")
	defer stmt.Close()

	if err != nil {
		return false, err
	}

	res, err := stmt.Exec(tagDevice.IDDevice, tagDevice.IDTag)
	if err != nil {
		return false, err
	}

	rows, _ := res.RowsAffected()

	if rows != 1 {
		return false, nil
	}

	return true, nil
}

// DettachTagDevice dettachTagDevice
func DettachTagDevice(tagDevice TagDevice) error {
	err := validate.Struct(tagDevice)

	if err != nil {
		return err
	}

	count := findTagDevice(tagDevice.IDTag, tagDevice.IDDevice)
	if count != 1 {
		return errors.New("Nenhum regisro encontrado")
	}

	stmt, err := db.Prepare("DELETE FROM sistema.tag_device WHERE id_device=$1 and id_tag=$2")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(tagDevice.IDDevice, tagDevice.IDTag)
	if err != nil {
		return err
	}

	return nil
}

// TagActivateDeactivate tagActivateDeactivate
func TagActivateDeactivate(id, idTag string) error {
	count := findTagDevice(idTag, id)
	if count != 1 {
		return errors.New("Nenhum regisro encontrado")
	}

	stmt, err := db.Prepare("UPDATE sistema.tag_device SET state = CASE WHEN state = 1 THEN 0 ELSE 1 END WHERE id_device=$1 and id_tag=$2")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, idTag)
	if err != nil {
		return err
	}

	return nil
}

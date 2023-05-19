package db

import (
	"database/sql"
)

func DbGetMailRequests(id int, offset, pageSize int, db *sql.DB) (*sql.Rows, error) {
	var query string
	var args []interface{}

	if offset < 0 || pageSize < 0 {
		query = "SELECT n2mr.id, n2mr.inserttime, n2mr.updatetime, n2mr.mail_config FROM not2_mail_request n2mr JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.not_user=$1"
		args = append(args, id)
	} else {
		query = "SELECT n2mr.id, n2mr.inserttime, n2mr.updatetime, n2mr.mail_config FROM not2_mail_request n2mr JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.not_user=$1 OFFSET $2 LIMIT $3"
		args = append(args, id, offset, pageSize)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func DbGetMailItemsByRequest(uid string, id int, offset, pageSize int, db *sql.DB) (*sql.Rows, error) {
	var query string
	var args []interface{}

	if offset < 0 || pageSize < 0 {
		query = "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr ON n2mr.id = item.mail_request JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.notification_id =$1 AND n2r.not_user =$2"
		args = append(args, uid, id)
	} else {
		query = "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr ON n2mr.id = item.mail_request JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.notification_id =$1 AND n2r.not_user =$2 OFFSET $3 LIMIT $4"
		args = append(args, uid, id, offset, pageSize)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func DbGetMailItemsByMailConfigId(confId int, id int, offset, pageSize int, db *sql.DB) (*sql.Rows, error) {
	var query string
	var args []interface{}

	if offset < 0 || pageSize < 0 {
		query = "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr on item.mail_request = n2mr.id JOIN not2_request n2r on n2mr.request_id = n2r.id WHERE n2mr.mail_config =$1 and n2r.not_user = $2"
		args = append(args, confId, id)
	} else {
		query = "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr on item.mail_request = n2mr.id JOIN not2_request n2r on n2mr.request_id = n2r.id WHERE n2mr.mail_config =$1 and n2r.not_user = $2 OFFSET $3 LIMIT $4"
		args = append(args, confId, id, offset, pageSize)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

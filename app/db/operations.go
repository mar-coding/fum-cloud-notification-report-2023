package db

import (
	"database/sql"
)

func DbGetMailRequests(id int, db *sql.DB) (*sql.Rows, error) {
	query := "SELECT n2mr.id, n2mr.inserttime, n2mr.updatetime, n2mr.mail_config FROM not2_mail_request n2mr JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.not_user=$1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func DbGetMailItemsByRequest(uid string, id int, db *sql.DB) (*sql.Rows, error) {
	query := "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr ON n2mr.id = item.mail_request JOIN not2_request n2r ON n2mr.request_id = n2r.id WHERE n2r.notification_id =$1 AND n2r.not_user =$2"
	rows, err := db.Query(query, uid, id)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func DbGetMailItemsByMailConfigId(confId int, id int, db *sql.DB) (*sql.Rows, error) {
	query := "SELECT item.message_id, item.state, item.inserttime, item.updatetime, item.receiver, n2mr.mail_config FROM not2_mail_item item JOIN not2_mail_request n2mr on item.mail_request = n2mr.id JOIN not2_request n2r on n2mr.request_id = n2r.id WHERE n2mr.mail_config =$1 and n2r.not_user = $2"
	rows, err := db.Query(query, confId, id)
	if err != nil {
		return nil, err
	}
	return rows, err
}

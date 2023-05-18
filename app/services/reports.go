package services

import (
	"database/sql"

	"github.com/mar-coding/fum-cloud-notification-report-2023/app/db"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/models"
)

func GetMailRequests(id int, sqlDB *sql.DB) []models.OutputEmail {
	rows, err := db.DbGetMailRequests(id, sqlDB)
	defer rows.Close()

	var results []models.OutputEmail
	for rows.Next() {
		var output models.OutputEmail
		err := rows.Scan(&output.MessageId, &output.RequestInsertTimestamp,
			&output.RequestUpdateTimestamp, &output.Mail_config)
		if err != nil {
			panic(err)
		}
		results = append(results, output)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return results
}

func GetMailItemsByRequest(uid string, id int, sqlDB *sql.DB) []models.OutputReq {
	rows, err := db.DbGetMailItemsByRequest(uid, id, sqlDB)
	defer rows.Close()

	var results []models.OutputReq
	for rows.Next() {
		var output models.OutputReq
		err := rows.Scan(&output.MessageId, &output.RequestState,
			&output.ItemInsertTimestamp, &output.ItemUpdateTimestamp,
			&output.Receiver, &output.MailConfigId)
		if err != nil {
			panic(err)
		}
		results = append(results, output)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return results
}

func GetMailItemsByMailConfigId(confId int, id int, sqlDB *sql.DB) []models.OutputReq {
	rows, err := db.DbGetMailItemsByMailConfigId(confId, id, sqlDB)
	defer rows.Close()

	var results []models.OutputReq
	for rows.Next() {
		var output models.OutputReq
		err := rows.Scan(&output.MessageId, &output.RequestState,
			&output.ItemInsertTimestamp, &output.ItemUpdateTimestamp,
			&output.Receiver, &output.MailConfigId)
		if err != nil {
			panic(err)
		}
		results = append(results, output)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return results
}

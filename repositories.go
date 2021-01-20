package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// db instance
var db *pgx.Conn

// connect to db
func connectDb(s DataSource) *pgx.Conn {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.User, s.Password, s.Dbname)
	db, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Connected to %s as %s\n", s.Dbname, s.User)
	return db
}

// ping db
func pingDb(db *pgx.Conn) {
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func startStream(key string) (int, error) {
	sql := `INSERT INTO streams (title, user_id, live, start_time)
	VALUES (
		(SELECT username FROM users WHERE stream_key = $1) || '''s Stream ' || NOW(),
		(SELECT id FROM users WHERE stream_key = $1),
		true,
		NOW()
	) RETURNING id;`

	var id int
	err := db.QueryRow(context.Background(), sql, key).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func stopStream(key string) (int, error) {
	sql := `UPDATE streams
	SET live = false, end_time = NOW()
	WHERE (
		user_id = (SELECT id FROM users WHERE stream_key = $1)
		AND live = true
	) RETURNING id;`

	var id int
	err := db.QueryRow(context.Background(), sql, key).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func queryStream(userID int) (Stream, error) {
	sql := `SELECT * FROM streams
	WHERE user_id = $1
	ORDER BY start_time DESC
	LIMIT 1`

	var stream Stream
	err := db.QueryRow(context.Background(), sql, userID).Scan(
		&stream.ID,
		&stream.Title,
		&stream.UserID,
		&stream.Live,
		&stream.StartTime,
		&stream.EndTime,
	)
	if err != nil {
		return stream, err
	}
	return stream, nil
}

func queryStreams(userID int) ([]Stream, error) {
	sql := `SELECT * FROM streams
	WHERE user_id = $1
	ORDER BY start_time DESC`
	rows, err := db.Query(context.Background(), sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var streams []Stream
	for rows.Next() {
		var stream Stream
		err = rows.Scan(
			&stream.ID,
			&stream.Title,
			&stream.UserID,
			&stream.Live,
			&stream.StartTime,
			&stream.EndTime,
		)
		if err != nil {
			return streams, err
		}
		streams = append(streams, stream)
	}
	err = rows.Err()
	if err != nil {
		return streams, err
	}
	return streams, nil
}

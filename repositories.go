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

func startStream(key string) (Stream, error) {
	sql := `INSERT INTO streams (title, user_id, live, start_time)
	VALUES (
		(SELECT username FROM users WHERE stream_key = $1) || '''s Stream ' || NOW(),
		(SELECT id FROM users WHERE stream_key = $1),
		true,
		NOW()
	) RETURNING *;`

	var stream Stream
	err := db.QueryRow(context.Background(), sql, key).Scan(
		&stream.ID,
		&stream.Title,
		&stream.Live,
		&stream.StartTime,
		&stream.EndTime,
		&stream.UserID,
	)
	if err != nil {
		return stream, err
	}
	return stream, nil
}

func stopStream(key string) (Stream, error) {
	sql := `UPDATE streams
	SET live = false, end_time = NOW()
	WHERE (
		user_id = (SELECT id FROM users WHERE stream_key = $1)
		AND live = true
	) RETURNING *;`

	var stream Stream
	err := db.QueryRow(context.Background(), sql, key).Scan(
		&stream.ID,
		&stream.Title,
		&stream.Live,
		&stream.StartTime,
		&stream.EndTime,
		&stream.UserID,
	)
	if err != nil {
		return stream, err
	}
	return stream, nil
}

func deleteStreams(userID int) error {
	sql := `DELETE FROM streams
	WHERE user_id = $1;`

	_, err := db.Exec(context.Background(), sql, userID)
	if err != nil {
		return err
	}
	return nil
}

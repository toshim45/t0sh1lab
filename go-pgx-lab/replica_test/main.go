package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	QueryWorkSelect                = `SELECT id, status FROM stock.work WHERE id=$1`
	QueryWorkUpdateStatusToFailed  = `UPDATE stock.work SET status='failed' WHERE id=$1`
	QueryWorkUpdateStatusToSuccess = `UPDATE stock.work SET status='success' WHERE id=$1`
)

func main() {
	ctx := context.Background()

	var (
		err                     error
		masterConn, replicaConn *pgx.Conn
		delay                   time.Duration
		workID                  uuid.UUID
	)

	if workID, err = uuid.Parse(os.Getenv("WORK_ID")); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] Failed WORK_ID: %v\n", err)
		os.Exit(1)
	}

	if delay, err = time.ParseDuration(os.Getenv("DELAY")); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] Failed DELAY: %v\n", err)
		os.Exit(1)
	}

	if masterConn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL")); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] Failed master connection: %v\n", err)
		os.Exit(1)
	}

	defer masterConn.Close(ctx)

	if replicaConn, err = pgx.Connect(ctx, os.Getenv("REPLICA_DATABASE_URL")); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] Failed replica connection: %v\n", err)
		os.Exit(1)
	}

	defer replicaConn.Close(ctx)
	ResetWorkStatus(ctx, masterConn, workID)

	var startTime time.Time
	SelectWorkStatus(ctx, replicaConn, workID, startTime)
	startTime = time.Now()
	UpdateWorkStatusToFailed(ctx, masterConn, workID)
	for i := 0; i < 5; i++ {
		if SelectWorkStatus(ctx, replicaConn, workID, startTime) {
			break
		}
		time.Sleep(delay)
	}

	fmt.Printf("[PGX Test] DONE%v\n", time.Since(startTime))
}

func ResetWorkStatus(ctx context.Context, conn *pgx.Conn, id uuid.UUID) {
	if _, err := conn.Exec(ctx, QueryWorkUpdateStatusToSuccess, id); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] error update status: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[PGX Test][Master] reset status to success \n")
	time.Sleep(1 * time.Second)
}

func UpdateWorkStatusToFailed(ctx context.Context, conn *pgx.Conn, id uuid.UUID) {
	if _, err := conn.Exec(ctx, QueryWorkUpdateStatusToFailed, id); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] error update status: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[PGX Test][Master] status updated to failed \n")
}

func SelectWorkStatus(ctx context.Context, conn *pgx.Conn, id uuid.UUID, start time.Time) bool {
	var (
		workIDResult     uuid.UUID
		workStatusResult string
	)
	if err := conn.QueryRow(ctx, QueryWorkSelect, id).Scan(&workIDResult, &workStatusResult); err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test][Replica] Failed replica connection: %v\n", err)
		os.Exit(1)
	}

	var since string
	if !start.IsZero() {
		since = fmt.Sprintf("%v", time.Since(start))
	}

	fmt.Printf("[PGX Test][Replica] current status: %s %v\n", workStatusResult, since)
	return workStatusResult == "failed"
}

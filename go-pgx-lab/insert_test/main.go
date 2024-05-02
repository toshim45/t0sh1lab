package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	// TableName = "trial.items_name_btree"
	// TableName = "trial.items_name_gin"
	TableName = "trial.items_name_btree_gin"
)

var (
	MaxIteration, BatchSize int
	AttributeSelect         = `SELECT COUNT(id) FROM ` + TableName
	AttributeInsert         = `INSERT INTO ` + TableName + ` (id, number, name, created_at) VALUES ` //  (DEFAULT, DEFAULT, DEFAULT, DEFAULT)
)

type Attribute struct {
	ID     uuid.UUID
	Number int32
	Name   string
}

func NewAttribute(number int32) Attribute {
	id, _ := uuid.NewUUID()
	name := fmt.Sprintf("number-%d", number)
	return Attribute{
		ID:     id,
		Number: number,
		Name:   name,
	}
}

func main() {
	ctx := context.Background()

	maxItr, errc := strconv.Atoi(os.Getenv("MAX_ITERATION"))
	if errc != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] max-iteration-error: %v\n", maxItr)
		os.Exit(1)
	}
	batchSz, errc := strconv.Atoi(os.Getenv("BATCH_SIZE"))
	if errc != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] batch-size-error: %v\n", batchSz)
		os.Exit(1)
	}

	MaxIteration = maxItr
	BatchSize = batchSz

	fmt.Printf("[PGX Test] Init: maxItr %d, batchSz %d, table %s\n", MaxIteration, BatchSize, TableName)

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	defer conn.Close(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	var start time.Time

	n := Count(ctx, conn)
	start = time.Now()
	InsertOne(ctx, conn, n)
	fmt.Printf("[PGX Test] after-insert-one %v\n", time.Since(start))
	start = time.Now()
	Batch(ctx, conn, n+MaxIteration, MaxIteration, BatchSize, InsertArray)
	fmt.Printf("[PGX Test] after-insert-Array %v\n", time.Since(start))
	Count(ctx, conn)
}

func Count(ctx context.Context, conn *pgx.Conn) int {
	var count int
	if errq := conn.QueryRow(ctx, AttributeSelect).Scan(&count); errq != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] error query row: %v\n", errq)
		os.Exit(1)
	}

	fmt.Printf("[PGX Test] count: %d \n", count)
	return count
}

func InsertOne(ctx context.Context, conn *pgx.Conn, startNumber int) {
	for i := startNumber; i < MaxIteration+startNumber; i++ {
		m := NewAttribute(int32(i))
		sqlStr := AttributeInsert + `($1,$2,$3,$4)`
		if _, errq := conn.Exec(ctx, sqlStr, m.ID, m.Number, m.Name, time.Now()); errq != nil {
			fmt.Fprintf(os.Stderr, "[PGX Test] error insert-one: %v\n", errq)
			os.Exit(1)
		}
	}
}

func InsertArray(ctx context.Context, conn *pgx.Conn, startNumber, endNumber int) {
	j := 0
	paramStrList := []string{}
	valueList := []interface{}{}
	now := time.Now()
	for i := startNumber; i < endNumber; i++ {
		m := NewAttribute(int32(i))
		valueList = append(valueList, m.ID, m.Number, m.Name, now)
		paramStrList = append(paramStrList, fmt.Sprintf("($%d,$%d,$%d,$%d)", j+1, j+2, j+3, j+4))
		j += 4
	}

	sqlStr := AttributeInsert + strings.Join(paramStrList, ",")

	if _, errq := conn.Exec(ctx, sqlStr, valueList...); errq != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] error insert-array: %v\n", errq)
		fmt.Fprintf(os.Stderr, "[PGX Test] %s => %v\n", sqlStr, valueList)
		os.Exit(1)
	}
}

func Batch(ctx context.Context, conn *pgx.Conn, initial, maxSize, batchSize int, eachFn func(ctx context.Context, conn *pgx.Conn, start, end int)) {
	i := initial
	totalSize := initial + maxSize
	for i < totalSize {
		end := i + batchSize - 1
		if end > totalSize-1 {
			end = totalSize - 1
		}
		// fmt.Printf("[PGX Test] batch %d-%d\n", i, end)
		eachFn(ctx, conn, i, end)
		i = end + 1
	}
}

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
	AttributeSelect = `SELECT COUNT(id) FROM stock.x_attribute`
	AttributeInsert = `INSERT INTO stock.x_attribute (id, number, message, created_at, created_by) VALUES ` //  (DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT)
)

var (
	MaxIteration int
)

type Attribute struct {
	ID        uuid.UUID
	Number    int32
	Message   string
	CreatedAt time.Time
	CreatedBy uuid.UUID
}

func NewAttribute(number int32) Attribute {
	id, _ := uuid.NewUUID()
	message := fmt.Sprintf("number-%d", number)
	return Attribute{
		ID:        id,
		Number:    number,
		Message:   message,
		CreatedAt: time.Now(),
		CreatedBy: id,
	}
}

func main() {
	ctx := context.Background()

	maxItr, errc := strconv.Atoi(os.Getenv("MAX_ITERATION"))
	if errc != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] max-iteration-error: %v\n", maxItr)
		os.Exit(1)
	}

	MaxIteration = maxItr

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	defer conn.Close(context.Background())

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
	InsertArray(ctx, conn, n)
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
		sqlStr := AttributeInsert + `($1,$2,$3,$4,$5)`
		if _, errq := conn.Exec(ctx, sqlStr, m.ID, m.Number, m.Message, m.CreatedAt, m.CreatedBy); errq != nil {
			fmt.Fprintf(os.Stderr, "[PGX Test] error insert-one: %v\n", errq)
			os.Exit(1)
		}
	}
}

func InsertArray(ctx context.Context, conn *pgx.Conn, startNumber int) {
	j := 0
	paramStrList := []string{}
	valueList := []interface{}{}
	for i := startNumber; i < MaxIteration+startNumber; i++ {
		m := NewAttribute(int32(i))
		valueList = append(valueList, m.ID, m.Number, m.Message, m.CreatedAt, m.CreatedBy)
		paramStrList = append(paramStrList, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d)", j+1, j+2, j+3, j+4, j+5))
		j += 5
	}

	sqlStr := AttributeInsert + strings.Join(paramStrList, ",")

	if _, errq := conn.Exec(ctx, sqlStr, valueList...); errq != nil {
		fmt.Fprintf(os.Stderr, "[PGX Test] error insert-array: %v\n", errq)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sample-golang/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// データベース接続を設定
	config, err := pgxpool.ParseConfig(getDBURL())
	if err != nil {
		log.Fatalf("プール設定の解析に失敗しました: %v", err)
	}

	// 接続プールの初期化
	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("データベースに接続できません: %v", err)
	}
	defer conn.Close()

	queries := sqlc.New(conn) // 生成されたクエリインターフェースを初期化
	fmt.Println("データベースに正常に接続しました")

	// ルーティング設定
	r := setupRouter(queries)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func getDBURL() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
}



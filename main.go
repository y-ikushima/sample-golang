package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sample-golang/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	// データベース接続を設定
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, getDBURL())
	if err != nil {
		log.Fatalf("データベースに接続できません: %v", err)
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn) // 生成されたクエリインターフェースを初期化
	fmt.Println("データベースに正常に接続しました")

	// Ginを設定
	r := gin.Default()

	// ヘルスチェック用の簡単なエンドポイントを定義
	r.GET("/health", func(c *gin.Context) {
		// 単純なレスポンスとしてステータス200を返す
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	// データベースからデータを取得するエンドポイントを定義
	r.GET("/user", func(c *gin.Context) {
		idStr := c.Query("id") // クエリパラメータからIDを取得
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}
	
		id, err := strconv.Atoi(idStr) // stringからintに変換
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		user, err := getUsers(queries, int32(id)) // IDを引数として渡す
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーを取得できません"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

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

func getUsers(q *sqlc.Queries, id int32) (*sqlc.GetUserRow, error) {
	user, err := q.GetUser(context.Background(),id) // 生成されたクエリを呼び出す
	if err != nil {
		return nil, err
	}
	return &user, nil
}
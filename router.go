// router.go
package main

import (
	"net/http"
	"sample-golang/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)
func setupRouter(queries *sqlc.Queries) *gin.Engine {
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

		// ユーザーデータを取得する
		user, queryErr := getUsers(queries, int32(id)) // ゴルーチンを使用せずに呼び出し

		if queryErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーを取得できません"})
			return
		}

		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	

	return r
}

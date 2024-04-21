package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"trpc-query-demo/proto/query"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

// 定义全局Redis和MySQL
var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",
		DB:       0,
	})
	dsn = "root:@tcp(129.204.44.167:3306)/messageboard" // 数据库DSN
)

// 定义了QueryServiceImpl类型来实现protobuf文件中定义的服务接口
type QueryServiceImpl struct {
	query.UnimplementedQueryService
}

// NewQueryService函数用于创建服务的实例
func NewQueryService() *QueryServiceImpl {
	return &QueryServiceImpl{}
}

// NullStringToString将sql.NullString转化为普通string，处理了SQL返回的NULL值
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func (service *QueryServiceImpl) Query(ctx context.Context, req *query.QueryRequest) (*query.QueryResponse, error) {
	queryText := req.GetQueryText()
	if queryText == "" {
		return nil, fmt.Errorf("queryText parameter is missing")
	}

	// 尝试从Redis获取缓存结果
	cachedResult, err := rdb.Get(ctx, queryText).Result()
	if err == nil {
		var response query.QueryResponse
		if err := json.Unmarshal([]byte(cachedResult), &response); err == nil {
			return &response, nil
		}
		log.Printf("Failed to unmarshal cached result: %v", err)
		// 如果缓存解析失败，则继续查询数据库
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	startTime := time.Now()

	sqlQuery := "SELECT ID, AskText, TitleText, ReplyText FROM `messegetext` WHERE AskText LIKE ?"
	likeText := "%" + queryText + "%"
	rows, err := db.Query(sqlQuery, likeText)
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

	var rowsData []*query.RowData
	count := 0

	for rows.Next() {
		var ID int32
		var AskText sql.NullString
		var TitleText sql.NullString
		var ReplyText sql.NullString // 使用sql.NullString
		if err := rows.Scan(&ID, &AskText, &TitleText, &ReplyText); err != nil {
			log.Println("Failed to scan row:", err)
			continue
		}

		row := &query.RowData{
			Id:        ID,
			AskText:   NullStringToString(AskText),
			TitleText: NullStringToString(TitleText),
			ReplyText: NullStringToString(ReplyText),
		}

		rowsData = append(rowsData, row)
		count++
	}

	duration := time.Since(startTime).Milliseconds()

	response := &query.QueryResponse{
		Duration: fmt.Sprintf("%dms", duration),
		Count:    int32(count),
		Rows:     rowsData,
	}

	// 将结果序列化后存入Redis缓存
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
	} else {
		if err := rdb.Set(ctx, queryText, string(responseJSON), 10*time.Minute).Err(); err != nil {
			log.Printf("Failed to set cache: %v", err)
		}
	}

	return response, nil
}

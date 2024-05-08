package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"trpc-query-demo/proto/query"
    "github.com/golang/groupcache/lru"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

// 定义全局Redis和MySQL
var (
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis服务器地址
		Password: "",
		DB:       0,
	})
	dsn = "root:@tcp(129.204.44.167:3306)/messageboard" // 数据库DSN
    cache = lru.New(128) 
)

// 定义了QueryServiceImpl类型来实现protobuf文件中定义的服务接口
type QueryServiceImpl struct {
	query.UnimplementedQueryService
}

// NewQueryService函数用于创建服务的实例
func NewQueryService() *QueryServiceImpl {
	service := &QueryServiceImpl{}
	return service
}

// NullStringToString将sql.NullString转化为普通string，处理了SQL返回的NULL值
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func (service *QueryServiceImpl) Query(ctx context.Context, req *query.QueryRequest) (*query.QueryResponse, error) {
	// 确保service.cache已经初始化
	if cache == nil {
		return nil, fmt.Errorf("internal error: LRU cache is not initialized")
	}
	queryText := req.GetQueryText()
	if queryText == "" {
		return nil, fmt.Errorf("queryText parameter is missing")
	}
	page := req.GetPage() // 获取请求的页码
	if page <= 0 {
		page = 1 // 默认值为第1页
	}
	perPage := req.GetPerPage() // 获取每页的结果数量
	if perPage <= 0 {
		perPage = 100 // 默认值为每页100条
	}

	// 计算查询偏移量
	offset := (page - 1) * perPage

	// 生成包含查询文本、页码和每页条数的唯一缓存键
	cachedKey := fmt.Sprintf("%s_%d_%d", queryText, page, perPage)

	// 首先尝试从内存缓存中获取
	if val, ok := cache.Get(cachedKey); ok {
		// 确认从LRU缓存中取出的数据确实是[]byte类型
		cachedBytes, isByteSlice := val.([]byte)
		if !isByteSlice {
			log.Println("Cache value is not of type []byte")

			return nil, fmt.Errorf("cache value is of incorrect type")
		}

		var response query.QueryResponse
		if err := json.Unmarshal(cachedBytes, &response); err == nil {
			return &response, nil
		} else {
			log.Println("Failed to unmarshal from LRU cache:", err)
			cache.Remove(cachedKey)
		}
	}

	// 尝试从Redis获取缓存结果
	cachedResult, err := rdb.Get(ctx, cachedKey).Result()
	if err == nil {
		var response query.QueryResponse
		if err := json.Unmarshal([]byte(cachedResult), &response); err == nil {
			// 将结果同时存储在内存缓存中
			cache.Add(cachedKey, []byte(cachedResult))
			return &response, nil
		}
		log.Printf("Failed to unmarshal cached result: %v", err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	startTime := time.Now()

	sqlQuery := "SELECT ID, AskText, TitleText, ReplyText FROM `messegetext` WHERE AskText LIKE ? LIMIT ? OFFSET ?"
	likeText := "%" + queryText + "%"
	rows, err := db.Query(sqlQuery, likeText, perPage, offset)
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
		// 创建一个新的context，用于Redis操作，设置更长的超时时间
		cacheCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel() // 确保新创建的context被正确释放

		// 使用新的context来设置缓存
		if err := rdb.Set(cacheCtx, cachedKey, string(responseJSON), 10*time.Minute).Err(); err != nil {
			log.Printf("Failed to set cache: %v", err)
		} else {
			// Redis缓存设置成功后，更新到内存(LRU)缓存
			cache.Add(cachedKey, responseJSON) // 注意此处直接存储序列化后的[]byte类型
			log.Println("Cache updated in both Redis and LRU")
		}
	}

	return response, nil
}

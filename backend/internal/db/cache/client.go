package cache

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	EmployeeId   string    `json:"employeeId"`
	TenantId     string    `json:"tenantId"`
	DepartmentId string    `json:"departmentId"`
	IP           string    `json:"ip"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

var CacheClient *redis.Client

func NewCacheClient() (*redis.Client, error) {
	ctx := context.Background()
	DB, err := strconv.Atoi(os.Getenv("DB_REDIS_DB"))
	if err != nil {
		DB = 0
	}
	CacheClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_REDIS_ADDR"),
		Password: os.Getenv("DB_REDIS_PASSWORD"),
		DB:       DB,

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 4,
	})

	if _, err := CacheClient.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("error conecting to redis")
	}
	return CacheClient, nil
}

func GenerateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func SaveSession(ctx context.Context, sessionID string, sess Session) error {
	if CacheClient == nil {
		return fmt.Errorf("could not connect to cache client")
	}
	data, _ := json.Marshal(sess)
	log.Print(data)
	return CacheClient.Set(ctx, "sess:"+sessionID, data, time.Until(sess.ExpiresAt)).Err()
}

func SaveQuery(ctx context.Context, key string, response interface{}) error {
	if CacheClient == nil {
		return fmt.Errorf("could not connect to cache client")
	}
	data, _ := json.Marshal(response)
	return CacheClient.Set(ctx, key, data, time.Until(time.Now().Add(1*time.Minute))).Err()
}

func GetSession(ctx context.Context, sessionID string) (*Session, error) {
	if CacheClient == nil {
		return nil, fmt.Errorf("could not connect to cache client")
	}
	data, err := CacheClient.Get(ctx, "sess:"+sessionID).Bytes()
	if err != nil {
		return nil, err
	}
	var sess Session
	json.Unmarshal(data, &sess)
	return &sess, nil
}

func GetQuery[T any](ctx context.Context, key string) (*T, error) {
	data, err := CacheClient.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteQuery(ctx context.Context, key string) error {
	if CacheClient == nil {
		return fmt.Errorf("could not connect to cache client")
	}
	return CacheClient.Del(ctx, key).Err()
}

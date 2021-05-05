package utils

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// ConnectToCache
// connects to single redis instance
// Takes context, host, password and db
// Returns client, error
func ConnectToCache(ctx context.Context, host, password string, db int) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
		PoolSize: 100,
	})
	if err = client.Ping(ctx).Err(); err != nil {
		log.Printf("[r] unable to ping redis because %v", err)
		return nil, err
	}
	log.Printf("[r] connected to redis successfully")
	return client, nil
}

// SaveToCache
// Saves data to redis
// Takes context, key, data, expiry, client
// Returns error
func SaveToCache(ctx context.Context, key string, data interface{}, expiry time.Duration, client *redis.Client) (err error) {
	log.Printf("[r] saving data to redis under key %s", key)
	if err = client.Set(ctx, key, data, expiry).Err(); err != nil {
		log.Printf("[r] unable to save data to redis under key %s because %v", key, err)
		return err
	}
	log.Printf("[r] saved data to redis under key %v", err)
	return nil
}

// ReadFromCache
// Reads data from redis
// Takes context, key, client
// Returns data, error
func ReadFromCache(ctx context.Context, key string, client *redis.Client) (data interface{}, err error) {
	log.Printf("[r] reading data from redis under key %s", key)
	dataStr, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("[r] unable to read data from redis %s because %v", key, err)
		return
	}
	if err = json.Unmarshal([]byte(dataStr), &data); err != nil {
		log.Printf("[r] unable to unmarshal data in %s because %v", key, err)
		return
	}
	log.Printf("[r] got data from redis key %s", key)
	return data, nil
}

// DeleteFromCache
// Deletes a key from redis
// Takes context, key, client
// Returns error
func DeleteFromCache(ctx context.Context, key string, client *redis.Client) (err error) {
	log.Printf("[r] Deleting redis key %s", key)
	if err = client.Del(ctx, key).Err(); err != nil {
		log.Printf("[r] unable to delete redis %s", key)
		return
	}
	log.Printf("[r] deleted redis key %s successfully", key)
	return nil
}

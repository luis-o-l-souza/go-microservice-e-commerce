package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/luis-o-l-souza/go-microservice-e-commerce/internal/cart"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) GenerateKey(userId int) string {
	return fmt.Sprintf("cart:%d", userId)
}

func (r *RedisRepository) Get(ctx context.Context, userId int) (*cart.Cart, error) {
	key := r.GenerateKey(userId)

	val, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return &cart.Cart{UserID: userId, Items: []cart.CartItem{}}, nil
	} else if err != nil {
		return nil, err
	}

	var c cart.Cart

	err = json.Unmarshal([]byte(val), &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *RedisRepository) Save(ctx context.Context, c *cart.Cart) error {
	key := r.GenerateKey(c.UserID)

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, 24*time.Hour).Err()
}

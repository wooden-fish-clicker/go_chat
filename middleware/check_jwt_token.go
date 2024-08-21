package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/wooden-fish-clicker/chat/pkg/redis"
	"github.com/wooden-fish-clicker/chat/pkg/utils"
)

func ValidateWebSocketToken(r *http.Request) (*utils.Claims, error) {

	token := r.URL.Query().Get("token")
	if token == "" {
		return nil, errors.New("未提供 token")
	}

	// 檢查 redis 裡面的黑名單 token
	err := checkRedisJwtBlackList(token)
	if err != nil && err.Error() != "redis: nil" {
		return nil, errors.New("發生錯誤")
	} else if err == nil {
		return nil, errors.New("此 token 已經失效")
	}

	// 驗證 JWT
	claims, err := utils.ParseJwtToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func checkRedisJwtBlackList(token string) error {
	_, err := redis.Rd.Get(context.Background(), "jwt:blacklist:"+token).Result()

	if err != nil {
		return err
	}

	return nil
}

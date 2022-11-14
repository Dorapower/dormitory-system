package cache

import (
	"dormitory-system/src/database"
	"strconv"
	"time"
)

const CatchExpireTime = time.Hour * 24 * 7

func SetRefreshTokenCache(rToken string, uId int) error {
	var redisDb = database.RedisDb

	// e.g.  refresh_token_1
	key := "refresh_token_" + strconv.Itoa(uId)

	err := redisDb.Set(key, rToken, CatchExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRefreshTokenCache(uId int) string {
	var redisDb = database.RedisDb
	rToken, err := redisDb.Get("refresh_token_" + strconv.Itoa(uId)).Result()
	if err != nil {
		return ""
	}
	return rToken
}

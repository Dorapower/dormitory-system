package cache

import (
	"dormitory-system/src/database"
	"strconv"
	"time"
)

const TokenCacheExpireTime = time.Hour * 24 * 7
const RoomCacheExpireTime = time.Hour * 1

func SetRefreshTokenCache(rToken string, uId int) error {
	var redisDb = database.RedisDb

	// e.g.  refresh_token_1
	key := "refresh_token_" + strconv.Itoa(uId)

	err := redisDb.Set(key, rToken, TokenCacheExpireTime).Err()
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

func DeleteRefreshTokenCache(uId int) error {
	var redisDb = database.RedisDb
	err := redisDb.Del("refresh_token_" + strconv.Itoa(uId)).Err()
	if err != nil {
		return err
	}
	return nil
}

// SetBuildingCache Set building's rest beds on cache
func SetBuildingCache(bId, gender, count int) error {
	var redisDb = database.RedisDb

	key := "building_" + strconv.Itoa(bId) + "_" + strconv.Itoa(gender)

	err := redisDb.Set(key, count, RoomCacheExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetBuildingCache Get building's rest beds from cache
func GetBuildingCache(bId, gender int) int {
	var redisDb = database.RedisDb

	key := "building_" + strconv.Itoa(bId) + "_" + strconv.Itoa(gender)

	value, err := redisDb.Get(key).Result()
	if err != nil {
		return -1
	}
	count, _ := strconv.Atoi(value)
	return count
}

// SetRoomCache  Set room's rest beds on cache
func SetRoomCache(rId, count int) error {
	var redisDb = database.RedisDb

	key := "room_" + strconv.Itoa(rId)

	err := redisDb.Set(key, count, RoomCacheExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetRoomCache  Get room's rest beds from cache
func GetRoomCache(rId int) int {
	var redisDb = database.RedisDb

	key := "room_" + strconv.Itoa(rId)

	value, err := redisDb.Get(key).Result()
	if err != nil {
		return -1
	}
	count, _ := strconv.Atoi(value)
	return count
}

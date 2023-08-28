package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"tiktok/pkg/constant"
	"tiktok/pkg/utils"
)

type logRedisData struct {
	Cmd []*logRedisCmdData `json:"cmd"`
}

type logRedisCmdData struct {
	Cmd string
	Err error
}

// 返回是否有错误并打印日志
func handlerErrs(cmds []redis.Cmder, ctx *gin.Context) bool {
	data := logRedisData{}
	for _, cmd := range cmds {
		if cmd.Err() != nil {
			data.Cmd = append(data.Cmd, &logRedisCmdData{
				Cmd: cmd.String(),
				Err: cmd.Err(),
			})
		}
	}
	if len(data.Cmd) == 0 {
		return false
	}

	utils.LogWithRID(constant.LMRedis, &data, ctx).Debug("multi cmd err")
	return true
}

// 返回是否有错误并打印日志
func handlerErr(cmd redis.Cmder, ctx *gin.Context) bool {
	if cmd.Err() != nil {
		utils.LogWithRID(constant.LMRedis, &logRedisData{
			Cmd: []*logRedisCmdData{
				{
					Cmd: cmd.String(),
					Err: cmd.Err(),
				},
			},
		}, ctx).WithError(cmd.Err()).Debug("cmd err")
		return true
	}
	return false
}

// 返回是否有错误并打印日志
func cmdsString(cmds []redis.Cmder) string {
	sb := strings.Builder{}
	size := len(cmds)
	for i, cmd := range cmds {
		sb.WriteString(strconv.FormatInt(int64(i+1), 10))
		sb.WriteString(": ")
		sb.WriteString(cmd.String())
		if i+1 != size {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

// LoadIfNotExists 将指定key加载到redis中
func LoadIfNotExists(key int64, rdb *redis.Client, LoadFunc func(key int64, rdb *redis.Client) error) error {
	exists, err := rdb.Exists(Ctx, strconv.FormatInt(key, 10)).Result()
	if err != nil {
		return err
	}
	//缓存不存在
	if exists == 0 {
		ok, err := FavMutex.Put(key)
		if err != nil {
			return err
		}
		if ok {
			err = LoadFunc(key, rdb)
			FavMutex.Del(key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

package kafka

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"tiktok/biz/dao"
	"tiktok/pkg/errno"
)

var FavoriteMq = GetKafka("Favorite", "AddFavorite")

// ConsumeFavorite 消费逻辑，在协程中异步写入数据库
func ConsumeFavorite(tKafka *TKafka) {
	msg, err := tKafka.ReadMsg()
	if err != nil {
		logrus.WithField("kafka", msg).Warn("kafka读取失败")
	}
	UserId, _ := strconv.ParseInt(string(msg.Key), 10, 64)
	value := strings.Split(string(msg.Value), " ")
	VideoId, _ := strconv.ParseInt(string(value[0]), 10, 64)
	if len(value) == 1 {
		err = dao.DelFavorite(UserId, VideoId)
	} else {
		createAt, _ := strconv.ParseInt(string(value[1]), 10, 64)
		err = dao.AddFavorite(UserId, VideoId, createAt)
	}
	if err != nil {
		if err == errno.FavoriteRelationAlreadyExistErr {
			logrus.WithFields(logrus.Fields{
				"UserId":  UserId,
				"VideoId": VideoId,
				"Action":  len(value),
			}).Warn("重复点赞或取消点赞")
		}
		logrus.WithFields(logrus.Fields{
			"UserId":  UserId,
			"VideoId": VideoId,
			"Action":  len(value),
		}).Error("数据写入失败")
	}

}

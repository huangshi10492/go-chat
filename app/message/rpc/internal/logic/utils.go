package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func getSeq(ctx context.Context, rds *redis.Redis, userId int64) (int64, error) {
	//field := fmt.Sprintf("%d_seq", userId)
	//res, err := rds.HincrbyCtx(ctx, "user_message_seq", field, 1)
	//if err != nil {
	//	return 0, err
	//}
	//return int64(res), nil
	lock := redis.NewRedisLock(rds, "user_message_seq")
	for i := 0; i < 10; i++ {
		acquire, err := lock.Acquire()
		switch {
		case err != nil:
			return 0, nil
		case acquire:
			// 获取到锁
			defer lock.Release() // 释放锁
			// 业务逻辑
			field := fmt.Sprintf("%d_seq", userId)
			res, err := rds.HincrbyCtx(ctx, "user_message_seq", field, 1)
			if err != nil {
				return 0, err
			}
			return int64(res), nil
		}
	}
	return 0, errors.New("获取锁失败")
}

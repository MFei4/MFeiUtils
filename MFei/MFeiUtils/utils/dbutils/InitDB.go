package dbutils

import (
	"github.com/mfei/mfeiutils/common/MFei"
	"github.com/mfei/mfeiutils/common/sun"
	"github.com/mfei/mfeiutils/utils/dbutils/redisUtil"
)

// GetDB 开放给外部获得db连接
func GetDB(typ string) {
	MFei.DB = setup(typ)
	sun.Redis = redisUtil.Setup()
}

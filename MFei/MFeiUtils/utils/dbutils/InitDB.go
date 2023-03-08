package dbutils

import (
	"MFeiUtils/common/MFei"
	"MFeiUtils/common/sun"
	"MFeiUtils/utils/dbutils/redisUtil"
)

// GetDB 开放给外部获得db连接
func GetDB(typ string) {
	MFei.DB = setup(typ)
	sun.Redis = redisUtil.Setup()
}

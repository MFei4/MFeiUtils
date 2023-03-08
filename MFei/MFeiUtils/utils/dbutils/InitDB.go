package dbutils

import (
	"github.com/MFei/MFeiUtils/common/MFei"
	"github.com/MFei/MFeiUtils/common/sun"
	"github.com/MFei/MFeiUtils/utils/dbutils/redisUtil"
)

// GetDB 开放给外部获得db连接
func GetDB(typ string) {
	MFei.DB = setup(typ)
	sun.Redis = redisUtil.Setup()
}

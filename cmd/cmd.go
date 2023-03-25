package cmd

import (
	"gokyrie/conf"
	"gokyrie/global"
	"gokyrie/router"
	"gokyrie/utils"
)

func Start() {
	var initErr error
	//初始化配置
	conf.InitConfig()
	//初始化日志
	global.Logger = conf.InitLogger()
	//初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	//初始化redis
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}
	//初始化路由
	router.InitRouter()
	//db.AutoMigrate(&model.Dept{}, &model.User{}, &model.Role{}, &model.Menu{})
}

func Clean() {}

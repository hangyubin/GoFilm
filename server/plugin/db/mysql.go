package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/config"
)

var Mdb *gorm.DB
var cfg = config.LoadConfig()

func InitDatabase() (err error) {
	// client 相关属性设置
	var dialector gorm.Dialector

	// 使用PostgreSQL驱动
	dialector = postgres.New(postgres.Config{
		DSN: cfg.Database.GetDSN(),
	})

	Mdb, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",                          //设置创建表时的前缀
			SingularTable: true, //是否使用 结构体名称作为表名 (关闭自动变复数)
			//NameReplacer:  strings.NewReplacer("spider_", ""), // 替表名和字段中的 Me 为 空
		},
		Logger: logger.Default.LogMode(logger.Info), //设置日志级别为Info
	})
	
	if err != nil {
		return err
	}
	
	// 获取底层的sql.DB对象，设置连接池参数
	sqlDB, err := Mdb.DB()
	if err != nil {
		return err
	}
	
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(3600)      // 设置连接可复用的最大时间（秒）
	sqlDB.SetConnMaxIdleTime(1800)      // 设置连接在空闲状态下的最大生存时间（秒）
	
	return nil
}

package db

import (
	"gorm.io/driver/mysql"
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

	// 根据配置的数据库类型选择不同的驱动
	switch cfg.Database.Type {
	case "postgres":
		dialector = postgres.New(postgres.Config{
			DSN: cfg.Database.GetDSN(),
		})
	default: // 默认使用mysql
		dialector = mysql.New(mysql.Config{
			DSN:                       cfg.Database.GetDSN(),
			DefaultStringSize:         255,   //string类型字段默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式
			DontSupportRenameColumn:   true,  // 用change 重命名列
			SkipInitializeWithVersion: false, // 根据当前Mysql版本自动配置
		})
	}

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

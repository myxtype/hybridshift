package conf

import "time"

// 全局配置
type GbeConfig struct {
	JwtSecret    jwtSecretConfig
	Logger       loggerConfig
	RestServer   serverConfig
	AppletServer serverConfig
	AdminServer  serverConfig
	SocketServer serverConfig
	DataSource   dataSourceConfig
	Redis        redisConfig
}

// JWT
type jwtSecretConfig struct {
	Admin string
	Rest  string
}

// 日志配置
type loggerConfig struct {
	Level    string // 打印日志的等级 debug/info/warn/error/panic/fatal
	Target   string // 日志打印目标："file" or "console"
	Filename string // 日志文件
}

// 接口服务配置
type serverConfig struct {
	Addr string // 监听地址
}

// 数据库配置
type dataSourceConfig struct {
	DSN         string        // 数据库配置
	MaxIdle     int           // 最大空闲数
	MaxOpen     int           // 最大连接数
	MaxLifetime time.Duration // 复用的最大时间：秒
	Migrate     bool          // 是否执行迁移
	LogDisabled bool          // 是否禁用SQL日志
}

// Redis配置
type redisConfig struct {
	Addr     string // 地址
	Password string // 密码
	DB       int    // 数据库
}

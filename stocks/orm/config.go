package orm

// DbSettings 数据库连接字符串属性
type DbSettings struct {
	DbName          string
	Host            string
	User            string
	Password        string
	Port            int
	Key             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // 分钟
}

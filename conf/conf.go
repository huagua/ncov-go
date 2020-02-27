package conf

import (
	"github.com/joho/godotenv"
	"ncov_go/model"
	"os"
)

func Init() {
	//加载.env文件
	godotenv.Load()

	//连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}

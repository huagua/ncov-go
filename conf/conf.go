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

	//判断该用户这天是否已经提交过
	count := 0
	if model.DB.Model(&model.Corp{}).Where("corpid = 100000001").Count(&count); count == 0 {
		model.DB.Create(&model.Corp{
			Corpid:       "100000001",
			TemplateCode: "北京大学",
		})
	}
}

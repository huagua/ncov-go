package model

//数据库迁移
func migration() {
	//自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserInfo{})
	DB.AutoMigrate(&DailyInfo{})
	DB.AutoMigrate(&Corp{})
}

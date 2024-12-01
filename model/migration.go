package model

func migration() {
	DB.Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(&User{}, &Task{})
}

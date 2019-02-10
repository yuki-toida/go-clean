package main

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}

func main() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.LogMode(true)
	// db.DropTableIfExists(&User{}, &Profile{}, &Email{})
	db.AutoMigrate(&User{}, &Profile{}, &Email{})

	create(db)

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	// SELECT * FROM `profiles`  WHERE `profiles`.`deleted_at` IS NULL AND ((`id` IN ('1','2')))
	// SELECT * FROM `emails`  WHERE `emails`.`deleted_at` IS NULL AND ((`user_id` IN ('1','2')))
	var users []User
	db.Preload("Profile").Preload("Emails").Find(&users)

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = '1' ORDER BY `users`.`id` ASC LIMIT 1
	// SELECT * FROM `emails`  WHERE `emails`.`deleted_at` IS NULL AND ((`user_id` = '1'))
	var emails []Email
	user := User{Model: gorm.Model{ID: 1}}
	db.First(&user)
	db.Model(&user).Related(&emails)

	// SELECT * FROM `profiles`  WHERE `profiles`.`deleted_at` IS NULL AND ((`profiles`.`name` = '2')) ORDER BY `profiles`.`id` ASC LIMIT 1
	var profile Profile
	db.Where(&Profile{Name: "2"}).First(&profile)

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN ('1','2')))
	var inUsers []User
	db.Where([]uint{1, 2}).Find(&inUsers)

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` NOT IN ('1','3')))
	var notInUsers []User
	db.Not([]uint{2, 3}).Find(&notInUsers)

	//UPDATE `profiles` SET `created_at` = '2019-02-07 19:11:03', `updated_at` = '2019-02-07 19:11:03', `deleted_at` = NULL, `name` = '22'  WHERE `profiles`.`deleted_at` IS NULL AND `profiles`.`id` = '2'
	profile.Name = "22"
	db.Save(&profile)

	// UPDATE `profiles` SET `deleted_at`='2019-02-07 19:31:54'  WHERE `profiles`.`deleted_at` IS NULL AND `profiles`.`id` = '2'
	db.Delete(&profile)
}

func create(db *gorm.DB) {
	for i := 1; i < 3; i++ {
		p := Profile{Name: strconv.Itoa(i)}
		db.Create(&p)
		u := User{ProfileID: p.ID}
		db.Create(&u)
		for j := 1; j < 4; j++ {
			email := strconv.Itoa(j) + "@hacobu.com"
			db.Create(&Email{Email: email, UserID: u.ID})
		}
	}

	// SELECT * FROM `profiles`  WHERE `profiles`.`deleted_at` IS NULL AND ((`profiles`.`name` = '3')) ORDER BY `profiles`.`id` ASC LIMIT 1
	// INSERT INTO `profiles` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ('2019-02-07 15:54:48','2019-02-07 15:54:48',NULL,'3')
	var profile Profile
	db.FirstOrCreate(&profile, Profile{Name: "3"})

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`profile_id` = '3')) ORDER BY `users`.`id` ASC LIMIT 1
	// INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`profile_id`) VALUES ('2019-02-07 15:54:48','2019-02-07 15:54:48',NULL,'3')
	var user User
	db.FirstOrCreate(&user, User{ProfileID: profile.ID})
}

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

type User struct {
	gorm.Model
	Profile   Profile
	ProfileID uint
	Emails    []Email
}

type Profile struct {
	gorm.Model
	Name string
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint `gorm:"index"`
}

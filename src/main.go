package main

import (
	"strconv"

	"github.com/yuki-toida/go-clean/src/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/go-clean/src/adapter/repositories"
	"github.com/yuki-toida/go-clean/src/external/config"
	"github.com/yuki-toida/go-clean/src/external/mysql"
	"github.com/yuki-toida/go-clean/src/external/web"
)

func main() {
	c := config.Load()
	db := mysql.Connect(c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
	defer mysql.Close()

	db.LogMode(true)
	db.DropTableIfExists(&repositories.User{}, &repositories.Profile{}, &repositories.Email{})
	db.AutoMigrate(&repositories.User{}, &repositories.Profile{}, &repositories.Email{})

	for i := 1; i < 3; i++ {
		p := repositories.Profile{Name: strconv.Itoa(i)}
		db.Create(&p)
		u := repositories.User{ProfileID: p.ID}
		db.Create(&u)
		for j := 1; j < 4; j++ {
			e := strconv.Itoa(j) + "@hacobu.jp"
			db.Create(&repositories.Email{Email: e, UserID: u.ID})
		}
	}

	repository := registry.NewRepository(db)
	usecase := registry.NewUseCase(repository)

	router := web.NewRouter(usecase)
	router.Run(":8088")
}

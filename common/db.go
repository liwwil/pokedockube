package common

import (
	"fmt"

	"github.com/TanyaEIEI/pokedex/graph/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDb() (*gorm.DB, error) {
	dsn := "postgres://admin:password123@postgres:5432/pokedoc?sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("open db complete")
	db.AutoMigrate(&model.Pokemon{}, &model.PokemonType{}, &model.PokemonCategory{}, &model.PokemonAbility{})
	fmt.Println("connected DB")

	return db, nil
}

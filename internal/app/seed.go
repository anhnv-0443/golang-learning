package app

import (
	"context"
	"go-app/config"
	"go-app/internal/domain"
	"go-app/pkg/errors"
	"go-app/pkg/postgres"

	"github.com/spf13/viper"

	roleRepository "go-app/internal/modules/role/repository"
	userRepository "go-app/internal/modules/user/repository"
	languageCodeRepository "go-app/internal/modules/language_code/repository"
	positionRepository "go-app/internal/modules/position/repository"
)

var pathJSON = "cmd/seed/data.json"

type seedData struct {
	Roles []domain.Role `json:"roles"`
	Users []domain.User `json:"users"`
	Language_codes []domain.LanguageCode `json:"language_codes"`
	Positions []domain.Position `json:"positions"`
}

// Seed is function that seed data
func Seed(dbConfig config.Database) error {
	db, err := postgres.NewGormDB(dbConfig)
	if err != nil {
		return errors.Wrap(err)
	}

	userRepo := userRepository.NewRepository(db)
	roleRepo := roleRepository.NewRepository(db)
	languageCodeRepo := languageCodeRepository.NewRepository(db)
	positionRepo := positionRepository.NewRepository(db)

	viper.SetConfigFile(pathJSON)
	err = viper.ReadInConfig()
	if err != nil {
		return errors.Wrap(err)
	}

	data := seedData{}
	err = viper.Unmarshal(&data)
	if err != nil {
		return errors.Wrap(err)
	}

	for _, r := range data.Roles {
		err = roleRepo.Store(context.Background(), &r)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	for _, u := range data.Users {
		err = userRepo.Store(context.Background(), &u)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	for _, u := range data.Language_codes {
		err = languageCodeRepo.Store(context.Background(), &u)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	for _, u := range data.Positions {
		err = positionRepo.Store(context.Background(), &u)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

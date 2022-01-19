package repository

import "todos/src/profile/model"

type ProfileRepository interface {
	Save(*model.Profile) error
}

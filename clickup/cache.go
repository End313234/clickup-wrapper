package clickup

import "errors"

type UserCache []User
type SpaceCache []Space

func (uc UserCache) Find(callback func(User) bool) (User, error) {
	for _, user := range uc {
		if callback(user) {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (sc SpaceCache) Find(callback func(Space) bool) (Space, error) {
	for _, space := range sc {
		if callback(space) {
			return space, nil
		}
	}
	return Space{}, errors.New("space not found")
}

func (sc SpaceCache) Add(space Space) SpaceCache {
	sc = append(sc, space)
	return sc
}

type Cache struct {
	Users  UserCache
	Spaces SpaceCache
}

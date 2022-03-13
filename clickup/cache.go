package clickup

import (
	"errors"
)

type UserCache []User
type SpaceCache []Space
type WebhookCache []Webhook
type TeamCache []Team

// Iteratively finds an User in the cache by a condition
func (uc *UserCache) Find(callback func(User) bool) (User, error) {
	for _, user := range *uc {
		if callback(user) {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

// Iteratively finds a Space in the cache by a condition
func (sc *SpaceCache) Find(callback func(Space) bool) (Space, error) {
	for _, space := range *sc {
		if callback(space) {
			return space, nil
		}
	}
	return Space{}, errors.New("space not found")
}

// Adds a Space to the cache
func (sc *SpaceCache) Add(space Space) *SpaceCache {
	*sc = append((*sc), space)
	return sc
}

// Removes a Space from the cache
func (sc *SpaceCache) Remove(index int) *SpaceCache {
	(*sc)[len(*sc)-1], (*sc)[index] = (*sc)[index], (*sc)[len(*sc)-1]
	(*sc) = (*sc)[:len(*sc)-1]
	return sc
}

// Updates a Space in the cache
func (sc *SpaceCache) Update(space Space) *SpaceCache {
	for index, element := range *sc {
		if element.Id == space.Id {
			(*sc)[index] = space
			break
		}
	}

	return sc
}

// Iteratively finds a Webhook in the cache by a condition
func (wh *WebhookCache) Find(callback func(Webhook) bool) (Webhook, error) {
	for _, webhook := range *wh {
		if callback(webhook) {
			return webhook, nil
		}
	}
	return Webhook{}, errors.New("webhook not found")
}

// Adds a Webhook to the cache
func (wh *WebhookCache) Add(webhook Webhook) *WebhookCache {
	*wh = append((*wh), webhook)
	return wh
}

// Removes a Webhook from the cache
func (wh *WebhookCache) Remove(index int) *WebhookCache {
	(*wh)[len(*wh)-1], (*wh)[index] = (*wh)[index], (*wh)[len(*wh)-1]
	(*wh) = (*wh)[:len(*wh)-1]
	return wh
}

// Updates a Webhook in the cache
func (wh *WebhookCache) Update(webhook Webhook) *WebhookCache {
	for index, element := range *wh {
		if element.Id == webhook.Id {
			(*wh)[index] = webhook
			break
		}
	}

	return wh
}

// Finds a Team in the cache
func (tc *TeamCache) Find(callback func(Team) bool) (Team, error) {
	for _, element := range *tc {
		if callback(element) {
			return element, nil
		}
	}

	return Team{}, errors.New("team not found")
}

// Basic cache structure
type Cache struct {
	Users    *UserCache
	Spaces   *SpaceCache
	Webhooks *WebhookCache
	Teams    *TeamCache
}

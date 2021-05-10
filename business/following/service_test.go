package following_test

import (
	"testing"

	"github.com/iswanulumam/dependency-injection-go/business/following"
)

type InMemoryRepository struct{}

func newInMemoryRepository() following.Repository {
	return InMemoryRepository{}
}

func (repo InMemoryRepository) Store(entity following.Following) error {
	return nil
}

func TestInsertFollowing(t *testing.T) {
	followingRepository := newInMemoryRepository()
	followingService := following.NewService(followingRepository)
	if err := followingService.Insert("root from test", "I'm Root from test"); err != nil {
		t.Errorf("Got %v, expect nil  when inserting new following", err.Error())
	}
}

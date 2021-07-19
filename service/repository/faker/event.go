package faker

import (
	"log"
	"time"

	faker "github.com/bxcodec/faker/v3"
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type eventRepository struct {
}

func NewEventRepository() service.EventRepository {
	return &eventRepository{}
}

func (r *eventRepository) GetEvents() (result []models.Event, err error) {
	var a, b, c models.Event

	err = faker.FakeData(&a)
	if err != nil {
		log.Printf("[repository][postgres][Event] failed while generate fake data, %+v\n", err)
		return result, err
	}

	err = faker.FakeData(&b)
	if err != nil {
		log.Printf("[repository][postgres][Event] failed while generate fake data, %+v\n", err)
		return result, err
	}

	err = faker.FakeData(&c)
	if err != nil {
		log.Printf("[repository][postgres][Event] failed while generate fake data, %+v\n", err)
		return result, err
	}

	// NOTE: mock processing time
	time.Sleep(2 * time.Second)

	result = append(result, a)
	result = append(result, b)
	result = append(result, c)

	return result, nil
}

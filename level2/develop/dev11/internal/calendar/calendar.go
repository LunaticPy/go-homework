package calendar

import (
	"encoding/json"
	"errors"
	"strconv"
	"task/level2/calendar/config"
	"time"
)

var (
	errNoEvent = errors.New("no such event")
)

type Calendar struct {
	Data map[string]Event `json:"data"`
}

type Event struct {
	Id      string    `json:"id"`
	Context string    `json:"Context"`
	Date    time.Time `json:"date"`
}

func NewCalendar(cfg *config.Calendar) *Calendar {
	return &Calendar{
		Data: make(map[string]Event),
	}
}

func (c *Calendar) CreateEvent(data []byte) error {
	ev := Event{}
	err := json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}
	ev.Id = strconv.Itoa(len(c.Data))
	c.Data[ev.Id] = ev
	return nil
}
func (c *Calendar) UpdateEvent(data []byte) error {
	ev := Event{}
	err := json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}
	_, ok := c.Data[ev.Id]
	if !ok {
		return errNoEvent
	}

	c.Data[ev.Id] = ev

	return nil
}

func (c *Calendar) DelEvent(data []byte) error {
	ev := Event{}
	err := json.Unmarshal(data, &ev)
	if err != nil {
		return err
	}
	_, ok := c.Data[ev.Id]
	if !ok {
		return errNoEvent
	}

	delete(c.Data, ev.Id)
	return nil
}

func (c *Calendar) GetMonth(month time.Time) []Event {
	evList := make([]Event, 0)
	for i := range c.Data {
		diff := c.Data[i].Date.Sub(month).Hours() / 24
		if diff <= 15 {
			evList = append(evList, c.Data[i])
		}
	}
	return evList
}

func (c *Calendar) GetDay(day time.Time) []Event {
	evList := make([]Event, 0)
	for i := range c.Data {
		diff := c.Data[i].Date.Sub(day).Hours()
		if diff <= 24 {
			evList = append(evList, c.Data[i])
		}
	}
	return evList
}
func (c *Calendar) GetWeek(date time.Time) []Event {
	evList := make([]Event, 0)

	for i := range c.Data {
		diff := c.Data[i].Date.Sub(date).Hours() / 24
		if diff <= 3 {
			evList = append(evList, c.Data[i])
		}
	}
	return evList
}

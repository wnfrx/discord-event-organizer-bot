package cron

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
	"github.com/wnfrx/discord-event-organizer-bot/models"
)

type cronJobHandler struct {
	c    *cron.Cron
	jobs []models.CronJob

	session *discordgo.Session
}

func NewJobHandler(
	session *discordgo.Session,
) *cronJobHandler {
	return &cronJobHandler{
		c:       cron.New(),
		session: session,
	}
}

func (h *cronJobHandler) InitJobHandlers() (err error) {
	h.AddJob("testJob", models.CronJobScheduleEveryMinute, h.testJob)
	h.AddJob("testFiveMinuteJob", "0 */5 * * * *", h.testFiveMinuteJob)

	return nil
}

func (h *cronJobHandler) start() {
	h.c.Start()
}

func (h *cronJobHandler) stop() {
	h.c.Stop()
}

func (h *cronJobHandler) register(j models.CronJob) (err error) {
	h.jobs = append(h.jobs, j)

	err = h.c.AddFunc(j.Schedule, j.Fn)
	if err != nil {
		fmt.Printf("Error occured while registering job %s\n", j.Name)
		return err
	}

	return nil
}

func (h *cronJobHandler) AddJob(name, schedule string, fn func()) (err error) {
	h.stop()

	job := models.CronJob{
		Name:     name,
		Schedule: schedule,
		Fn:       fn,
	}

	if err = h.register(job); err != nil {
		return err
	}

	h.start()

	return nil
}

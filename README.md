
# Discord Event Organizer Bot

This project is made using [DiscordGo](https://github.com/bwmarrin/discordgo) as low level bindings to the [Discord](https://discord.com/) chat client API, and also implemented [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

This bot is made to organizing any events (offline, online, and etc.) and reminds all participants the event schedule.


## Features

- Show All Events (Ongoing, Upcoming, Past)
- Create Event
- Join Event
- Cancel Event
- Add Participant to an Event
- Event Schedule Reminder

  
## Run Locally

Clone the project

```bash
  git clone git@github.com:wnfrx/discord-event-organizer-bot.git
```

Go to the project directory

```bash
  cd my-project
```

Start the application

```bash
  go run main.go
```

  
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`BOT_APPLICATION_ID`

`BOT_TOKEN`

  
## Database Table

For database, this project is using PostgreSQL as main database repository (you can add another repository based on your preferences).

**Class Diagram:**

![App Database SQL Table](https://via.placeholder.com/468x300?text=App+Screenshot+Here)

  

## Command Reference

#### Ping

```bash
  /ping
```

#### Show Events

```bash
  /show-event [filter]
```

| Parameter | Type     | Description                                                  |
| :-------- | :------- | :----------------------------------------------------------- |
| `filter`  | `string` | **Optional**. Values: upcoming (default), ongoing, past, all |




*More commands to be added soon.*

  
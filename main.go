package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/wnfrx/discord-event-organizer-bot/config"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

func init() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func catch(err error) {
	if err != nil {
		log.Fatalf("Something went wrong while preparing app, %+v", err)
	}
}

func main() {
	app := config.NewConfig()
	catch(app.InitDiscordSession())
	catch(app.InitServices())
	catch(app.Run())

	defer app.Stop()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}

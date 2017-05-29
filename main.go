package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"github.com/qbunt/eta-announce-go/ifttt"
	"github.com/qbunt/eta-announce-go/twilio"
	"github.com/qbunt/eta-announce-go/utils"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func setupServer() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		from := c.Query("from")
		to := c.Query("to")
		phone := c.Query("phone")

		timeInTraffic, myETA, err := calcEta(from, to)
		check(err)
		if phone != "" {
			twilio.Notify(phone, myETA)
			ifttt.Notify(timeInTraffic, myETA)
		}

		pretty.Println(myETA)
		c.String(http.StatusOK, "'%s ETA' generated at %s", myETA, time.Now())
	})
	router.Run(":" + os.Getenv("SERVER_PORT"))
}

func calcEta(f string, t string) (string, string, error) {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("MAPS_KEY")))
	check(err)

	origin := []string{f}
	destination := []string{t}

    // add a walking time buffer to start time
	nowish := strconv.FormatInt(time.Now().Add(parseTime(os.Getenv("DEPARTURE_OFFSET"))).Unix(), 10)

	eta := &maps.DistanceMatrixRequest{
		Origins:       origin,
		Destinations:  destination,
		Mode:          "driving",
		DepartureTime: nowish,
		TrafficModel:  "best_guess",
		Units:         "imperial",
	}
	// get the actual ETA off the distance matrix lib
	etaResponse := getETA(eta, c)
	predictedETA, err := time.ParseDuration(etaResponse.Rows[0].Elements[0].DurationInTraffic.String())
	delayedTime := time.Now().Add(predictedETA)
	check(err)

	timeInTraffic := fmt.Sprint(utils.Round(predictedETA.Minutes(), .5, 0))

	completedEta := delayedTime.Format(time.Kitchen)
	return timeInTraffic, completedEta, err
}

func main() {
	err := godotenv.Load()
	check(err)
	setupServer()
}

func parseTime(s string) time.Duration {
	parsedTime, err := time.ParseDuration(s)
	check(err)
	return parsedTime
}

func getETA(body *maps.DistanceMatrixRequest, client *maps.Client) *maps.DistanceMatrixResponse {
	resp, err := client.DistanceMatrix(context.Background(), body)
	check(err)
	return resp
}

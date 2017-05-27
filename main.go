package main

import (
    "googlemaps.github.io/maps"
    "github.com/kr/pretty"
    "golang.org/x/net/context"
    "log"
    "time"
    "strconv"
    "github.com/joho/godotenv"
    "os"
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

func check(err error) {
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }
}

func setupServer() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        from := c.Query("from")
        to := c.Query("to")

        myResp := calcEta(from, to)
        pretty.Println(myResp)
        c.String(http.StatusOK, "%s", myResp)
    })
    router.Run(":"+os.Getenv("SERVER_PORT"))
}

func calcEta(from string, to string) string {
    c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("MAPS_KEY")))
    check(err)

    origin := []string{from}
    destination := []string{to}

    now := strconv.FormatInt(time.Now().Unix(), 10)
    eta := &maps.DistanceMatrixRequest{
        Origins:       origin,
        Destinations:  destination,
        Mode:          "driving",
        DepartureTime: now,
        TrafficModel:  "best_guess",
        Units:         "imperial",
    }
    // get the actual ETA off the distance matrix lib
    etaResponse := getETA(eta, c)
    apiEta := etaResponse.Rows[0].Elements[0].DurationInTraffic.String()

    // delay the time by a touch to get out to the car
    delayedTime := time.Now().Add(addDelay(os.Getenv("DEPARTURE_OFFSET")))
    predictedETA, err := time.ParseDuration(apiEta)
    check(err)

    completedEta := delayedTime.Add(predictedETA).Format(time.Kitchen)
    return completedEta
}

func main() {
    err := godotenv.Load()
    check(err)

    setupServer()
}

func addDelay(s string) time.Duration {
    delay, err := time.ParseDuration(s)
    check(err)
    return delay
}

func getETA(body *maps.DistanceMatrixRequest, client *maps.Client) (*maps.DistanceMatrixResponse) {
    resp, err := client.DistanceMatrix(context.Background(), body)
    check(err)
    return resp
}

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
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("MAPS_KEY")))

    if err != nil {
        log.Fatalf("Fatal error %s", err)
    }

    origin := []string{os.Getenv("WORK_ADDRESS")}
    destination := []string{os.Getenv("HOME_ADDRESS")}
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
    etaResponse, etaErr := getETA(eta, c)

    if etaErr != nil {
        log.Fatalf("fatal error: %s", etaErr)
    }
    myResponse := etaResponse.Rows[0].Elements[0]

    pretty.Println(myResponse.DurationInTraffic.String())
    pretty.Println(myResponse.Duration.String())
}

func getETA(body *maps.DistanceMatrixRequest, client *maps.Client) (*maps.DistanceMatrixResponse, error) {
    resp, err := client.DistanceMatrix(context.Background(), body)
    return resp, err
}

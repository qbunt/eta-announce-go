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

    if err!=nil {
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
    }

    resp, err := c.DistanceMatrix(context.Background(), eta)
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }

    pretty.Println(resp)

}

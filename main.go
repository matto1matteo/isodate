package main

import (
    "fmt"
    "time"
)

func isoDate(d time.Time) string {
    return fmt.Sprintf(
        "%04d%02d%02d%02d%02d%02d", 
        d.Year(), d.Month(), d.Day(), d.Hour(), d.Minute(), d.Second(),
    )
}


func main() {
    now := time.Now()

    fmt.Println(isoDate(now))
}

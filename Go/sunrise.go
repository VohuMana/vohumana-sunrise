package main

import 
(
    "fmt"
    "math"
    "time"
)

// Constants to be used for calculations
const (
    twoPi = math.Pi * 2.0
    sunRadius = 0.53
    atmosphericRefraction = 34.0 / 60.0
    degToRad = math.Pi / 180.0
    radToDeg = 180.0 / math.Pi
)

// CalculateSunriseAndSet calculates the sunrise and sunset for a give lat and lon
func CalculateSunriseAndSet(lat, lon float64, timezone, day, month, year, hour int, isDst bool) (time.Time, time.Time) {
    return time.Now(), time.Now()
}

// ForcePiRange forces a number in the range [0, PI]
func ForcePiRange(x float64) float64 {
    b := x / twoPi
    a := twoPi * (b - float64(int(x)))
    if (a < 0) {
        a += twoPi
    }

    return a
}

// ConvertToJulian will convert a given year/month/day/hour into the Julian representation of that time
func ConvertToJulian(year, month, day, hour int) float64 {
    julian := -7 * (year + (month + 9) / 12) / 4 + 275 * month / 9 + day
    julian += year * 367
    return float64(float64(julian) - 730531.5 + float64(hour) / 24.0)
}

// CalculateHourAngle will calculate the hour angle at a lattitude
func CalculateHourAngle(lat, declination float64) float64 {
    dfo := degToRad * (0.5 * sunRadius + atmosphericRefraction)

    if (lat < 0.0) {
        dfo = -dfo
    }

    fo := math.Tan(declination + dfo) * math.Tan(lat * degToRad)

    if (fo > 0.999999) {
        fo = 1.0
    }

    return math.Asin(fo) + math.Pi / 2.0
}

// CalculateTwilightHourAngle calculates the hour angle for twilight at a lattitude
func CalculateTwilightHourAngle(lat, declination float64) float64 {
    dfo := degToRad * 6.0

    if (lat < 0.0) {
        dfo = -dfo
    }

    fo := math.Tan(declination + dfo) * math.Tan(lat * degToRad)

    return math.Asin(fo) + math.Pi / 2.0
}

// CalculateL TODO: I need a book on how this math works
func CalculateL(jullianDate float64) float64 {
    return ForcePiRange(280.461 * degToRad + .9856474 * degToRad * jullianDate)
}

// CalculateG TODO: Seriously how were these magic numbers calculated
func CalculateG(jullianDate float64) float64 {
    return ForcePiRange(357.528 * degToRad + .9856003 * degToRad * jullianDate)
}

// EclipticLongitude I'm not sure quite what this does
func EclipticLongitude(L, G float64) float64 {
    return ForcePiRange(L + 1.915 * degToRad * math.Sin(G) + 0.02 * degToRad * math.Sin(2.0 * G))
}

func main() {
    fmt.Println("HelloWorld")
}
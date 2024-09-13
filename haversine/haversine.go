package main

import (
	"fmt"
	"math"
)

const earthRadius = 6371 // Earth radius in kilometers

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert latitude and longitude from degrees to radians
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	// Calculate differences
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := earthRadius * c

	return distance
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func main() {
	// Example points (latitude and longitude in degrees)
	lat1 := -8.477043900671797
	lon1 := 114.12958367303607
	lat2 := -8.476691953742693
	lon2 := 114.12922783878675

	// Calculate distance
	distance := haversine(lat1, lon1, lat2, lon2)

	// Output result
	fmt.Printf("Distance between the points: %.2f km\n", distance)
}

package usecase

import (
	"math"
)

func IsInRange(lat1, lon1, lat2, lon2, radiusKm float64) bool {
	// Haversine formula to calculate the distance between two points on the Earth
	const R = 6371 // Radius of the Earth in kilometers
	dLat := (lat2 - lat1) * (3.141592653589793 / 180)
	dLon := (lon2 - lon1) * (3.141592653589793 / 180)
	a := (math.Sin(dLat/2) * math.Sin(dLat/2)) + (math.Sin(lat1*(3.141592653589793/180)) * math.Sin(lat2*(3.141592653589793/180)) * math.Sin(dLon/2) * math.Sin(dLon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c // Distance in kilometers

	return distance <= radiusKm
}

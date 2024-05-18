package main

import (
	"math"
)

func minimumRefuelStops(stations [][]int, targetDistance int, startFuel int) int {
	dp := make([][]int, len(stations)+1)
	dp[0] = make([]int, len(stations)+1)
	dp[0][0] = startFuel
	for i := 1; i <= len(stations); i++ {
		dp[i] = make([]int, len(stations)+1)
		if startFuel >= stations[i-1][0] {
			dp[i][0] = startFuel
		}
	}

	for stationIndex := 1; stationIndex <= len(stations); stationIndex++ {
		for stops := 1; stops <= len(stations); stops++ {
			if dp[stationIndex-1][stops-1] >= stations[stationIndex-1][0] {
				maxDistExcludeCurrentStation := dp[stationIndex-1][stops]
				maxDistIncludeCurrentStation := 0
				if stationIndex == 1 {
					maxDistIncludeCurrentStation = dp[stationIndex-1][stops-1] + stations[stationIndex-1][1]
					// Needed to cpmment this because this was giving the remaining fuel at station S in the car and that was causing
					//trouble to identify if the target distance can be achieved or not
					//- (stations[stationIndex-1][0] - 0)
				} else {
					maxDistIncludeCurrentStation = dp[stationIndex-1][stops-1] + stations[stationIndex-1][1] //- (stations[stationIndex-1][0] - stations[stationIndex-2][0])
				}

				dp[stationIndex][stops] = int(math.Max(float64(maxDistExcludeCurrentStation), float64(maxDistIncludeCurrentStation)))
				if dp[stationIndex][stops] >= targetDistance {
					return stops
				}
			}
		}
	}

	return -1
}

func TestMinimumRefuelStops() int {

	stations := [][]int{{10, 60}, {20, 25}, {30, 30}, {60, 40}}
	targetDistance := 120
	startFuel := 10
	return minimumRefuelStops(stations, targetDistance, startFuel)
}

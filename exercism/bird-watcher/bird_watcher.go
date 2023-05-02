package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	acc := 0
	for _, b := range birdsPerDay {
		acc += b
	}

	return acc
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	acc := 0

	for i, b := range birdsPerDay {
		if i/7 == week-1 {
			acc += b
		}
	}

	return acc
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}

	return birdsPerDay
}

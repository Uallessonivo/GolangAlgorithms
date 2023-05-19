package arithmetic

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	rate := float64(productionRate) * successRate / 100
	return rate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / float64(60))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var (
		cars_cost int = 10000
		ten_cars  int = 95000
	)
	return uint(((carsCount / 10) * ten_cars) + ((carsCount % 10) * cars_cost))
}

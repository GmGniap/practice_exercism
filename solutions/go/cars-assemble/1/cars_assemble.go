package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRate = successRate / 100
	var numCarsProduced float64 = float64(productionRate) * successRate
	return numCarsProduced
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successfulCarsPerHr := CalculateWorkingCarsPerHour(productionRate, successRate)
	numCarsProducedPerMin := successfulCarsPerHr / 60
	return int(numCarsProducedPerMin)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupOfTen int = int(carsCount / 10)
	var carRemaining int = carsCount % 10
	var groupOfTenCost int = 95000
	var individualCost int = 10000
	var totalCostForGroup int = groupOfTen * groupOfTenCost
	var totalCostForRemaining int = carRemaining * individualCost
	return uint(totalCostForGroup + totalCostForRemaining)
}

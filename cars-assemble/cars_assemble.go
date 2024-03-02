package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// OneCarCost is the total cost of an individual car
const OneCarCost = 10000
// GroupCarCost is the total cost of a group of cars.
const GroupCarCost = 95000
// GroupSize is the number of cars in a grouping
const GroupSize = 10

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupCost int = carsCount / GroupSize * GroupCarCost
	var indivCost int = carsCount % GroupSize * OneCarCost
	return uint(groupCost + indivCost)
}

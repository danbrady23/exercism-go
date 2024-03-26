package elon

import "fmt"

// Drive method updates the values in Car struct by reference.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance method returns a string describing the distance
// driven in meters for a specific Car.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery method returns a string describing the state
// of the battery charge for a specific Car.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)

}

// CanFinish method returns a boolean indicating if a Car will be able
// to finish a track of supplied distance.
func (c *Car) CanFinish(trackDistance int) bool {
	time := c.battery / c.batteryDrain
	return trackDistance <= (c.speed * time)
}

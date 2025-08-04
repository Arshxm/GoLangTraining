package main

type Car struct {
	battery int
	speed   int
}

func NewCar(speed, battery int) *Car {
	return &Car{
		battery: battery,
		speed:   speed,
	}
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	if car.battery == 100 {
		return
	}
	car.battery += minutes / 2
	if car.battery > 100 {
		car.battery = 100
	}
}
func TryFinish(car *Car, distance int) string {
	
	speed := GetSpeed(car)
	battery := GetBattery(car)
	time := float64(distance) / float64(speed)

	if battery < distance/2 {
		battery = 0
		return ""
	}else {
		battery -= int(distance)/2
		return time
	}
	
}

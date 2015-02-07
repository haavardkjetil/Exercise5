package	main

import(
	"driver"
)

func main() {
	driver.Init()

	println(driver.GetStopSignal())
	println(driver.GetObstructionSignal())
	x := driver.GetFloorSensorSignal()
	println(x)

}
package	main

import(
	"driver"
	"log"
)

func main() {
	// Initialize hardware
    if !driver.Init() {
        log.Fatal("Unable to initialize elevator hardware!\n")
    }

    println("Press STOP button to stop elevator and exit program.\n")

    driver.SetMotorDirection(-1)

    for {
        // Change direction when we reach top/bottom floor
    	if driver.GetFloorSensorSignal() == 3 {
            driver.SetMotorDirection(-1)
        } else if driver.GetFloorSensorSignal() == 0 {
            driver.SetMotorDirection(1)
        }

        // Stop elevator and exit program if the stop button is pressed
        if driver.GetStopSignal() {
        	driver.SetMotorDirection(0)
            break
        }
    }
}
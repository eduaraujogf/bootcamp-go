package pkg

import "fmt"

func Climate() {
	var (
		temperature         float32
		humidity            float32
		atmosphericPressure float32
	)
	temperature, humidity, atmosphericPressure = 11.5, 81.3, 1.1
	fmt.Println(temperature, humidity, atmosphericPressure)

}

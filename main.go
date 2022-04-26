package main

import (
	"./log"
	// "./user"
	// "./cars"
	// "./calc"
	// "./channel"
	// "./osy"
	// "./iot"
	// "./examples"
	// "./network"
	"./users"
)

func main() {
	log.SpeedTest(func() {
		// user.ShowUser()
		// cars.CreateCars()
		// calc.GetFactorials()
		// channel.WaitGroupExample()
		// osy.WorkWithFile()
		// iot.IotPeople()
		// examples.MyIo()
		// examples.WriteTxtFile()
		// network.OnlyServer()
		users.UserSelectors()
	})
}

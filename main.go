package main

import (
	"my_golang_pkgs/log"
	// "my_golang_pkgs/user"
	// "my_golang_pkgs/cars"
	// "my_golang_pkgs/calc"
	"my_golang_pkgs/channel"
	// "my_golang_pkgs/osy"
	// "my_golang_pkgs/iot"
	// "my_golang_pkgs/examples"
	// "my_golang_pkgs/network"
	// "my_golang_pkgs/users"
	// "my_golang_pkgs/helpers"
)

func main() {
	log.SpeedTest(func() {
		// user.ShowUser()
		// cars.CreateCars()
		// calc.GetFactorials()
		// channel.WaitGroupExample()
		channel.RunProcessWithMutex()
		// osy.WorkWithFile()
		// iot.IotPeople()
		// examples.MyIo()
		// examples.WriteTxtFile()
		// examples.Closure()
		// network.OnlyServer()
		// users.UserSelectors()
		// helpers.TypesDetection()
	})
}

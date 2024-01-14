package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

var logger service.Logger

// myService is a simple struct to represent our service.
type myService struct{}

func (m *myService) Start(s service.Service) error {
	// Start your service logic here.
	go m.run()
	return nil
}

func (m *myService) Stop(s service.Service) error {
	// Stop your service logic here.
	return nil
}

func (m *myService) run() {
	// Your service logic goes here.
	for {
		log.Println("MyService is running...")
	}
}

func main() {
	svcConfig := &service.Config{
		Name:        "MyService",
		DisplayName: "My Go Service",
		Description: "A simple Go service running on Windows",
	}

	prg := &myService{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatalf("Valid actions: %q\n", service.ControlAction)
		}
		return
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

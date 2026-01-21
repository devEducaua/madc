package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Service struct {
    name string
    path string
	stats bool
}

func (s *Service) inService(command...string) error {
    path, err := exec.LookPath("docker")
    if err != nil {
        return err
    }

	cmd := exec.Command(path, command...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (s *Service) upService() error {
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    s.inService("compose", "-f", home + s.path, "up", "-d")
    
    fmt.Printf("service: %v is up\n", s.name)
	s.stats = true
	return nil
}

func (s *Service) downService() error {
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    s.inService("compose", "-f", home + s.path, "down")
    
	s.stats = false
    fmt.Printf("service: %v is down\n", s.name)
	return nil
}

func (s *Service) listServices() {
		
}

func getServiceByName(services []Service, serviceArg string) (*Service, error) {
	for _, s := range services {
		if s.name == serviceArg {
			return &s, nil
		}
	}
    return nil, errors.New("this service doens't exist")
}

func readConfigFile() []Service {
	services := []Service{}

	contents, err := os.ReadFile("madc.conf")
    if err != nil {
        log.Fatal("file reading error", err)
    }
	cont := string(contents)
	lines := strings.Split(cont, "\n")

	for i := 0; i < len(lines) - 1; i++ {
		idx := strings.Index(lines[i], ":")
		if idx < 0 {
			os.Exit(1)
		}

		var s Service
		s.name = lines[i][0:idx]
		s.path = lines[i][:idx]
		s.stats = false
		
		services = append(services, s)
	}

	return services
}

func main() {
	services := readConfigFile()

	args := os.Args[1:]

    if len(args) == 0 {
		for _, s := range services {
			fmt.Printf("%v: %v\n", s.name, s.stats)
		}
		os.Exit(0)
    }

    s, err := getServiceByName(services, args[0])
    if err != nil {
        log.Fatal(err)
    }

    if len(args) <= 1 {
        log.Fatalf("pass the command for the service: %v", s.name)
    }

    switch args[1] {
        case "u":
            s.upService()
        case "d":
            s.downService()
        default:
            log.Fatalf("command not found: %v", args[1])
    }
}

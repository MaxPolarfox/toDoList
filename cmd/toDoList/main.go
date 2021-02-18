package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/MaxPolarfox/toDoList/toDoList"
	"github.com/MaxPolarfox/toDoList/pkg/types"
)

const ServiceName = "toDoList"
const EnvironmentVariable = "APP_ENV"

func main() {
	// Load current environment
	env := os.Getenv(EnvironmentVariable)

	// load config options
	options := loadEnvironmentConfig(env)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", options.Port) )
	if err != nil {
		log.Fatalf("failed to listen %v: %v", options.Port, err)
	}

	s := toDoList.Server{}

	grpcServer := grpc.NewServer()

	toDoList.RegisterToDoListServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC over port %v: %v",options.Port, err)
	}
}

// loadEnvironmentConfig will use the environment string and concatenate to a proper config file to use
func loadEnvironmentConfig(env string) types.Options {
	configFile := "config/" + ServiceName + "/" + env + ".json"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		panic(err)
	}
	return parseConfigFile(configFile)
}

func parseConfigFile(configFile string) types.Options {
	var opts types.Options
	byts, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byts, &opts)
	if err != nil {
		panic(err)
	}

	return opts
}
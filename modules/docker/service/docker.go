package service

import "github.com/docker/docker/client"

// DockerService is the service for docker module.
type DockerService struct {
	Client *client.Client
}

// NewDockerService creates and returns a new docker service.
func NewDockerService() *DockerService {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &DockerService{
		Client: client,
	}
}

// Get

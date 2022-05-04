package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"os"
)

const (
	datapath      = "data"
	containername = "ethpostgres"
	port          = "5432"
	pass          = "secretpassword"
)

func main() {
	Stop()
	err := Deploy()
	if err != nil {
		fmt.Println(err)
		Stop()
	}

}
func Stop() error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	err = cli.ContainerKill(ctx, containername, "")
	err = cli.ContainerRemove(ctx, containername, types.ContainerRemoveOptions{})

	return err
}

func Deploy() error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, containername, types.ContainerStartOptions{})
	if err != nil {
		//service is not started
		fmt.Errorf("error %s", err)

	} else {
		return nil
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/postgres", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)

	if _, err := os.Stat(datapath); os.IsNotExist(err) {
		os.Mkdir(datapath, 0777)
	}
	natPort := fmt.Sprintf("%s/tcp", port)
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(natPort): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}
	goPath := os.Getenv("GOPATH")

	dataDirPath := fmt.Sprintf("%s/src/github.com/ezuhl/eth/cmd/postgres/%s", goPath, datapath)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        "postgres",
		Volumes:      map[string]struct{}{fmt.Sprintf("%s:/var/lib/postgresql/data", dataDirPath): struct{}{}},
		Env:          []string{fmt.Sprintf("POSTGRES_PASSWORD=%s", pass)},
		ExposedPorts: nat.PortSet{nat.Port(natPort): struct{}{}},
	}, hostConfig, nil, nil, containername)
	if err != nil {
		return err
	}

	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

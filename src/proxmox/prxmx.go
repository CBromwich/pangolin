package prxmx

import (
	"os"
	"github.com/luthermonson/go-proxmox"
	"github.com/joho/godotenv"
	"context"
	"log"
	//"fmt"
	//"encoding/json"
)

type ClientConfig struct {
	ProxClient *proxmox.Client
}


func InitClient() *proxmox.Client {
	godotenv.Load()
	proxmoxToken  := os.Getenv("PROXMOX_TOKEN")
	proxmoxSecret := os.Getenv("PROXMOX_SECRET")

	client := proxmox.NewClient("https://proxmox.matthewsbromwich.com/api2/json",
		proxmox.WithAPIToken(proxmoxToken, proxmoxSecret))
	
	return client
}

func (client *ClientConfig) GetVMs() proxmox.VirtualMachines {
	node, _ := client.ProxClient.Node(context.Background(), "pve")
	virtualmachines, _ := node.VirtualMachines(context.Background())
	return virtualmachines
}

func (client *ClientConfig) GetVM(id int) proxmox.VirtualMachine {
	node, _ := client.ProxClient.Node(context.Background(), "pve")
	vm, err := node.VirtualMachine(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}
	return *vm
}

func GetHostName(vm proxmox.VirtualMachine) (string, error) {
	hostname, err := vm.AgentGetHostName(context.Background())
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func GetIP(vm proxmox.VirtualMachine) ([]*proxmox.AgentNetworkIface, error) {
	ip, err := vm.AgentGetNetworkIFaces(context.Background())
	if err != nil {
		return []*proxmox.AgentNetworkIface{}, err
	}
	return ip, nil
}

func ParseIP(iface []*proxmox.AgentNetworkIface) string {
	for _, vminterface := range iface {
		for _, ip := range vminterface.IPAddresses {
			if ip.IPAddressType == "ipv4" {
				return ip.IPAddress
			}
		}
	}
	return ""
}

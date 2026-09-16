package main

import (
	"log"
	"github.com/rivo/tview"
	"matthewsbromwich.com/pangolin/proxmox"
	"github.com/luthermonson/go-proxmox"
	//"context"
)
func main() {

	proxmoxClient := prxmx.InitClient()
	clientCfg := prxmx.ClientConfig{ProxClient: proxmoxClient}
	app := tview.NewApplication()
	newTextBox := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	header := newTextBox("Pangolin")
	clientList := tview.NewList().
		AddItem("Quit", "Press Q to exit", 'q', func () {app.Stop()})
	infoForm := tview.NewForm()

	//main grid
	grid := tview.NewGrid().
		SetRows(1, 5).
		SetColumns(3, 3).
		AddItem(header, 0, 0, 1, 6, 0, 0, false).
		SetBorders(true).
		AddItem(clientList, 1, 0, 5, 5, 0, 0, true).
		AddItem(infoForm,    1, 5, 5, 1, 0, 0, false)
		
	//TODO	
	// Get relevant information from proxmox for IP/CPU/RAM usage
	
	PopulateList(clientList, infoForm, clientCfg.GetVMs())

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		log.Fatal(err)
	}
}

func PopulateList(list *tview.List, panel *tview.Form, vms proxmox.VirtualMachines) {
	listLen := list.GetItemCount()
	for _, machine := range vms {
		list.InsertItem(listLen-1, machine.Name, machine.Status, 0, func() {
			panel.Clear(true)
			hostname, _ := prxmx.GetHostName(*machine)
			iface, _ := prxmx.GetIP(*machine)
			ip := prxmx.ParseIP(iface)
			panel.AddTextView("Hostname: ", hostname, 20, 1, false, false)
			panel.AddTextView("IP:       ", ip, 20, 1, false, false)
		})
	}
}


package main

import (
	"fmt"
	"net"
	"sync"
)

// ArpTable represents an ARP table, i.e., a mapping of hardware MAC addresses to IP addresses.
type ArpTable struct {
	sync.Mutex
	arpTable map[string]net.IP
}

// Add adds an ARP table entry.
func (a *ArpTable) Add(hwAddr net.HardwareAddr, ip net.IP) {
	a.Lock()
	defer a.Unlock()

	a.arpTable[string(hwAddr)] = ip
}

// Print prints the ARP table.
func (a *ArpTable) Print() {
	fmt.Printf("ARP table (%d):\n", len(a.arpTable))
	for mac, ip := range a.arpTable {
		fmt.Printf(" - %v == %v\n",
			net.HardwareAddr([]byte(mac)),
			net.IP(ip))
	}
}

// NewArpTable initialized an ARP table.
func NewArpTable() *ArpTable {
	at := &ArpTable{}
	at.arpTable = make(map[string]net.IP)
	return at
}
package main

import (
	"fmt"
)

type Virtmach struct{
	ip string
	hostname string
	diskgb int
	ram int
}

// ipset(string) - change IP address
// use a pointer here as we are making a change
func (vm *Virtmach)ipset(ip string){
	vm.ip = ip // change the value of ip associated with Virtmach
}

// diskexpand(int) - expand disk by int
// use a pointer here as we are making a change
func (vm *Virtmach)diskexpand(gb int){
	vm.diskgb = vm.diskgb + gb // increase
}

func(vm *Virtmach) increaseram(ram int) {
	vm.ram = vm.ram + ram
}
func (vm *Virtmach)display(){
	fmt.Println("Primary IP Address:", vm.ip) // primary IP address
	fmt.Println("Hostname:", vm.hostname)     // hostname
	fmt.Println("Disk GB:", vm.diskgb)        // diskgb
	fmt.Println("RAM:", vm.ram)               // ram
}

func main() {

	// create a virtual machine
	vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8} // ip, hostname, diskgb, ram

	vm1.display() // show the current state

	vm1.diskexpand(3) // increase by 3 GB

	vm1.ipset("192.168.77.33") // change the IP address

	vm1.display() // show the changes

	vm1.increaseram(4)
	vm1.display()
}
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


func (vm *Virtmach)ipset(ip string){
	vm.ip = ip // change the value of ip associated with Virtmach
}

func (vm *Virtmach) expanddisk(gb int){
	vm.diskgb = vm.diskgb + gb // increase
}

func(vm *Virtmach) increaseram(ram int) {
	vm.ram = vm.ram + ram
}
func (vm *Virtmach)display(){
	fmt.Printf("The virtual machine info: %+v \n", *vm)
}

func main() {

	// create a virtual machine
	vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8} // ip, hostname, diskgb, ram

	vm1.display() // show the current state

	vm1.expanddisk(3) // increase by 3 GB

	vm1.ipset("192.168.77.33") // change the IP address

	vm1.display() // show the changes

	vm1.increaseram(4)
	vm1.display()

	vm2 := Virtmach{"10.0.0.8", "elephant", 32, 16}

	vm2.display();

	slice := []*Virtmach{&vm1, &vm2}
	slice[1].display()
	slice[1].increaseram(16)
	slice[1].display()

	vm2.display()

	slice2 := []Virtmach{vm1, vm2}
	slice2[1].display()
	slice2[1].increaseram(-16)
	slice2[1].display()   //back to ram 16
	vm2.display() //but vm2 didn't change
}
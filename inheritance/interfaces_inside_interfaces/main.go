package main

type Species interface{
	eat()
}

type Animal interface {
	Species
	speak()
}

type Plants interface {
	Species
	getOxygen()
}


func main(){

}

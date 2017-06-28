package main

type Buffer struct {
}

func (b Buffer) process() {

}

func (b Buffer) get() {

}

var available = make(chan Buffer, 10)
var toController = make(chan Buffer)

func worker() {
	for {
		var b Buffer
		select {
		case b = <-available:
		default:
			b = Buffer{}
		}
		b.get()
		toController <- b
	}
}

func controller() {
	for {
		b := <-toController
		b.process()
		select {
		case available <- b:
		default:
		}
	}
}
func main() {
	go controller()
	go worker()
}

package main

import "strconv"

//type Mutex interface {
//	Lock(who string)
//	Unlock(who string)
//}
//
//func run(lock Mutex, who string) {
//	println(who + " starting")
//	i := 0
//	for {
//		lock.Lock(who)
//		j := i
//		for ; i < j+10; i++ {
//			time.Sleep(100 * time.Millisecond)
//			println(who + " " + strconv.Itoa(i))
//		}
//		lock.Unlock(who)
//	}
//}

type A interface {
	A(string) string
	B(string) string
}

type A2 struct {
	float64
}

type B struct {
	A2
	x string
}

type C struct {
	A2
	y string
	int
}

func (this A2) A(a string) string {
	return "A2(" + strconv.FormatFloat(this.float64, byte('f'), 3, 32) + ")::A(" + a + ")"
}

//	func (this *B) A(a string) string {
//		return this.x + ".A(" + a + ")"
//	}
func (this *B) B(a string) string {
	return this.x + ".B(" + a + ")"
}

func (this *C) A(a string) string {
	return this.A2.A(a) + " + " + this.y + "->A(" + a + ", " + strconv.Itoa(this.int) + ")"
}

func (this *C) B(a string) string {
	return this.y + "->B(" + a + ")"
}

func main() {
	//m := NewMutex2()
	//go run(m, "A")
	//go run(m, "B")
	//go run(m, "C")
	//go run(m, "D")
	//time.Sleep(time.Minute)

	//serverChan := make(chan *grpc.Server, 5)
	//go comms.RunMemberServer(serverChan)
	//time.Sleep(time.Second)
	//for i := 0; i < 10; i++ {
	//	go func(){
	//    log.Println("at " + comms.RunMemberClient("hello"))
	//	}()
	//}
	//time.Sleep(10 * time.Second)
	//A := <-serverChan
	//A.GracefulStop()
	i := 0
	for _, v := range []A{&B{x: "Bx"}, &C{y: "Cy", int: 7, A2: A2{
		float64: 1.0,
	}}} {
		println(v.A(strconv.Itoa(i)))
		i++
		println(v.B(strconv.Itoa(i)))
		i++

	}

}

package main

import(
	"log"
	"car/car"
	"car/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	car.PolicyHistoryDBInit()
	go car.InitProducer(config.GetMode())
	// view 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go car.InitConsumer(config.GetMode())
	e := car.RouteInit()

	e.Logger.Fatal(e.Start(":8082"))
}

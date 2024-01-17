package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// This is the max limit which Pizza shop can make in a Day
const numberOfPizzas int = 10

var pizzaMade, pizzaFailed, totalPiz int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makingPizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		// 5 is the max limit which rand can return
		delay := rand.Intn(5) + 1
		log.Printf("Recived order no %d\n", pizzaNumber)

		randomNumber := rand.Intn(10) + 1
		msg := ""
		success := false

		if randomNumber < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		totalPiz++

		fmt.Printf("Making pizza #%d. It will take %d seconds....\n\n", pizzaNumber, delay)

		// delay of the time which we take above : Basically time to make a pizza
		time.Sleep(time.Duration(delay) * time.Second)

		if randomNumber <= 2 {
			msg = fmt.Sprintf("** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if randomNumber <= 4 {
			msg = fmt.Sprintf("** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzaShop(pizzaMaker *Producer) {
	// Let's track all the pizzas
	var i = 0

	// Keep run always until and unless we don't rececive quit notificaition
	for {
		currentPizza := makingPizza(i)

		if currentPizza != nil {

			i = currentPizza.pizzaNumber
			select {
				// We tried to make the pizza and pushed into the data chanel (Not sure pizza is ready or not)
			case pizzaMaker.data <- *currentPizza:
				
			case quitChan := <-pizzaMaker.quit:
				// Close the both channel of pizzaMaker
				close(pizzaMaker.data)
				close(quitChan)
				return 
			}
		}
	}

}

func main() {
	// Seeding the random no
	rand.Seed(time.Now().UnixNano())

	color.Green("\n\nPizza shop is open,You can order !!\n\n")

	pizzaComing := Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// Run Producer in a different go routine 
	go pizzaShop(&pizzaComing)

	// Run consumer in a different go routine 

	for entry := range pizzaComing.data{
		if entry.pizzaNumber <= numberOfPizzas{
			// Till now we are not sure Pizza is ready or not : So will add a check there as well 

			if entry.success {
				// It means pizza is ready and out out delivery 
				color.Green(entry.message)
				color.Green("Order no %d is ready and out for delivery\n", entry.pizzaNumber)
			}else{
				color.Red(entry.message)
				color.Red("Sorry, We can not delivery your order , Order no %d\n", entry.pizzaNumber)
			}

		}else{
			color.Blue("\nWe are done for Today and Shop Closed\n")
			
			err := pizzaComing.Close()
			if err != nil {
				color.Red("Error while closing data : " , err)
			}
		}
	}


	// We are done with Whole Program here
	color.Yellow("---------------------------------------------------------------------------------------------------------------------------------------------")
	finalMsg := fmt.Sprintf("We got total %d orders in which we successfully made %d orders and we failed to make %d orders", totalPiz , pizzaMade , pizzaFailed)
	color.White(finalMsg)
	color.Yellow("---------------------------------------------------------------------------------------------------------------------------------------------")

}

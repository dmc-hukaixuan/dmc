package number

import (
	"fmt"
	"math/rand"
	"service-cool/kernel/system/util/number"
	"strconv"
	"time"
)

type DateChecksum struct{}

func (*DateChecksum) TicketNumberBuild() string {
	counter := number.TicketNumberCounterAdd()
	fmt.Println(" counter is  ", counter)
	// Pad ticket number with leading '0' to length 5.
	counter_s := fmt.Sprintf("%05d", counter)

	// The runtime environment is fixed, so rand.Intn will always return the same number.
	// To get a different number, you need to provide a different seed number for the generator
	rand.Seed(time.Now().Unix())

	// Create new ticket number
	ticket_number := time.Now().Format("20060102") + counter_s + strconv.Itoa(rand.Intn(10))
	fmt.Println("ticket_number ", ticket_number, " counter_s: ", counter_s)
	return ticket_number
}

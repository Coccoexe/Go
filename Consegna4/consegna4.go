package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Gruppo struct {
	nome     string
	nPalline int
}

type Tunnel struct {
	libero bool
}

func transumanza(g Gruppo, t chan Tunnel, c1 chan int, wg *sync.WaitGroup) {
	for g.nPalline > 0 {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mandaPersona(&g, t, c1)
	}
	wg.Done()
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int) {
	tunnel := <-t
	if tunnel.libero {
		tunnel.libero = false
		t <- tunnel
		fmt.Println("qui")
		select {
		case x := <-c1:
			fmt.Println("scontro ", g.nome)
			// GO si arrabbia se non usate una varabile...
			x = x - x
		case <-time.After(time.Second):
			tunnel := <-t
			tunnel.libero = true
			t <- tunnel
			fmt.Println("passato")
			g.nPalline = g.nPalline - 1
			fmt.Println("rimangono ", g.nPalline, " nel gruppo ", g.nome)
		}
		//time.Sleep(time.Duration(4) * time.Second)
	} else {
		//time.Sleep(time.Duration(4) * time.Second)
		c1 <- 1
		tunnel.libero = true
		t <- tunnel
		//time.Sleep(time.Duration(4) * time.Second)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	gruppo1 := Gruppo{"destra", 5}
	gruppo2 := Gruppo{"sinistra", 5}

	c1 := make(chan int, 1)

	tunnelChannel := make(chan Tunnel, 1)
	tunnel := Tunnel{true}
	tunnelChannel <- tunnel

	var wg sync.WaitGroup

	wg.Add(2)
	go transumanza(gruppo1, tunnelChannel, c1, &wg)
	go transumanza(gruppo2, tunnelChannel, c1, &wg)
	wg.Wait()

	time.Sleep(time.Minute)
}

/*
Scrivete un programma che simuli l’ordinazione, la cottura e l’uscita dei piatti in un ristorante.
10 clienti ordinano contemporaneamente i loro piatti.
In cucina vengono preparati in un massimo di 3 alla volta, essendoci solo 3 fornelli.
Il tempo necessario per preparare ogni piatto è fra i 4 e i 6 secondi.
Dopo che un piatto viene preparato, viene portato fuori da un cameriere, che impiega 3 secondi a portarlo fuori.
Ci sono solamente 2 camerieri nel ristorante.
	●Creare la strutture Piatto e Cameriere col relativo campo “nome”.
	●Creare le funzioni ordina che aggiunge il piatto a un buffer di piatti da fare;
	 creare la function cucina che cucina ogni piatto e lo mette in lista per essere consegnato;
	 creare la function consegna che fa uscire un piatto dalla cucina.
	●Ogni cameriere può portare solo un piatto alla volta.
	●Usate buffered channels per svolgere il compito.
	●Attenzione: se per cucinare un piatto lo mandate nel buffer fornello di capienza 3 e lo ritirate dopo 3 secondi,
	 non è detto che ritiriate lo stesso piatto che avete messo sul fornello.
	 Tenetelo in memoria. Ovviamente la vostra soluzione potrebbe differire dalla mia e questo hint potrebbe non servirvi.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Piatto struct {
	nome  string
	tempo int
}

type Ordine struct {
	piatto Piatto
	nome   string
}

type Cameriere struct {
	nome string
}

func ordina(o Ordine, staff chan Cameriere, fornelli chan int, wg *sync.WaitGroup) {
	var wgo sync.WaitGroup

	wgo.Add(1)
	go cucina(o, fornelli, &wgo)
	wgo.Wait()
	wgo.Add(1)
	go consegna(o, staff, &wgo)
	wgo.Wait()

	wg.Done()
}

func cucina(o Ordine, fornelli chan int, wgo *sync.WaitGroup) {
	fornelli <- 1
	time.Sleep(time.Duration(o.piatto.tempo) * time.Second)
	fmt.Println("Cucinato " + o.piatto.nome)
	<-fornelli
	wgo.Done()
}

func consegna(o Ordine, staff chan Cameriere, wgo *sync.WaitGroup) {
	c := <-staff
	time.Sleep(3 * time.Second)
	fmt.Println("consegnato " + o.piatto.nome + " a " + o.nome + " da " + c.nome)
	staff <- c
	wgo.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	c1 := Cameriere{"Cameriere1"}
	c2 := Cameriere{"Cameriere2"}
	staff := make(chan Cameriere, 2)
	staff <- c1
	staff <- c2

	fornelli := make(chan int, 3)

	p1 := Piatto{"Pasta", 6}
	p2 := Piatto{"Gnocchi", 6}
	p3 := Piatto{"Cotoletta", 5}
	p4 := Piatto{"Grigliata", 5}
	p5 := Piatto{"Gelato", 4}
	p6 := Piatto{"Tiramisu", 4}
	menu := []Piatto{p1, p2, p3, p4, p5, p6}

	clienti := []string{"Cliente1", "Cliente2", "Cliente3", "Cliente4", "Cliente5",
		"Cliente6", "Cliente7", "Cliente8", "Cliente9", "Cliente10"}

	wg.Add(len(clienti))
	for i := range clienti {
		choose := rand.Intn(len(menu))
		ordine := Ordine{menu[choose], clienti[i]}
		go ordina(ordine, staff, fornelli, &wg)
		fmt.Println("Ordinato " + ordine.piatto.nome + " da " + ordine.nome)
	}
	wg.Wait()
}

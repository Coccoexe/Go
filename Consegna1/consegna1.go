/*
Scrivete un programma che simuli una agenzia di viaggi che deve gestire le prenotazioni per due diversi viaggi da parte di 7 clienti.
Ogni cliente fa una prenotazione per un viaggio in una delle due mete disponibili (Spagna e Francia),
ognuna delle quali ha un numero minimo di partecipanti per essere confermata (rispettivamente 4 e 2).
	●Creare la struttura Cliente col relativo campo “nome”.
	●Creare la struttura Viaggio col rispettivo campo “meta”.
	●Creare la function prenota, che prende come input una persona e che prenota uno a caso dei due viaggi.
	●Creare una function stampaPartecipanti che alla fine del processo stampa quali viaggi sono confermati e quali persone vanno dove.
	●Ogni persona può prenotarsi al viaggio contemporaneamente.
	●Create tutte le classi e function che vi servono, ma mantenete la struttura data dalle due strutture e le due function che ho elencato sopra.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	nome string
}

type Viaggio struct {
	meta             string
	prenotazioni     []Cliente
	min_prenotazioni int
}

func prenota(c *Cliente, ch chan []Viaggio, wg *sync.WaitGroup) {
	v := <-ch
	choose := rand.Intn(len(v))
	v[choose].prenotazioni = append(v[choose].prenotazioni, *c)
	ch <- v
	wg.Done()
}

func stampaPartecipanti(v []Viaggio) {
	for j := range v {
		fmt.Println("Partecipanti prenotati per ", v[j].meta)
		for i := range v[j].prenotazioni {
			fmt.Println(i, ": ", v[j].prenotazioni[i].nome)
		}

		if v[j].min_prenotazioni > len(v[j].prenotazioni) {
			fmt.Println("-ATTENZIONE- viaggio annullato per numero di partecipanti insufficiente")
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	clienti := []Cliente{{"cliente1"}, {"cliente2"}, {"cliente3"}, {"cliente4"}, {"cliente5"}, {"cliente6"}, {"cliente7"}}
	sl := []Cliente{}
	v1 := Viaggio{"Spagna", sl, 4}
	v2 := Viaggio{"Francia", sl, 2}

	viaggi := []Viaggio{v1, v2}
	ch := make(chan []Viaggio, 1)
	ch <- viaggi

	var wg sync.WaitGroup

	wg.Add(len(clienti))
	for i := range clienti {
		go prenota(&clienti[i], ch, &wg)
	}
	wg.Wait()

	stampaPartecipanti(viaggi)

}

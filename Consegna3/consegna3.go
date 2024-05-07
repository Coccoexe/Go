/*
Scrivete un programma che simuli un lavoro fatto da tre operai, ognuno dei quali deve usare un martello,
un cacciavite e un trapano per fare un lavoro. Devono usare il cacciavite DOPO il trapano e il martello in un qualsiasi momento.
Ovviamente, possono fare solo un lavoro alla volta.
Gli attrezzi a disposizione sono: due trapani, un martello e un cacciavite,
quindi I tre operai devono aspettare di avere a disposizione gli attrezzi per usarli.
Modellate questa situazione minimizzando il più possibile le attese.●Creare la struttura Operaio col relativo campo “nome”.
●Creare la strutture Martello, Cacciavite e Trapano che devono essere “prese” dagli operai.
●Nelle function che creerete, inserite una stampa che dica quando l’operaio x ha preso l’oggetto y e quando ha finito di usarlo.
●Hint sulla logica: ogni operaio può avere solo un oggetto alla volta e ogni oggetto può essere in mano a un solo operaio.
●Per assicurarmi che ogni operaio abbia in mano un solo oggetto, posso mettere ogni operaio in un channel,
 e prima di cercare di prendere un oggetto...
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Operaio struct {
	nome  string
	doneM bool
	doneC bool
	doneT bool
}

type Martello struct {
	time int
}

type Cacciavite struct {
	time int
}

type Trapano struct {
	time int
}

func do(o Operaio, m chan Martello, c chan Cacciavite, t chan Trapano, wg *sync.WaitGroup) {
	for {
		if len(m) > 0 && o.doneM == false {
			mar := <-m
			fmt.Println(o.nome, " prende martello")
			time.Sleep(time.Duration(mar.time) * time.Second)
			fmt.Println(o.nome, " posa martello")
			m <- mar
			o.doneM = true
		}
		if len(c) > 0 && o.doneT == true && o.doneC == false {
			cac := <-c
			fmt.Println(o.nome, " prende cacciavite")
			time.Sleep(time.Duration(cac.time) * time.Second)
			fmt.Println(o.nome, " posa cacciavite")
			c <- cac
			o.doneC = true
		}
		if len(t) > 0 && o.doneT == false {
			tra := <-t
			fmt.Println(o.nome, " prende trapano")
			time.Sleep(time.Duration(tra.time) * time.Second)
			fmt.Println(o.nome, " posa trapano")
			t <- tra
			o.doneT = true
		}

		if o.doneC == true && o.doneM == true && o.doneT == true {
			break
		}
	}
	wg.Done()
}

func main() {
	martelli := make(chan Martello, 1)
	martelli <- Martello{2}

	cacciaviti := make(chan Cacciavite, 1)
	cacciaviti <- Cacciavite{3}

	trapani := make(chan Trapano, 2)
	trapani <- Trapano{1}
	trapani <- Trapano{1}

	operai := []Operaio{{"Operaio1", false, false, false}, {"Operaio2", false, false, false}, {"Operaio3", false, false, false}}

	var wg sync.WaitGroup

	wg.Add(len(operai))
	for i := range operai {
		go do(operai[i], martelli, cacciaviti, trapani, &wg)
	}
	wg.Wait()
}

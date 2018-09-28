package main

import (
	"github.com/go-loco/restful/rest"
)

func main() {

	var r *rest.Response

	go func() {
		r = <-rest.AsyncGet("https://api.mercadolibre.com/sites")
	}()

	//time.Sleep(1 * time.Millisecond)

	println(r.String())

	//ch1, ch2 := make(chan string), make(chan string)

	//rb1, rb2 := new(rest.RequestBuilder), new(rest.RequestBuilder)
	//rb2 := new(rest.RequestBuilder)
	/*
			go func() {
				/*h := make(http.Header)
				h.Add("X-Test", "test")

				rb1.Headers = h

				data := rb1.Get("https://api.mercadolibre.c/sites")
				println(data.String())
				ch1 <- "done"
			}()

			go func() {
				/*h := make(http.Header)
				h.Add("X-Test", "test")

				rb1.Headers = h

				rb2.Get("https://api.mercadolibre.com/sites")
				ch2 <- "done"
			}()

		var a, b, c *rest.FutureResponse

		ch := rest.ForkJoin(func(ct *rest.Concurrent) {
			a = ct.Get("https://api.mercadolibre.com/sites")
			b = ct.Get("https://api.mercadolibre.com/sites")
			c = ct.Get("https://api.mercadolibre.com/sites")
			println("FORK JOIN")
			print(a.Response())
		})

		for val := range ch {
			println(val.String())
		}

		//<-ch1
		//<-ch2
		println("DONE")
	*/
}

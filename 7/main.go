package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	test := or(
		sig(23*time.Second),
		sig(15*time.Second),
		sig(13*time.Second),
		sig(9*time.Second),
		sig(2*time.Second),
	)

	fmt.Print("lallafgpaginawgnwae")

	<-test

	fmt.Printf("fone after %v", time.Since(start))
}

//func or(channels ...<-chan interface{}) <-chan interface{} {
//	for {
//		for i := range channels {
//			select {
//			case <-channels[i]:
//				return channels[i]
//			default:
//				continue
//			}
//		}
//	}
//}

func or(channels ...<-chan interface{}) <-chan interface{} {
	test := make(chan interface{})
	go func() {
		for {
			for i := range channels {
				select {
				case <-channels[i]:
					close(test)
					return
				default:
					continue
				}
			}
		}
	}()
	return test
}

package gostatestore

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type readOp struct {
	key  string
	resp chan interface{}
}
type writeOp struct {
	key  string
	val  interface{}
	resp chan bool
}

var reads chan readOp
var writes chan writeOp
var wg sync.WaitGroup

func NewStore() {
	reads = make(chan readOp)
	writes = make(chan writeOp)
	Listener()
	wg.Wait()
}

func Listener() {
	go func() {
		fmt.Println("listening")
		var state = make(map[string]interface{})
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
}

func Reader(f interface{}) {
	e := reflect.ValueOf(f)
	if e.Kind() != reflect.Ptr {
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("reading")

		read := readOp{
			key:  e.Type().String(),
			resp: make(chan interface{})}
		reads <- read
		reply := <-read.resp
		////--- Extract Value without specifying Type
		// if &reply != nil {
		r := reflect.ValueOf(reply)
		if r.Kind() != reflect.Ptr {
			return
		}
		rn := reflect.Indirect(r)
		for i := 0; i < rn.NumField(); i++ {
			varName := rn.Type().Field(i).Name
			varType := rn.Type().Field(i).Type
			varValue := rn.Field(i).Interface()
			if !reflect.ValueOf(varValue).IsZero() {
				e.Elem().Set(reflect.ValueOf(reply).Elem())
			}
			fmt.Printf("read key: %s, struct: %v, struct field: %v struct field type: %v struct field value: %v\n", read.key, e.Type(), varName, varType, varValue)
		}
		// }
		time.Sleep(time.Millisecond)
	}()
}

func Writer(f interface{}) {
	e := reflect.ValueOf(f)
	if e.Kind() != reflect.Ptr {
		return
	}
	en := reflect.Indirect(e)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("writing")

		write := writeOp{
			key:  e.Type().String(),
			val:  f,
			resp: make(chan bool)}
		writes <- write
		<-write.resp
		for i := 0; i < en.NumField(); i++ {
			varName := en.Type().Field(i).Name
			varType := en.Type().Field(i).Type
			varValue := en.Field(i).Interface()
			fmt.Printf("write key: %s, struct: %v, struct field: %v struct field type: %v struct field value: %v\n", write.key, e.Type(), varName, varType, varValue)
		}
		time.Sleep(time.Millisecond)
	}()
}

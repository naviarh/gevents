# Gevents

**An implementation of the event model for Golang.**

## Overview

A purpose of this project is to create an event mechanism that allos to organize generation of events in different places of your program and to subscribe listeners to these events. The subscribers can receive necessary information from event generators.
This helps to organize an interaction between different parts of the program, for example, that are isolated in different goroutines.

This package has minimal complexity of the source code, and can be easily integrated into any program.

## In some cases, there is a need to use Gevents:

 **- Multiple listeners subscribed to the same event source.**
 
 **- The source does not know which listeners are subscribed.**
 
 **- Listeners can subscribe and unsubscribe dynamically.**

## Why I should use this:

 - Very easy to use event model.
 - No dependencies on other packages at all.
 - Full control over events and reactions to them.
 - Simple source code that is easy to upgrade.

## How to use:

 1. Install package (and update)

    ```sh
    go get -u github.com/naviarh/gevents
    ```

 2. Import Package

    ```go
    import "github.com/naviarh/gevents"
    ```

 3. Register the event globally

    ```go
    var event1 gevent.String
    ```

 4. Event channels initialization

    ```go
    func init() {
    	...
		event1.Init()
		...
	}
    ```

 5. The event occurs when a message is sent to OUT channel

    ```go
    event1.OUT <- "This is the event!"
    ```

 6. To add the event listener

    ```go
    event1.Add()
    ```

 7. To remove the event listener

    ```go
    event1.Del()
    ```

 8. The example of reactive code for event (by IN channel)

    ```go
    go func() {
		event1.Add()
		defer event1.Del()
		for {
			var := <-event1.IN
			fmt.Println(var)
		}
	}()
    ```

 9. One listener can subscribe to several events

    ```go
    go func() {
		event1.Add()
		event2.Add()
		defer event1.Del()
		defer event2.Del()
		for {
			select {
			case var1 := <-event1.IN:
				fmt.Println(var1)
			case var2 := <-event2.IN:
				fmt.Println(var2)
		}
	}()
    ```

 0. The following types of events channels exist

    ```go
    gevent.Interface	interface{}
	gevent.Bool		bool
	gevent.Int		int
	gevent.String		string
    ```

 1. During initialization, you can specify parameters

    ```go
    event1.Init(size_OUT, size_IN)
    // size_OUT - the size of the buffer of OUT channel, default is 100 (int)
    // size_IN - the size of the buffer of IN channel, default is 0 (int)
    ```

 2. You can add or remove several event listeners

    ```go
    event1.Add(N)
    event1.Del(N)
    // N - number (int)
    ```


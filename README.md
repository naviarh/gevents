# Gevents

**An implementation of the event model for Golang.**

## Overview

The purpose of this project is to create an event mechanism that allow you to organize the generation of events in different places of your program and sign consumers to these events, which can receive the necessary information from event generators. Thus, it is possible to organize the interaction between different parts of the program, for example, which are isolated in different goroutines.

This package has minimal complexity of the source code, and can be easily integrated into any program.

## In some cases, there is a need to use Gevents:

 **- Multiple consumers subscribed to the same event source.**
 
 **- The source does not know which consumers are subscribed.**
 
 **- Consumers can subscribe and unsubscribe dynamically.**

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
		event1.init()
		...
	}
    ```

 5. The event occurs when a message is sent to OUT channel

    ```go
    event1.OUT <- "This is the event!"
    ```

 6. To add the event listener

    ```go
    event1.Increment()
    ```

 7. To remove the event listener

    ```go
    event1.Decrement()
    ```

 8. The example of reactive code for event (by IN channel)

    ```go
    go func() {
		event1.Increment()
		defer event1.Decrement()
		for {
			var := <-event1.IN
			fmt.Println(var)
		}
	}()
    ```

 9. One consumer can listen to several events

    ```go
    go func() {
		event1.Increment(2)
		defer event1.Decrement(2)
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
    gevent.String	interface{}
	gevent.Bool	bool
	gevent.Int	int
	gevent.String	string
    ```


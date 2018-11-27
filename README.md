# events

**An implementation of the event model for Golang.**

## Why I should use this:

 - Very easy to use event model.
 - No dependencies on other packages.
 - Full control over events and reactions to them.
 - Simple source code that is easy to upgrade. 

## How to use:

 1. Install package

    ```sh
    go get github.com/.../events
    ```

 2. Import Package

    ```go
    import . "github.com/.../events"
    ```

 3. Register events globally

    ```go
    var (
    	event1 = E.AddEvent()
    	...
    )
    ```

 4. Add event-responsive function

    ```go
    func myFunc() {
        return true
    }
    ```

 5. Add a response to the event, we get the response ID and the result of the successful addition

    ```go
    num1, res := E.AddReaction(event1, myFunc)
    if res {
        println("Added")
    }
    ```

 6. The event occurs using the function

    ```go
    E.Event(event1)
    ```

 7. You can pause the response to an event

    ```go
    res = E.StopReaction(num1)
    if res {
        println("Suspended")
    }
    ```

 8. You can resume responding to an event

    ```go
    res = E.StartReaction(num1)
    if res {
        println("Resumed")
    }
    ```

 9. You can remove the response to the event

    ```go
    res = E.DelReaction(num1)
    if res {
        println("Removed")
    }
    ```

 1. You can delete the event itself (with all added responses to it)

    ```go
    res = E.DelEvent(event1)
    if res {
        println("Removed")
    }
    ```

 
 **Event response functions run in goroutines!**
  

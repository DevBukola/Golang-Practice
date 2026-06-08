### 1. What is a goroutine?

    Goroutine is a function that runs cocurrently (can overlap in execution). In this project, it allows the program to start checking all websites at almost the same time, and thereby making the runtime shorter.

### 2. Why does the main function not automatically wait for goroutines?

    The main function does not automatically wait for goroutines because goroutines run independently, and main() does not automatically "knows" it should wait for them.
    Once main() finishes, the program exits; it doesn't care if goroutines are still running.

    main() = the boss leaving the office.
    goroutines = the workers still doing tasks.
    When the boss leaves, the cmpany shuts down, even if workers ae still doing tasks.
    Example:
    go fmt.Println("Hello from goroutine.")
    fmt.Println("Hello from main.")

    - main() may finish first.
    - program exits immediately.
    - goroutine might never run.

    So, output is unreliable.
    Go behaves like this because it is desgned to be fast, lightweight, non-blocking. So,goroutines are not automatically synchronized with main(). The developer must explicitly control timng by making main() wait correctly:
    1. Using time.Sleep (this is not recommended for real apps. Okay? Okay!)
    2. Using sync.WaitGroup (Better approach! This is the best way and proper Go solution.)

    In summary, main() does NOT automatically wait because goroutines are independent, and main is just another goroutine with special power to exit the program.

### 3. What problem does `sync.WaitGroup` solves?

    Since the main() function does not automatically wait for goroutines, and the program exits once it finishes even if goroutines are still running, `sync.WaitGroup` helps control timing by making main() wait correctly for goroutines to finish before the program exits.

### 4. Why should 'Add' be called before starting goroutine?

    'Add' must be called first to correctly register work before exection begins.

### 5. Why should 'Done' be deferred inside the goroutines?

    Done() is deferred to ensre it always runs when the goroutine finishes regardless of how it exits.

### 6. What is a channel used for?

    A channel is used to send data from one goroutine to another. Goroutines do work concurrently. Chanels let them communicate and share data safely.

### 7. Who sends data into the result channel?

    The goroutine that has the result sends the data to the channel.

### 8. Who receives data from the result channel?

    The receiver is whichever goroutine reads from the channel. In most cases, that is the main goroutine.

### 9. Why should the result channel be closed?

    A channel does not always need to be closed. You close a channel to tell the receivers that no more values will e sent.

### 10. Why should the channel not be closed before goroutines finish?

    The channel not be closed before goroutines finish because a goroutine might still be trying to send data into the channel. If the goroutine runs after the channelis closed, Go will panic.

    panic: send on closed channel.

### 11. What causes a deadlock?

    A deadlock happens when goroutines are waiting for something that will never happen The program gets stuck because nobody can move forward.

    ```go
    exampleChan := make(chan string)
    message := <- exampleChan
    fmt.Println(message)
    ```
    - main waits to receive from exampleChan.
    - Nobody sends anything.
    Result:
    fatal error: all goroutines are asleep - deadlock!

### 12. What causes a panic when using channels?

    Channels panic when:
    1. Sending to a closed channel.

    ```go
    close(ch)
    ch <- "hello" //panic
    ```
    2. Closing a channel twice.

    ```go
    close(ch)
    close(ch) //panic
    ```
    3. Closing a nil channel.

    ```go
    var ch chan string
    close(ch) //panic
    ```

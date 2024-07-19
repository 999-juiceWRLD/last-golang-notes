# Nice Golang

## Goroutine

A goroutine is a lightweight thread managed by the Go runtime, created by `go` keyword and right before a function call.

## Channels

Channels are a typed conduit (tube, pipe) through which you can send and receive values with the channel operator, `<-`.

Like maps and slices, channels must be created before use. They also use the same `make` keyword:

```golang
ch := make(chan int)

// send data to a channel

ch <- 69

// receive data from a channel

n := <- ch
```

## Buffered Channels

Buffered channels are a type of channel in Go that allow you to store a limited number of values before they are received by a corresponding goroutine. Unlike unbuffered channels, which block both the sender and receiver until they synchronize, buffered channels provide a way to decouple sending and receiving operations to a certain extent.

```golang
ch := make(chan int, 3)  // Creates a buffered channel with a capacity of 3
```

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

A buffered channel allows the goroutine that is adding data to the buffered channel to keep running and doing things, even if the goroutines reading from the channel are starting to fall behind a little bit.

For example, you might have one goroutine that is receiving HTTP requests and you want it to be as fast as possible. However you also want it to queue up some background job, like sending an email, which could take a while. So the HTTP goroutine just parses the user's request and quickly adds the background job to the buffered channel. The other goroutines will process it when they have time. If you get a sudden surge in HTTP requests, the users will not notice any slowness in the HTTP if your buffer is big enough.

**Note:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note:** Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

Also, we can block and wait until something is sent on a channel using the following syntax

```golang
<- ch
```

This will block until it pops a single item off the channel, then continue, discarding the item.

## Select

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

It's designed for managing multiple communication operations in a concise and efficient manner. The `select` statement lets you wait on multiple channel operations and proceed with the one that becomes ready first.

A `select` statement is used to listen to multiple channels at the same time. It is similar to a `switch` statement but for channels.

```golang
select {
    case i, ok := <- chInts:
        fmt.Println(i)
    case s, ok := <- chStrings:
        fmt.Println(s)
}
```

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The `ok` variable in the example above refers to whether or not the channel has been closed by the sender yet.

## Mutexes

Stands for *mutual exclusion*, mutexes allow us to *lock* access to data. This ensures that we can control which goroutines can access certain data at which time.

Before jumping to mutex, it is important to understand the concept of critical section in concurrent programming. When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time. This section of code that modifies shared resources is called critical section.

**Note:** maps are **not** thread-safe, therefore they are not safe for concurrent use. If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.

The principle problem that mutexes help us avoid is the *concurrent read/write problem*. This problem arises when one thread is writing to a variable while another thread is reading from that same variable *at the same time*. When this happens, a Go program will panic because the reader could be reading bad data while it's being mutated in place.

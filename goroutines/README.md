# goroutines


## What are Goroutines?

Goroutines are functions or methods that run concurrently with other functions or methods.

## Advantages of Goroutines over threads

* Goroutines are extremely cheap when compared to threads.
* The Goroutines are multiplexed to fewer number of OS threads. 
* Goroutines communicate using channels. 

## How to start a Goroutine?

* When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
* The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
# panicrecover

## When should panic be used?

* One important factor is that you should avoid panic and recover and use errors where ever possible. Only in cases where the program just cannot continue execution should a panic and recover mechanism be used.

There are two valid use cases for panic.

1. An unrecoverable error where the program cannot simply continue its execution. 
2. A programmer error. 

## What is recover?

* recover is a builtin function which is used to regain control of a panicking goroutine.

## Panic, Recover and Goroutines

* Recover works only when it is called from the same goroutine. It's not possible to recover from a panic that has happened in a different goroutine. 

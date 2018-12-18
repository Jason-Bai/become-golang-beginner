# Defer

## What is Defer?

* Defer statement is used to execute a function call just before the function where the defer statement is present returns.

## Arguments evaluation

* The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.

## Stack of defers

* When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.
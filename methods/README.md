# methods

## Value receivers in methods vs value arguments in functions

* When a function has a value argument, it will accept only a value argument.
* When a method has a value receiver, it will accept both pointer and value receivers.

## Pointer receivers in methods vs pointer arguments in functions.

* Similar to value arguments, functions with pointer arguments will accept only pointers whereas methods with pointer receivers will accept both value and pointer receiver.

## Methods on non struct types

* To define a method on a type, the definition of the receiver type of the method and the definition of the method should be in the same package.
# types

The following are the basic types available in go:

* bool
* Numeric Types
  * int8, int16, int32, int64, int
  * uint8, uint16, uint32, uint64, uint
  * float32, float64
  * complex64, complex128
  * byte
  * rune
* string

## Signed integers

int8: represents 8 bit signed integers 
size: 8 bits 
range: -128 to 127

int16: represents 16 bit signed integers 
size: 16 bits 
range: -32768 to 32767

int32: represents 32 bit signed integers 
size: 32 bits 
range: -2147483648 to 2147483647

int64: represents 64 bit signed integers 
size: 64 bits 
range: -9223372036854775808 to 9223372036854775807

int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer. 
size: 32 bits in 32 bit systems and 64 bit in 64 bit systems. 
range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

## Unsigned integers

uint8: represents 8 bit unsigned integers 
size: 8 bits 
range: 0 to 255

uint16: represents 16 bit unsigned integers 
size: 16 bits 
range: 0 to 65535

uint32: represents 32 bit unsigned integers 
size: 32 bits 
range: 0 to 4294967295

uint64: represents 64 bit unsigned integers 
size: 64 bits 
range: 0 to 18446744073709551615

uint : represents 32 or 64 bit unsigned integers depending on the underlying platform. 
size : 32 bits in 32 bit systems and 64 bits in 64 bit systems. 
range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

## Floating point types

float32: 32 bit floating point numbers 
float64: 64 bit floating point numbers

## Other numeric types

byte is an alias of uint8 
rune is an alias of int32
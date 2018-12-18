# interface ii

## Implementing interfaces using pointer receivers vs value receivers

* The reason is that it is legal to call a pointer-valued method on anything that is already a pointer or whose address can be taken. The concrete value stored in an interface is not addressable and hence it is not possible for the compiler to automatically take the address of a in line no. 45 and hence this code fails.
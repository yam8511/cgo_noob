#include <stdint.h> // for uintptr_t
#include <stdio.h>

// A Go function
extern char *MyGoPrint(uintptr_t handle);
extern char *Concat(char *a, char *b);

// A C function
void myprint(uintptr_t handle)
{
    printf("%s\n", Concat("c call", MyGoPrint(handle)));
}

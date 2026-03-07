#include <string.h>

#define c_Bytes(ptr, n) \
    ((ptr) ? unsafe_Slice(ptr, n) : (so_Slice){NULL, 0, 0})

#define c_String(ptr) \
    ((ptr) ? unsafe_String(ptr, strlen((const char*)(ptr))) : (so_String){NULL, 0})

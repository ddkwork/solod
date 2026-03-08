#include <string.h>
#include "so/builtin/builtin.h"

static inline so_Slice c_Bytes(void* ptr, size_t n) {
    return ptr ? unsafe_Slice(ptr, n) : (so_Slice){NULL, 0, 0};
}

static inline so_String c_String(void* ptr) {
    char* s = (char*)(ptr);
    return ptr ? unsafe_String(s, strlen(s)) : (so_String){NULL, 0};
}

static inline char* c_CharPtr(void* ptr) {
    return (char*)ptr;
}

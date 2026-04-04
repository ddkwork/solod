#include <assert.h>
#include <string.h>
#include "so/builtin/builtin.h"

#define c_Alignof(T) ((so_int)alignof(T))

static inline void c_Assert(bool cond, const char* msg) {
    assert((cond) && msg);
}

static inline so_Slice c_Bytes(void* ptr, size_t n) {
    return ptr ? (so_Slice){ptr, n, n} : (so_Slice){&so_Nil, 0, 0};
}

static inline char* c_CharPtr(void* ptr) {
    return (char*)ptr;
}

#define c_Sizeof(T) ((so_int)sizeof(T))

#define c_Slice(T, ptr, len, cap) \
    (ptr ? (so_Slice){ptr, (size_t)len, (size_t)cap} : (so_Slice){&so_Nil, 0, 0})

static inline so_String c_String(void* ptr) {
    char* s = (char*)(ptr);
    return ptr ? (so_String){s, strlen(s)} : (so_String){(char*)&so_Nil, 0};
}

static inline void* c_PtrAdd(void* ptr, so_int offset) {
    return (char*)ptr + (size_t)offset;
}

#define c_PtrAs(T, ptr) ((T*)(ptr))

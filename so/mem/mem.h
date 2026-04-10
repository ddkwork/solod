#include <assert.h>
#include <string.h>
#include "so/builtin/builtin.h"

// SwapByte swaps n bytes between a and b.
// Panics if either a or b is nil.
//
// SwapByte temporarily allocates a buffer of size n
// on the stack, so it's not suitable for large n.
static inline void mem_SwapByte(void* a, void* b, so_int n) {
    assert(a != NULL && "mem: nil pointer");
    assert(b != NULL && "mem: nil pointer");
    assert(n >= 0 && "mem: negative size");
    if (n == 0) return;

    size_t size = (size_t)n;
    char tmp[size];
    memcpy(tmp, a, size);
    memcpy(a, b, size);
    memcpy(b, tmp, size);
}

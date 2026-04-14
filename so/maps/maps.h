// keyHash hashes a key, dispatching to string or inline hash.
#define maps_keyHash(K, key_ptr, seed) _Generic((K){0}, \
    so_String: maps_hashString(key_ptr, seed),          \
    default: maps_hash(key_ptr, sizeof(K), seed))

// equal compares two typed key pointers for equality.
#define maps_keyEqual(K, a, b)                                       \
    _Generic((K){0},                                                 \
        so_String: so_string_eq(*(so_String*)(a), *(so_String*)(b)), \
        default: memcmp((a), (b), sizeof(K)) == 0)

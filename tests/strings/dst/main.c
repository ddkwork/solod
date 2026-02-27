#include "main.h"

int main(void) {
    so_String str = so_strlit("Hi 世界!");
    for (so_int i = 0; i < so_len(str); i++) {
        uint8_t chr = so_index(str, uint8_t, i);
        so_println("%s %lld %s %u", "i =", i, "chr =", chr);
    }
    for (so_int i = 0; i < so_len(str);) {
        int _iw = 0;
        so_rune r = so_utf8_decode(str, i, &_iw);
        so_println("%s %lld %s %d", "i =", i, "r =", r);
        i += _iw;
    }
    for (so_int i = 0; i < so_len(str);) {
        int _iw = 0;
        so_utf8_decode(str, i, &_iw);
        so_println("%s %lld", "i =", i);
        i += _iw;
    }
    for (so_int _ = 0; _ < so_len(str);) {
        int __w = 0;
        so_rune r = so_utf8_decode(str, _, &__w);
        so_println("%s %d", "r =", r);
        _ += __w;
    }
    so_String s1 = so_strlit("hello");
    so_String s2 = so_strlit("world");
    if (so_string_eq(s1, s2) || so_string_eq(s1, so_strlit("hello"))) {
        so_println("%s", "ok");
    }
}

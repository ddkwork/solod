#include "main.h"

int main(void) {
    so_String a = so_strlit("initial");
    so_println("%s", a.ptr);
    so_int b = 1;
    so_int c = 2;
    so_println("%lld %lld", b, c);
    bool d = true;
    (void)d;
    so_int e = 0;
    so_println("%lld", e);
    so_String f = so_strlit("apple");
    so_println("%s", f.ptr);
}

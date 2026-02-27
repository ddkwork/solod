#include "main.h"

so_int main_Rect_Area(void* self) {
    main_Rect* r = (main_Rect*)self;
    return r->width * r->height;
}

so_int main_Rect_Perim(void* self, so_int n) {
    main_Rect* r = (main_Rect*)self;
    return n * (2 * r->width + 2 * r->height);
}

static so_int calc(main_Shape s) {
    return s.Perim(s.self, 2) + s.Area(s.self);
}

static bool isRect(main_Shape s) {
    bool ok = (s.Area == main_Rect_Area);
    return ok;
}

static so_int asRect(main_Shape s) {
    bool ok = (s.Area == main_Rect_Area);
    if (!ok) {
        return 0;
    }
    main_Rect r = *((main_Rect*)s.self);
    return main_Rect_Area(&r);
}

int main(void) {
    main_Rect r = {.width = 10, .height = 5};
    main_Shape s = (main_Shape){.self = &r, .Area = main_Rect_Area, .Perim = main_Rect_Perim};
    calc(s);
    calc((main_Shape){.self = &r, .Area = main_Rect_Area, .Perim = main_Rect_Perim});
    calc((main_Shape){.self = &r, .Area = main_Rect_Area, .Perim = main_Rect_Perim});
    (void)isRect(s);
    so_int a = asRect(s);
    (void)a;
}

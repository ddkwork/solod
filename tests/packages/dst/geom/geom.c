#include "geom.h"
const double geom_Pi = 3.14159;

static double rectArea(double width, double height) {
    return width * height;
}

double geom_RectArea(double width, double height) {
    return rectArea(width, height);
}

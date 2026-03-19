#include "main.h"

// -- Implementation --

int main(void) {
    {
        // time.Date and time.Time properties.
        time_Time t = time_Date(2021, time_May, 10, 12, 33, 44, 777888999, time_UTC);
        if (time_Time_Year(t) != 2021) {
            so_panic("unexpected Time.Year");
        }
        if (time_Time_Month(t) != time_May) {
            so_panic("unexpected Time.Month");
        }
        if (time_Time_Day(t) != 10) {
            so_panic("unexpected Time.Day");
        }
        if (time_Time_Hour(t) != 12) {
            so_panic("unexpected Time.Hour");
        }
        if (time_Time_Minute(t) != 33) {
            so_panic("unexpected Time.Minute");
        }
        if (time_Time_Second(t) != 44) {
            so_panic("unexpected Time.Second");
        }
        if (time_Time_Nanosecond(t) != 777888999) {
            so_panic("unexpected Time.Nanosecond");
        }
        so_println("%" PRId64 " %" PRId64 " %" PRId64 " %" PRId64 " %" PRId64 " %" PRId64 " %" PRId64, time_Time_Year(t), time_Time_Month(t), time_Time_Day(t), time_Time_Hour(t), time_Time_Minute(t), time_Time_Second(t), time_Time_Nanosecond(t));
    }
}

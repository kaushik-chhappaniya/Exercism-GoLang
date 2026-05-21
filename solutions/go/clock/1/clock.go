package clock

import "fmt"

type Clock struct {
    minutes int
}

const day = 24 * 60

func normalize(minutes int) int {
    minutes %= day
    if minutes < 0 {
        minutes += day
    }
    return minutes
}

func New(hour, minute int) Clock {
    return Clock{minutes: normalize(hour*60 + minute)}
}

func (c Clock) Add(minutes int) Clock {
    return Clock{minutes: normalize(c.minutes + minutes)}
}

func (c Clock) Subtract(minutes int) Clock {
    return Clock{minutes: normalize(c.minutes - minutes)}
}

func (c Clock) String() string {
    h := c.minutes / 60
    m := c.minutes % 60
    return fmt.Sprintf("%02d:%02d", h, m)
}
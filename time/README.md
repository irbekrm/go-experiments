# time

## time.Now()
- returns a `time.Time` object that includes both wall clock and monotonic clock readings
- `Time.Round` can be used to strip off the monotonic clock reading and round the wall clock reading

## time.Parse()
- takes a layout and a string representing time and returns `time.Time`
- common layouts are defined in the time package i.e `time.RFC3339`



## Useful resources
- CLOCK_REALTIME (wall clock) vs CLOCK_MONOTONIC https://stackoverflow.com/questions/3523442/difference-between-clock-realtime-and-clock-monotonic
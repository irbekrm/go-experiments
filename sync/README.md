# sync

## sync.Pool

Implements a [free list](https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Freelist.html)

Can be used to avoid allocating memory anew for an object inside a function that is called multiple times (and potentially in parallel). Instead of creating a new object each time, pick one from pool (will create new if pool is empty- specify how to with sync.Pool.New) and put back after using it.

Run benchmark tests `go -bench=. -benchmem` to see the difference between `PrintWithPool` and `PrintWithoutPool`

## Useful resources:

- Official docs https://golang.org/pkg/sync/#Pool
- justforfunc tutorial https://www.youtube.com/watch?v=rfXSrgIGrKo&t=720s&ab_channel=justforfunc%3AProgramminginGo
- about free lists https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Freelist.html
- great blog post about `sync` package https://medium.com/@teivah/a-closer-look-at-go-sync-package-9f4e4a28c35a
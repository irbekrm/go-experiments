# pprof

## Examples from https://blog.golang.org/pprof

### With Linux `time` util

Run a script that measures real time, system time, user time & memory usage of a simple go program using `time` util

`docker run -it -v $(pwd):/tmp/app -w /tmp/app golang:alpine ./run.sh`

### Write cpu profile to a file at runtime, interpret with `go tool pprof`

`pprof.StartCPUProfile(<filepath>)` has been added to code

`go build -o rsa main.go`

`./rsa --cpuprofile cpu.prof --bits 4096`

`go tool pprof cpu.prof` and:

* `web` command will generate a graph that can be seen in browser if [Graphviz](http://www.graphviz.org/about/) is 

* `topN` show the biggest CPU consumers

* `list <regexp>` show how much cpu time spent on each line of functions matching regexp- really useful

### Write memory profile to a file at runtime, interpret with `go tool pprof`

* `pprof.WriteHeapProfile(<filename>)` has  been added to code

* `pprof.WriteHeapProfile` should be  ran _after_ the code runs https://golang.org/pkg/runtime/pprof/#hdr-Profiling_a_Go_program

* `./rsa  --memprofile mem.prof --bits 4096`

* `go tool pprof mem.prof` 

* `topN`, `web`, `list <regexp>`

* `sample_index` set before running commands- allocated objects/space or inuse objects/space *useful*
i.e `sample_index=alloc_space`
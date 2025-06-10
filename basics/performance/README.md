benchmark:
```shell
  go test -bench . bench_test.go
```

benchmark with memory:
```shell
    go test -bench . -benchmem bench_test.go
```

out files for pprof:
```shell
  go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 bench_test.go
```

pprof cpu:
```shell
    go tool pprof performance.test cpu.out
```

pprof mem:
```shell
    go tool pprof performance.test mem.out
```

pprof commands:
- top - самые долгие методы
- list Unmarshal - сколько времени занимали операции внутри метода
- web - построить граф

для mem:
- alloc_objects -> top - где выделялись объекты

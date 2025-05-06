### Small Files - Compare Speeds By Base 10 Orders of Magnitude

Use the filegen util to create random files:
e.g.

```bash
# assuming pwd is filegen
 go run main.go -dir ~/filegen_out/10k -max 5000 -min 100 -num 10000
 go run main.go -dir ~/filegen_out/100k -max 5000 -min 100 -num 100000
 go run main.go -dir ~/filegen_out/1M -max 5000 -min 100 -num 1000000
```

And use the

```bash
# assuming pwd is serial-full-read
# using go run
go run main.go -path ~/filegen_out/10k
# build and run
go build -o checker
./checker -path ~/filegen_out/10k
```

I havent built in a way to scale... yes its just a for loop, but I am unsure of the different ways I might take the serial-full-read
so I've decided to keep the interface simple and keep the extra logic elsewhere... such as in shell scripts `base10_orders_of_magnitude.sh`

#### Results Summary

| Files | Write Time (filegen) | Read Time |
| ----- | -------------------- | --------- |
| 10K   | 0.8s                 | ~0.2s     |
| 100K  | 5s                   | ~2s       |
| 1M    | 45s                  | ~13-17s   |

#### pprof top

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
```

flat flat% sum% cum cum%
1950ms 15.29% 15.29% 1950ms 15.29% internal/runtime/syscall.Syscall6
870ms 6.82% 22.12% 870ms 6.82% runtime.(\*mspan).base (inline)
560ms 4.39% 26.51% 560ms 4.39% indexbytebody
550ms 4.31% 30.82% 2230ms 17.49% runtime.mallocgc
350ms 2.75% 33.57% 2310ms 18.12% runtime.scanobject
310ms 2.43% 36.00% 310ms 2.43% runtime.nextFreeFast (inline)
290ms 2.27% 38.27% 490ms 3.84% runtime.findObject
270ms 2.12% 40.39% 270ms 2.12% runtime.duffcopy
270ms 2.12% 42.51% 1370ms 10.75% runtime.greyobject
190ms 1.49% 44.00% 190ms 1.49% runtime.memclrNoHeapPointers

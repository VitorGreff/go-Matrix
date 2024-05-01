## Go Concurrency
This code aims only to explore golang concurrency system. The idea is to perform a matrix multiplication using two different functions, one sequential and the other concurrent.

## How to run
```bash
go run cmd.go -r <rows-number> -c <columns-number> -t <threads-number>
```
All flags are optional, and its default values can be seen by using -h flag.
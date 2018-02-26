# hw3_bench

### Запуск:
   
   ```go test -v ```- чтобы проверить что ничего не сломалось
   
   ```go test -bench . -benchmem ``` - для просмотра производительности
   
   
   
#### До момента ускорения   
   ```
   BenchmarkSlow-4                5         227807840 ns/op        336854937 B/op    284170 allocs/op
   BenchmarkFast-4               10         181633730 ns/op        336854842 B/op    284167 allocs/op
   ```
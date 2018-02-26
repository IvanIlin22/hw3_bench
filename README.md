# hw3_bench

### Запуск:
   
   ```go test -v ```- чтобы проверить что ничего не сломалось
   
   ```go test -bench . -benchmem ``` - для просмотра производительности
   
   
   
#### До момента ускорения   
   ```
   BenchmarkSlow-4                5         227807840 ns/op        336854937 B/op    284170 allocs/op
   BenchmarkFast-4               10         181633730 ns/op        336854842 B/op    284167 allocs/op
   ```
   
##### Стадии ускорения

1. После замены стандартного unmarshal на easyjson от mail.ru (кодогенерация), избавился от двух циклов

```
BenchmarkFast-4               10         115124640 ns/op        335701931 B/op    239236 allocs/op
```

2. Замена использования регулярных выраженией на strings.Contains (поиск строки в строке) 

```
BenchmarkFast-4              300           5124503 ns/op         4386214 B/op      14843 allocs/op
```
 
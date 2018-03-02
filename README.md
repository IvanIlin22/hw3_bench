# hw3_bench

Есть функиця, которая что-то там ищет по файлу. Но делает она это не очень быстро. Надо её оптимизировать.

В данном случае к чему прийти нужно:

B/op < 559910

allocs/op < 12506

### Снятие тестов

```
go test -bench . -benchmem -cpuprofile cpu.out -memprofile mem.out -memprofilerate=1 main_test.go fast.go common.go fast_easyjson.go
```

Исследование Оперативной памяти
```
go tool pprof .\main.test.exe .\mem.out
```

### Запуск:
   
   ```go test -v ```- чтобы проверить что ничего не сломалось
   
   ```go test -bench . -benchmem ``` - для просмотра производительности
   
### Определения показателей

   1. Кол-во выполненных итераций
   2. ns/op - сколько занимает одна операция (нано секунд)
   3. B/op - кол-во байт на операцию
   4. allocs/op - Кол-во аллокаций памяти
   
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
3. После установки sync.Pool - повторного использования памяти

```
BenchmarkFast-4              300           4271814 ns/op         4150028 B/op      12623 allocs/op
``` 

4. Построчное считывание файла
```
BenchmarkFast-4               50          25690114 ns/op         2415471 B/op       9621 allocs/op
```

5. Финальная оптимизация, отсек лишние преобразования

```
bytes.Contains(row, []byte("Android"))
```

```
BenchmarkFast-4              500           3443097 ns/op          375658 B/op       4658 allocs/op
```
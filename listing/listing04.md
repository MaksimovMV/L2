Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
fatal error: all goroutines are asleep - deadlock!
 
 Deadlock возникает, когда группа goroutines ждет друг друга, и ни одна из них не может продолжить. 
 Каналы необходимо закрывать, range не читает с закрытых каналов

```

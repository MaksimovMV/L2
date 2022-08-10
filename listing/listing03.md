Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil> - Пакет fmt форматирует значение ошибки, вызывая его строковый метод Error()
false - error представляет из себя интерфейс с методом Error() string, интерфейс равен
nil только в случае если его тип и значения полей равны нулю

```

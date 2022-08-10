package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func cd(a string) error {
	err := os.Chdir(a)
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(wd)
	return nil
}

func scan() {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() && s.Text() != "quit" {
		ss := strings.Split(s.Text(), "|")
		var wg sync.WaitGroup
		wg.Add(len(ss))
		for _, a := range ss {
			go func(a string) {
				defer wg.Done()
				args := strings.Fields(a)
				if len(args) < 1 {
					return
				}
				switch args[0] {
				case "cd":
					if len(args) < 2 {
						fmt.Println("too few arguments")
						return
					}
					err := cd(args[1])
					if err != nil {
						fmt.Println(err)
						return
					}
				case "pwd":
					wd, err := os.Getwd()
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(wd)
				case "echo":
					fmt.Println(strings.Join(args[1:], " "))
				case "kill":
					if len(args) < 2 {
						fmt.Println("too few arguments")
						return
					}
					pid, err := strconv.Atoi(args[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					proc, err := os.FindProcess(pid)
					if err != nil {
						fmt.Println(err)
						return
					}
					err = proc.Kill()
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("Process killed")
				case "ps":
					c := exec.Command("TASKLIST")
					c.Stdin = os.Stdin
					c.Stdout = os.Stdout
					c.Stderr = os.Stderr
					c.Run()
				case "fork":
					var procAttr os.ProcAttr
					procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}
					_, err := os.StartProcess(args[1], args[1:], &procAttr)
					if err != nil {
						fmt.Println(err)
						return
					}
				case "exec":
					programPath, err := exec.LookPath(args[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					if err := syscall.Exec(programPath, args[1:], os.Environ()); err != nil {
						fmt.Println(err)
						return
					}
				default:
					fmt.Printf("%v is unknown command\n", args[0])
					return
				}
			}(a)
		}
		wg.Wait()

		fmt.Print("> ")
	}
}

func main() {

	wd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(wd)
	fmt.Print("> ")

	scan()
}

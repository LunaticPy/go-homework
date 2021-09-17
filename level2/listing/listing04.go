package main

func main() {
	ch := make(chan int)
	go func() {
		// пишем 10 значений
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// читаем 10 значений
	// уходим в блокировку, тк 11е мы никогда не получим
	for n := range ch {
		println(n)
	}
}

// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /home/lunatic/work/level2/24.go:11 +0xa7
// exit status 2

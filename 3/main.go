package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	cnt := len(slice)
	ans := 0

	//создаём счетчик горутин, т.к. мы знаем кол-во элементов в слайсе
	var wg sync.WaitGroup
	wg.Add(cnt)

	//создаём мьютекс для предотвращения гонки
	var mx sync.Mutex

	for _, v := range slice {
		go func(n int) {
			defer wg.Done()
			mx.Lock()
			ans += n * n
			mx.Unlock()
		}(v)
	}
	wg.Wait()

	fmt.Println(ans)
}

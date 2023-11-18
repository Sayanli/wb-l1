package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	cnt := len(slice)
	ans := make([]int, 0, cnt)

	//создаем счетчик горутин, т.к. мы знаем кол-во элементов в слайсе
	var wg sync.WaitGroup
	wg.Add(cnt)

	//создаем мьютекс для предотвращения гонки
	var mx sync.Mutex

	for _, v := range slice {
		go func(v int) {
			defer wg.Done()
			mx.Lock()
			ans = append(ans, v*v)
			mx.Unlock()
		}(v)
	}
	wg.Wait()

	// т.к. горутины завершатся в разное время, ответ будет неупорядоченным
	fmt.Println(ans)
}

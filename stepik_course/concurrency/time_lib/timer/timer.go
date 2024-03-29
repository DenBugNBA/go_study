package main

import "time"

func main() {
	t := time.NewTimer(time.Second) // создаем новый таймер, который сработает через 1 секунду
	go func() {
		<-t.C // C - канал, который должен вернуть значение через заданное время
	}()
	t.Stop() // но мы можем остановить таймер и раньше установленного времени

	t.Reset(time.Second * 2) // пока таймер не сработал, мы можем сбросить его, установив новый срок выполнения
	<-t.C
}

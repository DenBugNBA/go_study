package main

import "fmt"

func main() {
	v := 5
	// Переменную v инициализировали со значением 5. Теперь там лежит 5.
	p := &v
	// Нашли адрес в памяти для переменной v, и записали его в переменную р.
	// Теперь р - это указатель, и там лежит адрес v. Адрес области памяти выглядит примерно так: 0х000012.
	// Переменная р не знает, что лежит в v, но знает ее адрес. Легко запомнить: амперсанд напоминает слово "адрес".
	// Пэ инициализируем как адрес вэ.
	fmt.Print(*p, " ")
	// Звездочку просто запомнить как глазок в двери, когда мы приходим по адресу типа 0х000012 и берем то,
	// что там лежит. А лежит там v, все с тем же значением. Печатаем его, потом пробел.
	changePointer(p)
	// Берем копию р (там лежит какой-то адрес по-прежнему) и передаем его в функцию.
	// Го в свои функции передает всегда копию переменной, а оригинал остается в мейн, это важно.
	// Идем и смотрим что происходит в функции...
	fmt.Print(*p)
	// Функция нам ничего не вернула, мы пришли ни с чем, поэтому пошли по старому адресу р,
	// увидели глазок, открыли дверь и взяли там то же самое, что и в прошлый раз.
}

func changePointer(p *int) {
	// Функция принимает копию адреса, но ничего не возвращает. Здесь звездочка путает нас тем, что стоит перед типом.
	// Но мы все равно представляем глазок. Мы в функции, стоим перед дверью по адресу р,
	// нам осталось открыть ее, и взять то что там лежит (а там по-прежнему v с неизменным значением).
	v := 3
	// Но что это? Мы не открываем дверь, а инициализируем абсолютно новую переменную v в этой функции.
	// В main про эту v ничего не известно. Это совсем другой Сережа (v).
	p = &v
	// А вот тут интересно! Мы берем эту р (копию той р, которая осталась в мейне), и говорим:
	// взять адрес v и положить его в р. У нас здесь другая v, значит и адрес в р будет другим.
	// Но кто об этом в курсе? В функцию мейн ничего не передается, ("все осталось в Вегасе"). Функция завершает работу.
}

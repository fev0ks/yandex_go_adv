package main

import (
	"fmt"
)

func main() {
	number := 17
	mask := 0
	c := 0
	for div := number; div >= 2; {
		div = div / 2
		c++
		mask = mask | 1<<c
	}
	newNumber := number | mask
	fmt.Println()
	fmt.Printf("number %d = %b\n", number, number)
	fmt.Printf("result %d = %b\n", newNumber, newNumber)

	//context_topick.CtxTask1()

	//ctx1 := context.WithValue(context.Background(), "key_1", "value_1")
	//ctx2 := context.WithValue(ctx1, "key_2", "value_2")
	//fmt.Println(ctx2.Value("key_1"))

	//currentTimeStr := "2021-09-19T15:59:41+03:00"
	//// скопируйте блок себе в IDE и допишите код
	//layout := time.RFC3339 //"2006-01-02T15:04:05Z07:00"
	//currentTime, err := time.Parse(layout, currentTimeStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(currentTime)
	//
	//currentTimeStr = "2021-09-19T15:59:41+03:00"
	//// скопируйте блок себе в IDE и допишите код
	////layout = "..."
	//currentTime, err = time.Parse(layout, currentTimeStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(currentTime)
	//
	//now := time.Now()
	//
	//fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	//fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
	//fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))
	//
	//truncTime := now.Truncate(24 * time.Hour)
	//fmt.Println(now)
	//fmt.Println(truncTime)
	//fmt.Println()
	//fmt.Println("Андрей родился 26 ноября 1993 года. Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.")
	//// скопируйте блок себе в IDE и допишите код
	//birthday := time.Date(1993, 11, 26, 0, 0, 0, 0, time.Local)
	//hundred := birthday.Add(100 * 365 * 24 * time.Hour)
	//days := int(hundred.Sub(now).Hours() / 24)
	//days2 := int(time.Until(hundred).Hours() / 24)
	//fmt.Println(days)
	//fmt.Println(days2)
	//
	////c := make(chan byte)
	//ticker := time.NewTicker(2 * time.Second)
	//
	//number := 0
	//
	//for number < 10 {
	//	<-ticker.C
	//	number++
	//	fmt.Println(int(now.Sub(time.Now()).Seconds()))
	//}
	//ticker.Stop()
}

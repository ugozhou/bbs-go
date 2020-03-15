package main

import (
	"bbs-go/model"
	"github.com/mlogclub/simple"
	"math/rand"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func main() {
	simple.Generate("./", "bbs-go", simple.GetGenerateStruct(&model.Waiter{}))
	simple.Generate("./", "bbs-go", simple.GetGenerateStruct(&model.Notice{}))
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(GetRandomString(10))
	//}
	//fmt.Println(model.Recharge)
	//fmt.Println(model.BuyMiner)

}

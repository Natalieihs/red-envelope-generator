package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RedEnvelope struct {
	Id     int
	Amount float64
}

func GenerateRedEnvelopes(totalAmount float64, totalCount int) []RedEnvelope {
	envelopes := make([]RedEnvelope, totalCount)

	rand.Seed(time.Now().UnixNano())

	// 计算每个红包的平均金额
	average := totalAmount / float64(totalCount)

	// 计算每个红包可以分配的最大金额
	maxAmount := 2 * average

	// 计算剩余金额
	remainingAmount := totalAmount

	// 计算剩余红包数量
	remainingCount := totalCount

	// 生成红包
	for i := 0; i < totalCount-1; i++ {
		// 计算当前红包的最大金额
		maxCurrent := remainingAmount - float64(remainingCount-1)*average

		// 随机生成红包金额
		amount := rand.Float64() * maxAmount

		// 如果生成的金额大于当前红包的最大金额，则使用当前红包的最大金额
		if amount > maxCurrent {
			amount = maxCurrent
		}

		// 将随机生成的金额加入到当前红包中
		envelopes[i].Id = i + 1
		envelopes[i].Amount = amount

		// 更新剩余金额和剩余红包数量
		remainingAmount -= amount
		remainingCount--
		maxAmount = 2 * remainingAmount / float64(remainingCount)
	}

	// 将剩余金额分配到最后一个红包中
	envelopes[totalCount-1].Id = totalCount
	envelopes[totalCount-1].Amount = remainingAmount

	return envelopes
}
func main() {
	envelopes := GenerateRedEnvelopes(100, 30)
	//切片中的数字保留两位小数  并且计算总和
	var sum float64
	for i := 0; i < len(envelopes); i++ {
		//envelopes[i].Amount = float64(int(envelopes[i].Amount*100)) / 100
		sum += envelopes[i].Amount
	}
	fmt.Println(sum)
	fmt.Println(envelopes)
}

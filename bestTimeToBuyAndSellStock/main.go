package main

func main() {

}

func maxProfit(prices []int) int {
	left := 0
	right := 1
	maxMoney := 0
	for right < len(prices) {
		if prices[left] < prices[right] {
			maxMoney = max(maxMoney, prices[right]-prices[left])
		} else {
			left = right
		}
		right++
	}

	return maxMoney
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

package main

func main() {
}

const mod int = 1000000007

func countRoutes(ls []int, start, end, fuel int) {
	n := len(ls)
	// 初始化缓存器
	// 之所以初始化为 -1
	// 是为了区分「某个状态下路径数量为0」和「某个状态尚未计算过两种情况」
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		fu := make([]int, fuel+1)
		for i := range fu {
			fu[i] = -1
		}
	}
}

package lesson6

/*给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/
//MinimumTotal: 三角形最小路径和
/*parameter
triangle: 形如[[2],[3,4],[6,5,7],[4,1,8,3]]的二维整数数组
return: 最小路径和
*/
func MinimumTotal(triangle [][]int) int {

	layers := len(triangle)
	for i := layers - 2; i >= 0; i-- {

		layers--
		for j := 0; j < layers; j++ {

			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func min(x int, y int) int {

	if x < y {

		return x
	}

	return y
}

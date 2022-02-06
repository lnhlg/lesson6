package lesson6

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//climbStairs: 计算上n阶台阶的方法数
/*parameter
n: 阶数
return: 上台阶的总方法数
*/
func ClimbStairs(n int) int {

	//0阶只有一种方法
	if n == 0 {

		return 1
	}

	//小于3阶方法数为阶数
	if n < 3 {

		return n
	}

	x, y, z := 0, 0, 1
	for i := 0; i < n; i++ {

		x = y
		y = z
		z = x + y
	}

	return z
}

package main

func isOK(x, y, row, col int) bool {
	if (x >= 0 && x < row && y >= 0 && y < col) {
		return true
	}
	return false
}

func dfs(grid [][]int, x int, y int) {

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}};

	if (grid[x][y] == 0) {
		return
	}
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		var tx int = x + dir[i][0]
		var ty int = y + dir[i][1]
		if (isOK(tx, ty, len(grid), len(grid[0]))) {
			dfs(grid, tx, ty)
		}

	}
}

func CountIslands(grid [][]int)int {
	var island int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if (grid[i][j] == 1) {
				island++
				dfs(grid, i, j)
			}
		}
	}
	return island
}

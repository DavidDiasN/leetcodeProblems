package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	col := make(map[int]map[byte]bool)
	row := make(map[int]map[byte]bool)
	square := map[[2]int]map[byte]bool{}

	for i := 0; i < 9; i++ {
		col[i] = make(map[byte]bool)
		row[i] = make(map[byte]bool)
	}
	r, c := 0, 0

	for r < 3 {
		for c < 3 {
			square[[2]int{r, c}] = make(map[byte]bool)
			c++
		}
		c = 0
		r++
	}

	for r := range board {
		for c, val := range board[r] {
			if val == 45 || val == 46 {
				continue
			}
			_, ok1 := row[r][val]
			_, ok2 := col[c][val]
			_, ok3 := square[[2]int{r / 3, c / 3}][val]
			if ok1 || ok2 || ok3 {
				return false
			} else {
				row[r][val] = false
				col[c][val] = false
				square[[2]int{r / 3, c / 3}][val] = false
			}

		}
	}

	return true
}

package game

import "fmt"

type Board struct {
	Rows [][]string
	Size int
}

func (b *Board) Print() {
	fmt.Println()
	for i := 0; i < len(b.Rows); i++ {
		for j := 0; j < len(b.Rows[i]); j++ {
			item := b.Rows[i][j]
			if item == "" {
				item = "-"
			}
			fmt.Printf("%s  ", item)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) Mark(row int, col int, symbol string) error {
	if len(b.Rows) <= row || len(b.Rows[row]) <= col {
		return fmt.Errorf("invalid row or column index")
	}

	if b.Rows[row][col] != "" {
		return fmt.Errorf("cell already marked")
	}

	b.Rows[row][col] = symbol
	return nil
}

func (b *Board) IsFull() bool {
	for _, row := range b.Rows {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func (b *Board) IsWinningMove(row int, col int, symbol string) bool {
	if b.Rows[row][col] != symbol {
		return false
	}

	return isRowFilled(b, row, symbol) ||
		isColumnFilled(b, col, symbol) ||
		isDiagonalFilled(b, symbol)
}

func isRowFilled(b *Board, row int, symbol string) bool {
	for j := 0; j < b.Size; j++ {
		if b.Rows[row][j] != symbol {
			return false
		}
	}
	return true
}

func isColumnFilled(b *Board, col int, symbol string) bool {
	for i := 0; i < b.Size; i++ {
		if b.Rows[i][col] != symbol {
			return false
		}
	}
	return true
}

func isDiagonalFilled(b *Board, symbol string) bool {
	leftFilled := true
	rightFilled := true
	len := b.Size

	for i := 0; i < len; i++ {
		if b.Rows[i][i] != symbol {
			leftFilled = false
		}

		if b.Rows[i][len-1-i] != symbol {
			rightFilled = false
		}
	}

	return leftFilled || rightFilled
}

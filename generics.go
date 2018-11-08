package main

import (
	"math"

	vector "github.com/strosel/goutil/Vector"
)

func generateHorizontal(p *Piece, b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}
	for i := 0; i < 8; i++ {
		x := i
		y := p.Pos.Y
		if x != p.Pos.X {
			if !p.AttackingAllies(x, y, b) {
				if !p.MoveThroughPieces(x, y, b) {
					moves = append(moves, vector.Vector2I{X: x, Y: y})
				}
			}
		}
	}
	return moves
}

func generateVertical(p *Piece, b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}
	for i := 0; i < 8; i++ {
		x := p.Pos.X
		y := i
		if i != p.Pos.Y {
			if !p.AttackingAllies(x, y, b) {
				if !p.MoveThroughPieces(x, y, b) {
					moves = append(moves, vector.Vector2I{X: x, Y: y})
				}
			}
		}
	}
	return moves
}

func generateDiagonal(p *Piece, b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}
	for i := 0; i < 8; i++ {
		x := i
		y := p.Pos.Y - (p.Pos.X - i)
		if x != p.Pos.X {
			if p.WithinBounds(x, y) {
				if !p.AttackingAllies(x, y, b) {
					if !p.MoveThroughPieces(x, y, b) {
						moves = append(moves, vector.Vector2I{X: x, Y: y})
					}
				}
			}
		}
	}

	for i := 0; i < 8; i++ {
		x := p.Pos.X + (p.Pos.Y - i)
		y := i
		if x != p.Pos.X {
			if p.WithinBounds(x, y) {
				if !p.AttackingAllies(x, y, b) {
					if !p.MoveThroughPieces(x, y, b) {
						moves = append(moves, vector.Vector2I{X: x, Y: y})
					}
				}
			}
		}
	}
	return moves
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func generateBoards(p Piece, b *Board, moves []vector.Vector2I) []*Board {
	boards := []*Board{}

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(p.Pos, m)
	}

	return boards
}

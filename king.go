package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type King struct {
	*Piece
}

func NewKing(x, y int, isWhite bool) *King {
	return &King{
		Piece: &Piece{
			Pos:             vector.Vector2I{x, y},
			Taken:           false,
			White:           isWhite,
			MovingThisPiece: false,
			Letter:          'K',
			Value:           99,
		},
	}
}

func (k King) Clone() PieceI {
	king := NewKing(k.Pos.X, k.Pos.Y, k.White)
	king.Taken = k.Taken
	return king
}

func (k King) CanMove(x, y int, b *Board) bool {
	if !k.WithinBounds(x, y) {
		return false
	}
	if k.AttackingAllies(x, y, b) {
		return false
	}
	if abs(x-k.Pos.X) <= 1 && abs(y-k.Pos.Y) <= 1 {
		return true
	}
	return false
}

func (k King) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			x := k.Pos.X + i
			y := k.Pos.Y + j
			if k.WithinBounds(x, y) {
				if i != 0 || j != 0 {
					if !k.AttackingAllies(x, y, b) {
						moves = append(moves, vector.Vector2I{X: x, Y: y})
					}
				}
			}
		}

	}
	return moves
}

func (k King) GenerateNewBoards(b *Board) []*Board {
	boards := []*Board{}
	moves := k.GenerateMoves(b)

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(k.Pos, m)
	}
	return boards
}

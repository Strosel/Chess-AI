package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type Bishop struct {
	*Piece
}

func NewBishop(x, y int, isWhite bool) *Bishop {
	return &Bishop{
		Piece: &Piece{
			Pos:    vector.Vector2I{x, y},
			Taken:  false,
			White:  isWhite,
			Letter: 'B',
			Value:  3,
		},
	}
}

func (bs Bishop) Clone() PieceI {
	bishop := NewBishop(bs.Pos.X, bs.Pos.Y, bs.White)
	bishop.Taken = bs.Taken
	return bishop
}

func (bs Bishop) CanMove(x, y int, b *Board) bool {
	if !bs.WithinBounds(x, y) {
		return false
	}
	if bs.AttackingAllies(x, y, b) {
		return false
	}

	//diagonal
	if abs(x-bs.Pos.X) == abs(y-bs.Pos.Y) {
		if bs.MoveThroughPieces(x, y, b) {
			return false
		}

		return true
	}
	return false
}

func (bs Bishop) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}

	diagonal := generateDiagonal(bs.Piece, b)
	moves = append(moves, diagonal...)

	return moves
}

func (bs Bishop) GenerateNewBoards(b *Board) []*Board {
	moves := bs.GenerateMoves(b)
	boards := generateBoards(*bs.Piece, b, moves)

	return boards
}

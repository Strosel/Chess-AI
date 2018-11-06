package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type Rook struct {
	*Piece
}

func NewRook(x, y int, isWhite bool) *Rook {
	return &Rook{
		Piece: &Piece{
			Pos:    vector.Vector2I{x, y},
			Taken:  false,
			White:  isWhite,
			Letter: 'R',
			Value:  5,
			Moves:  0,
		},
	}
}

func (r Rook) Clone() PieceI {
	rook := NewRook(r.Pos.X, r.Pos.Y, r.White)
	rook.Taken = r.Taken
	rook.Moves = r.Moves
	return rook
}

func (r Rook) CanMove(x, y int, b *Board) bool {
	if !r.WithinBounds(x, y) {
		return false
	}
	if r.AttackingAllies(x, y, b) {
		return false
	}

	if x == r.Pos.X || y == r.Pos.Y {
		if r.MoveThroughPieces(x, y, b) {
			return false
		}

		return true
	}

	return false
}

func (r Rook) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}

	horizontal := generateHorizontal(r.Piece, b)
	vertical := generateVertical(r.Piece, b)

	moves = append(moves, horizontal...)
	moves = append(moves, vertical...)

	return moves
}

func (r Rook) GenerateNewBoards(b *Board) []*Board {
	boards := []*Board{}
	moves := r.GenerateMoves(b)

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(r.Pos, m)
	}
	return boards
}

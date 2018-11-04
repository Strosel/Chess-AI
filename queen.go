package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type Queen struct {
	*Piece
}

func NewQueen(x, y int, isWhite bool) *Queen {
	return &Queen{
		Piece: &Piece{
			Pos:             vector.Vector2I{x, y},
			Taken:           false,
			White:           isWhite,
			MovingThisPiece: false,
			Letter:          'Q',
			Value:           9,
		},
	}
}

func (q Queen) Clone() PieceI {
	queen := NewQueen(q.Pos.X, q.Pos.Y, q.White)
	queen.Taken = q.Taken
	return queen
}

func (q Queen) CanMove(x, y int, b *Board) bool {
	if !q.WithinBounds(x, y) {
		return false
	}
	if q.AttackingAllies(x, y, b) {
		return false
	}

	if q.MoveThroughPieces(x, y, b) {
		return false
	}

	if x == q.Pos.X || y == q.Pos.Y {
		return true
	}

	if abs(x-q.Pos.X) == abs(y-q.Pos.Y) {
		return true
	}

	return false
}

func (q Queen) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}

	horizontal := generateHorizontal(q.Piece, b)
	vertical := generateVertical(q.Piece, b)
	diagonal := generateDiagonal(q.Piece, b)

	moves = append(moves, horizontal...)
	moves = append(moves, vertical...)
	moves = append(moves, diagonal...)

	return moves
}

func (q Queen) GenerateNewBoards(b *Board) []*Board {
	boards := []*Board{}
	moves := q.GenerateMoves(b)

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(q.Pos, m)
	}
	return boards
}

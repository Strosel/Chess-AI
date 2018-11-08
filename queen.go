package main

import (
	vector "github.com/strosel/goutil/Vector"
)

//Queen Defines the Queen Piece
type Queen struct {
	*Piece
}

//NewQueen Create a new Queen Piece
func NewQueen(x, y int, isWhite bool) *Queen {
	return &Queen{
		Piece: &Piece{
			Pos:    vector.Vector2I{X: x, Y: y},
			Taken:  false,
			White:  isWhite,
			Letter: 'Q',
			Value:  9,
		},
	}
}

//Clone Clone a new Queen Piece
func (q Queen) Clone() PieceI {
	queen := NewQueen(q.Pos.X, q.Pos.Y, q.White)
	queen.Taken = q.Taken
	return queen
}

//CanMove Check if the Queen can move to a point on the Board
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

//GenerateMoves Generaet a set of moves
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

//GenerateNewBoards Generate a new Board for each move
func (q Queen) GenerateNewBoards(b *Board) []*Board {
	moves := q.GenerateMoves(b)
	boards := generateBoards(*q.Piece, b, moves)

	return boards
}

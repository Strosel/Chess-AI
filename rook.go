package main

import (
	vector "github.com/strosel/goutil/Vector"
)

//Rook Defines the Rook Piece
type Rook struct {
	*Piece
}

//NewRook Create a new Rook Piece
func NewRook(x, y int, isWhite bool) *Rook {
	return &Rook{
		Piece: &Piece{
			Pos:    vector.Vector2I{X: x, Y: y},
			Taken:  false,
			White:  isWhite,
			Letter: 'R',
			Value:  5,
			Moves:  0,
		},
	}
}

//Clone Clone a new Rook Piece
func (r Rook) Clone() PieceI {
	rook := NewRook(r.Pos.X, r.Pos.Y, r.White)
	rook.Taken = r.Taken
	rook.Moves = r.Moves
	return rook
}

//CanMove Check if the Rook can move to a point on the Board
func (r Rook) CanMove(x, y int, b *Board) bool {
	if !r.Piece.CanMove(x, y, b) {
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

//GenerateMoves Generaet a set of moves
func (r Rook) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}

	horizontal := generateHorizontal(r.Piece, b)
	vertical := generateVertical(r.Piece, b)

	moves = append(moves, horizontal...)
	moves = append(moves, vertical...)

	return moves
}

//GenerateNewBoards Generate a new Board for each move
func (r Rook) GenerateNewBoards(b *Board) []*Board {
	moves := r.GenerateMoves(b)
	boards := generateBoards(*r.Piece, b, moves)

	return boards
}

package main

import (
	vector "github.com/strosel/goutil/Vector"
)

//Knight Defines the Knight Piece
type Knight struct {
	*Piece
}

//NewKnight Create a new Knight Piece
func NewKnight(x, y int, isWhite bool) *Knight {
	return &Knight{
		Piece: &Piece{
			Pos:    vector.Vector2I{X: x, Y: y},
			Taken:  false,
			White:  isWhite,
			Letter: 'H', //Horse
			Value:  3,
		},
	}
}

//Clone Clone a new Knight Piece
func (k Knight) Clone() PieceI {
	rook := NewKnight(k.Pos.X, k.Pos.Y, k.White)
	rook.Taken = k.Taken
	return rook
}

//CanMove Check if the Knight can move to a point on the Board
func (k Knight) CanMove(x, y int, b *Board) bool {
	if !k.WithinBounds(x, y) {
		return false
	}
	if k.AttackingAllies(x, y, b) {
		return false
	}

	if (abs(x-k.Pos.X) == 2 && abs(y-k.Pos.Y) == 1) || (abs(x-k.Pos.X) == 1 && abs(y-k.Pos.Y) == 2) {
		return true
	}

	return false
}

//GenerateMoves Generaet a set of moves
func (k Knight) GenerateMoves(b *Board) []vector.Vector2I {
	moves := []vector.Vector2I{}

	for i := -2; i < 3; i += 4 {
		for j := -1; j < 2; j += 2 {

			var x = i + k.Pos.X
			var y = j + k.Pos.Y
			if !k.AttackingAllies(x, y, b) {
				if k.WithinBounds(x, y) {
					moves = append(moves, vector.Vector2I{X: x, Y: y})
				}
			}
		}
	}
	for i := -1; i < 2; i += 2 {
		for j := -2; j < 3; j += 4 {

			var x = i + k.Pos.X
			var y = j + k.Pos.Y

			if k.WithinBounds(x, y) {
				if !k.AttackingAllies(x, y, b) {
					moves = append(moves, vector.Vector2I{X: x, Y: y})
				}
			}
		}
	}

	return moves
}

//GenerateNewBoards Generate a new Board for each move
func (k Knight) GenerateNewBoards(b *Board) []*Board {
	moves := k.GenerateMoves(b)
	boards := generateBoards(*k.Piece, b, moves)

	return boards
}

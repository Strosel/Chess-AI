package main

import (
	vector "github.com/strosel/goutil/Vector"
)

//King Defines the King Piece
type King struct {
	*Piece
}

//NewKing Create a new King Piece
func NewKing(x, y int, isWhite bool) *King {
	return &King{
		Piece: &Piece{
			Pos:    vector.Vector2I{X: x, Y: y},
			Taken:  false,
			White:  isWhite,
			Letter: 'K',
			Value:  99,
			Moves:  0,
		},
	}
}

//Clone Clone a new King Piece
func (k King) Clone() PieceI {
	king := NewKing(k.Pos.X, k.Pos.Y, k.White)
	king.Taken = k.Taken
	king.Moves = k.Moves
	return king
}

//CanMove Check if the King can move to a point on the Board
func (k King) CanMove(x, y int, b *Board) bool {
	if !k.WithinBounds(x, y) {
		return false
	}

	attacking := b.GetPieceAt(x, y)
	if attacking != nil && attacking.IsWhite() == k.White { // Moving on own piece
		if k.Moves == 0 && !b.InCheck(k.White) && attacking.GetLetter() == 'R' && attacking.GetMoves() == 0 { // Castling
			if (x == 7 && b.GetPieceAt(5, y) == nil && b.GetPieceAt(6, y) == nil && b.IsSafe(5, y, k.White) && b.IsSafe(6, y, k.White)) ||
				(x == 0 && b.GetPieceAt(3, y) == nil && b.GetPieceAt(2, y) == nil && b.IsSafe(3, y, k.White) && b.IsSafe(2, y, k.White)) {
				return true
			}
		}
		// Attacking Own Piece
		return false
	}

	if abs(x-k.Pos.X) <= 1 && abs(y-k.Pos.Y) <= 1 {
		return true
	}

	return false
}

//GenerateMoves Generaet a set of moves
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

	// Add castling if possible
	rook := k.Position()
	rook.X = 0
	if k.CanMove(rook.X, rook.Y, b) {
		moves = append(moves, rook)
	}
	rook.X = 7
	if k.CanMove(rook.X, rook.Y, b) {
		moves = append(moves, rook)
	}

	return moves
}

//GenerateNewBoards Generate a new Board for each move
func (k King) GenerateNewBoards(b *Board) []*Board {
	moves := k.GenerateMoves(b)
	boards := k.generateBoards(b, moves)

	return boards
}

//Move Move the piece on the board
func (k *King) Move(x, y int, b *Board) {
	attacking := b.GetPieceAt(x, y)
	if attacking != nil && attacking.IsWhite() != k.White {
		attacking.SetTaken(true)
	} else if attacking != nil && attacking.IsWhite() == k.White && attacking.GetLetter() == 'R' {
		if x == 7 {
			k.Pos = vector.Vector2I{X: 6, Y: y}
			attacking.Move(5, y, b)
		} else if x == 0 {
			k.Pos = vector.Vector2I{X: 2, Y: y}
			attacking.Move(3, y, b)
		}
		attacking.IncrementMoves()
		k.Moves++
		return
	}
	k.Moves++
	k.Pos = vector.Vector2I{X: x, Y: y}
}

package main

import (
	vector "github.com/strosel/goutil/Vector"
)

//Pawn Defines the Pawn Piece
type Pawn struct {
	*Piece
}

//NewPawn Create a new Pawn Piece
func NewPawn(x, y int, isWhite bool) *Pawn {
	return &Pawn{
		Piece: &Piece{
			Pos:           vector.Vector2I{X: x, Y: y},
			Taken:         false,
			White:         isWhite,
			Letter:        'p',
			Value:         1,
			Moves:         0,
			LastTurnMoved: 0,
		},
	}
}

//Clone Clone a new Pawn Piece
func (p Pawn) Clone() PieceI {
	pawn := NewPawn(p.Pos.X, p.Pos.Y, p.White)
	pawn.Taken = p.Taken
	pawn.Moves = p.Moves
	pawn.LastTurnMoved = p.LastTurnMoved
	return pawn
}

//CanMove Check if the Pawn can move to a point on the Board
func (p *Pawn) CanMove(x, y int, b *Board) bool {
	if !p.Piece.CanMove(x, y, b) {
		return false
	}

	if p.EnPassant(x, y, b) != nil {
		return true
	}

	attacking := b.IsPieceAt(x, y)
	if attacking {
		//if attacking a player
		if abs(x-p.Pos.X) == abs(y-p.Pos.Y) && ((p.White && (y-p.Pos.Y) == -1) || (!p.White && (y-p.Pos.Y) == 1)) {
			p.Moves++
			return true
		}
		return false
	}
	if x != p.Pos.X {
		return false
	}
	if (p.White && y-p.Pos.Y == -1) || (!p.White && y-p.Pos.Y == 1) {
		p.Moves++
		return true
	}
	if p.Moves < 1 && ((p.White && y-p.Pos.Y == -2) || (!p.White && y-p.Pos.Y == 2)) {
		if p.MoveThroughPieces(x, y, b) {
			return false
		}

		p.Moves++
		return true
	}

	return false
}

//GenerateMoves Generaet a set of moves
func (p Pawn) GenerateMoves(b *Board) []vector.Vector2I {
	x := 0
	y := 0
	moves := []vector.Vector2I{}

	for i := -1; i < 2; i += 2 {
		x = p.Pos.X + i
		if p.White {
			y = p.Pos.Y - 1
		} else {
			y = p.Pos.Y + 1
		}
		attacking := b.IsPieceAt(x, y)
		if attacking {
			if !p.AttackingAllies(x, y, b) {
				moves = append(moves, vector.Vector2I{X: x, Y: y})
			}
		}

		if p.EnPassant(x, y, b) != nil {
			moves = append(moves, vector.Vector2I{X: x, Y: y})
		}
	}

	x = p.Pos.X
	if p.White {
		y = p.Pos.Y - 1
	} else {
		y = p.Pos.Y + 1
	}
	if !b.IsPieceAt(x, y) && p.WithinBounds(x, y) {
		moves = append(moves, vector.Vector2I{X: x, Y: y})
	}

	if p.Moves < 1 {

		if p.White {
			y = p.Pos.Y - 2
		} else {
			y = p.Pos.Y + 2
		}
		if !b.IsPieceAt(x, y) && p.WithinBounds(x, y) {
			if !p.MoveThroughPieces(x, y, b) {
				moves = append(moves, vector.Vector2I{X: x, Y: y})
			}
		}
	}

	return moves
}

//GenerateNewBoards Generate a new Board for each move
func (p Pawn) GenerateNewBoards(b *Board) []*Board {
	moves := p.GenerateMoves(b)
	boards := p.generateBoards(b, moves)

	return boards
}

//Move Move the piece on the board
func (p *Pawn) Move(x, y int, b *Board) {
	attacking := p.EnPassant(x, y, b)
	if attacking != nil {
		attacking.SetTaken(true)
	}
	p.Piece.Move(x, y, b)
	p.LastTurnMoved = turn
}

//EnPassant Check if a move to x;y is an En Passant
func (p Pawn) EnPassant(x, y int, b *Board) PieceI {
	attacking := b.GetPieceAt(x, y)
	if abs(p.Pos.X-x) == 1 && abs(p.Pos.Y-y) == 1 && attacking == nil {
		if p.White {
			attacking = b.GetPieceAt(x, p.Pos.Y)
			if attacking != nil && attacking.GetLetter() == 'p' && attacking.GetMoves() == 1 && attacking.IsWhite() != p.White && p.Pos.Y == 3 && abs(turn-attacking.GetLastTurn()) <= 1 {
				return attacking
			}
		} else {
			attacking = b.GetPieceAt(x, p.Pos.Y)
			if attacking != nil && attacking.GetLetter() == 'p' && attacking.GetMoves() == 1 && attacking.IsWhite() != p.White && p.Pos.Y == 4 && abs(turn-attacking.GetLastTurn()) <= 1 {
				return attacking
			}
		}
	}

	return nil
}

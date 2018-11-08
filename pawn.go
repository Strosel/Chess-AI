package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type Pawn struct {
	*Piece
	LastTurnMoved int
}

func NewPawn(x, y int, isWhite bool) *Pawn {
	return &Pawn{
		Piece: &Piece{
			Pos:    vector.Vector2I{x, y},
			Taken:  false,
			White:  isWhite,
			Letter: 'p',
			Value:  1,
			Moves:  0,
		},
		LastTurnMoved: 0,
	}
}

func (p Pawn) Clone() PieceI {
	pawn := NewPawn(p.Pos.X, p.Pos.Y, p.White)
	pawn.Taken = p.Taken
	pawn.Moves = p.Moves
	pawn.LastTurnMoved = p.LastTurnMoved
	return pawn
}

func (p *Pawn) CanMove(x, y int, b *Board) bool {
	if !p.WithinBounds(x, y) {
		return false
	}
	if p.AttackingAllies(x, y, b) {
		return false
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

func (p Pawn) GenerateNewBoards(b *Board) []*Board {
	moves := p.GenerateMoves(b)
	boards := generateBoards(*p.Piece, b, moves)

	return boards
}

func (p *Pawn) Move(x, y int, b *Board) {
	attacking := b.GetPieceAt(x, y)
	if attacking != nil {
		attacking.SetTaken(true)
	}
	p.Pos = vector.Vector2I{x, y}
	p.LastTurnMoved = turn
}

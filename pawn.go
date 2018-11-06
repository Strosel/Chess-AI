package main

import (
	vector "github.com/strosel/goutil/Vector"
)

type Pawn struct {
	*Piece
	FirstTurn bool
}

func NewPawn(x, y int, isWhite bool) *Pawn {
	return &Pawn{
		Piece: &Piece{
			Pos:    vector.Vector2I{x, y},
			Taken:  false,
			White:  isWhite,
			Letter: 'p',
			Value:  1,
		},
		FirstTurn: true,
	}
}

func (p Pawn) Clone() PieceI {
	pawn := NewPawn(p.Pos.X, p.Pos.Y, p.White)
	pawn.Taken = p.Taken
	pawn.FirstTurn = p.FirstTurn
	return pawn
}

func (p Pawn) CanMove(x, y int, b *Board) bool {
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
			p.FirstTurn = false
			return true
		}
		return false
	}
	if x != p.Pos.X {
		return false
	}
	if (p.White && y-p.Pos.Y == -1) || (!p.White && y-p.Pos.Y == 1) {
		p.FirstTurn = false
		return true
	}
	if p.FirstTurn && ((p.White && y-p.Pos.Y == -2) || (!p.White && y-p.Pos.Y == 2)) {
		if p.MoveThroughPieces(x, y, b) {
			return false
		}

		p.FirstTurn = false
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

	if p.FirstTurn {

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
	boards := []*Board{}
	moves := p.GenerateMoves(b)

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(p.Pos, m)
	}
	return boards
}

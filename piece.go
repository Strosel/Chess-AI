package main

import (
	"fmt"

	"github.com/strosel/goutil/Vector"
)

//PieceI Piece interface
type PieceI interface {
	GenerateNewBoards(*Board) []*Board
	WithinBounds(int, int) bool
	Move(int, int, *Board)
	AttackingAllies(int, int, *Board) bool
	CanMove(int, int, *Board) bool
	MoveThroughPieces(int, int, *Board) bool
	GenerateMoves(*Board) []vector.Vector2I
	Position() vector.Vector2I
	IsTaken() bool
	SetTaken(bool)
	IsWhite() bool
	GetValue() int
	GetLetter() byte
	Clone() PieceI
	GetMoves() int
	IncrementMoves()
	GetLastTurn() int
}

//Piece Defines a generic Piece
type Piece struct {
	Pos                         vector.Vector2I //matrixpos
	Taken, White                bool
	Letter                      byte
	Value, Moves, LastTurnMoved int
}

//GenerateNewBoards Generates a new board for each possible Move of the Piece
func (p Piece) GenerateNewBoards(b *Board) []*Board {
	moves := p.GenerateMoves(b)
	boards := generateBoards(p, b, moves) // Replace ALL meth occurences to this version

	return boards
}

//GenerateMoves Generate moves for the Piece
func (p Piece) GenerateMoves(b *Board) []vector.Vector2I {
	return []vector.Vector2I{}
}

//WithinBounds Is the Piece still on the board
func (p Piece) WithinBounds(x, y int) bool {
	if x >= 0 && y >= 0 && x < 8 && y < 8 {
		return true
	}
	return false
}

//Move Move the piece to x, y on b
func (p *Piece) Move(x, y int, b *Board) {
	attacking := b.GetPieceAt(x, y)
	if attacking != nil {
		attacking.SetTaken(true)
	}
	p.Pos = vector.Vector2I{X: x, Y: y}
	p.Moves++
}

//AttackingAllies Is the Piece trying to attack an ally
func (p Piece) AttackingAllies(x, y int, b *Board) bool {
	attacking := b.GetPieceAt(x, y)
	if attacking != nil {
		if attacking.IsWhite() == p.White {
			return true
		}
	}
	return false
}

//CanMove Can the piece move to c, y on b
func (p Piece) CanMove(x, y int, b *Board) bool {
	fmt.Println("Can move")
	if !p.WithinBounds(x, y) {
		return false
	}
	return true
}

//MoveThroughPieces Will a move to x, y on b move through another Piece
func (p Piece) MoveThroughPieces(x, y int, b *Board) bool {
	stepDirectionX := x - p.Pos.X
	if stepDirectionX > 0 {
		stepDirectionX = 1
	} else if stepDirectionX < 0 {
		stepDirectionX = -1
	}
	stepDirectionY := y - p.Pos.Y
	if stepDirectionY > 0 {
		stepDirectionY = 1
	} else if stepDirectionY < 0 {
		stepDirectionY = -1
	}
	tempPos := vector.Vector2I{X: p.Pos.X, Y: p.Pos.Y}
	tempPos.X += stepDirectionX
	tempPos.Y += stepDirectionY
	for p.WithinBounds(tempPos.X, tempPos.Y) && (tempPos.X != x || tempPos.Y != y) {
		tmp := b.GetPieceAt(tempPos.X, tempPos.Y)
		if tmp != nil {
			return true
		}
		tempPos.X += stepDirectionX
		tempPos.Y += stepDirectionY
	}

	return false
}

//Position Get Piece.Pos
func (p Piece) Position() vector.Vector2I {
	return p.Pos
}

//IsTaken Get Piece.Taken
func (p Piece) IsTaken() bool {
	return p.Taken
}

//SetTaken Set Piece.Taken
func (p *Piece) SetTaken(set bool) {
	p.Taken = set
}

//IsWhite Get Piece.White
func (p Piece) IsWhite() bool {
	return p.White
}

//GetValue Get Piece.Value
func (p Piece) GetValue() int {
	return p.Value
}

//GetLetter Get Piece.Letter
func (p Piece) GetLetter() byte {
	return p.Letter
}

//GetMoves Get Piece.Moves
func (p Piece) GetMoves() int {
	return p.Moves
}

//IncrementMoves Increment Piece.Moves by 1
func (p *Piece) IncrementMoves() {
	p.Moves++
}

//GetLastTurn Get Piece.LastTurnMoved
func (p Piece) GetLastTurn() int {
	return p.LastTurnMoved
}

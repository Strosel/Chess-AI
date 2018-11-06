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
}

//Piece Defines a generic Piece
type Piece struct {
	Pos          vector.Vector2I //matrixpos
	Taken, White bool
	Letter       byte
	Value        int
}

//GenerateNewBoards Generates a new board for each possible Move of the Piece
func (p Piece) GenerateNewBoards(b *Board) []*Board {
	boards := []*Board{}
	moves := p.GenerateMoves(b)

	for i, m := range moves {
		boards = append(boards, b.Clone())
		boards[i].Move(p.Pos, m)
	}
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
	p.Pos = vector.Vector2I{x, y}
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
	tempPos := vector.Vector2I{p.Pos.X, p.Pos.Y}
	tempPos.X += stepDirectionX
	tempPos.Y += stepDirectionY
	for tempPos.X != x || tempPos.Y != y {
		tmp := b.GetPieceAt(tempPos.X, tempPos.Y)
		if tmp != nil {
			return true
		}
		tempPos.X += stepDirectionX
		tempPos.Y += stepDirectionY
	}

	return false
}

func (p Piece) Position() vector.Vector2I {
	return p.Pos
}

func (p Piece) IsTaken() bool {
	return p.Taken
}

func (p *Piece) SetTaken(set bool) {
	p.Taken = set
}

func (p Piece) IsWhite() bool {
	return p.White
}

func (p Piece) GetValue() int {
	return p.Value
}

func (p Piece) GetLetter() byte {
	return p.Letter
}

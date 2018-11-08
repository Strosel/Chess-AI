package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	vector "github.com/strosel/goutil/Vector"
)

//Board Defines a Board
type Board struct {
	BlackPieces, WhitePieces []PieceI
	Score                    int
}

//NewBoard Create a new Board
func NewBoard() *Board {
	b := &Board{}
	b.SetupPieces()
	return b
}

//SetupPieces Sets the board to start position
func (b *Board) SetupPieces() {
	//White Pieces
	b.WhitePieces = append(b.WhitePieces, NewKing(4, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewQueen(3, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewBishop(2, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewBishop(5, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewKnight(1, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewRook(0, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewKnight(6, 7, true))
	b.WhitePieces = append(b.WhitePieces, NewRook(7, 7, true))

	for i := 0; i < 8; i++ {
		b.WhitePieces = append(b.WhitePieces, NewPawn(i, 6, true))
	}

	//Black Pieces
	b.BlackPieces = append(b.BlackPieces, NewKing(4, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewQueen(3, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewBishop(2, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewBishop(5, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewKnight(1, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewRook(0, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewKnight(6, 0, false))
	b.BlackPieces = append(b.BlackPieces, NewRook(7, 0, false))

	for i := 0; i < 8; i++ {
		b.BlackPieces = append(b.BlackPieces, NewPawn(i, 1, false))
	}
}

//IsPieceAt Checs if there is a Piece in the given location
func (b Board) IsPieceAt(x, y int) bool {
	for _, wp := range b.WhitePieces {
		if !wp.IsTaken() && wp.Position().X == x && wp.Position().Y == y {
			return true
		}
	}

	for _, bp := range b.BlackPieces {
		if !bp.IsTaken() && bp.Position().X == x && bp.Position().Y == y {
			return true
		}
	}

	return false
}

//GetPieceAt Gets the Piece at the given location
func (b Board) GetPieceAt(x, y int) PieceI {
	for _, wp := range b.WhitePieces {
		if !wp.IsTaken() && wp.Position().X == x && wp.Position().Y == y {
			return wp
		}
	}

	for _, bp := range b.BlackPieces {
		if !bp.IsTaken() && bp.Position().X == x && bp.Position().Y == y {
			return bp
		}
	}

	return nil
}

//GenerateNewBoardsWhitesTurn Generates all possible Boards after 1 white move
func (b Board) GenerateNewBoardsWhitesTurn() []*Board {
	boards := []*Board{}
	for _, wp := range b.WhitePieces {
		if !wp.IsTaken() {
			tmp := wp.GenerateNewBoards(&b)
			boards = append(boards, tmp...)
		}
	}
	return boards
}

//GenerateNewBoardsBlacksTurn Generates all possible Boards after 1 black move
func (b Board) GenerateNewBoardsBlacksTurn() []*Board {
	boards := []*Board{}
	for _, bp := range b.BlackPieces {
		if !bp.IsTaken() {
			tmp := bp.GenerateNewBoards(&b)
			boards = append(boards, tmp...)
		}
	}
	return boards
}

//SetScore Sets the score of the board at current state
func (b *Board) SetScore() {
	b.Score = 0

	for _, wp := range b.WhitePieces {
		if !wp.IsTaken() {
			b.Score -= wp.GetValue()
		}
	}

	for _, bp := range b.BlackPieces {
		if !bp.IsTaken() {
			b.Score += bp.GetValue()
		}
	}
}

//Move Move the Piece at from to to
func (b *Board) Move(from, to vector.Vector2I) {
	pieceToMove := b.GetPieceAt(from.X, from.Y)
	if pieceToMove == nil {
		return
	}
	pieceToMove.Move(to.X, to.Y, b)
}

//Clone Clones the board
func (b Board) Clone() *Board {
	clone := &Board{}
	for _, wp := range b.WhitePieces {
		clone.WhitePieces = append(clone.WhitePieces, wp.Clone())
	}

	for _, bp := range b.BlackPieces {
		clone.BlackPieces = append(clone.BlackPieces, bp.Clone())
	}
	return clone
}

//IsDone Is the game done
func (b Board) IsDone() bool {
	return b.WhitePieces[0].IsTaken() || b.BlackPieces[0].IsTaken() //[0] Is the king
}

//IsDead Is any player dead
func (b Board) IsDead() bool {
	if whiteAI && whitesMove {
		return b.WhitePieces[0].IsTaken()
	}
	if blackAI && !whitesMove {
		return b.BlackPieces[0].IsTaken()
	}

	return false
}

//HasWon Has any player won
func (b Board) HasWon() bool {
	if whiteAI && whitesMove {
		return b.BlackPieces[0].IsTaken()
	}
	if blackAI && !whitesMove {
		return b.WhitePieces[0].IsTaken()
	}

	return false
}

func (b Board) String() string {
	out := ""

	// player on top

	if whiteAI && !blackAI {
		fmt.Printf("Player is %v\nAI is %v\n", aurora.BgBlack(aurora.Gray("Black")), aurora.BgGray(aurora.Black("White")))
		out += "  A B C D E F G H\n"
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				piece := b.GetPieceAt(x, y)
				if x == 0 {
					out += fmt.Sprintf("%v ", y+1)
				}
				if piece == nil {
					out += "  "
				} else if piece.IsWhite() {
					out += fmt.Sprintf("%v ", aurora.BgGray(aurora.Black(string(piece.GetLetter()))))
				} else if string(piece.GetLetter()) != " " {
					out += fmt.Sprintf("%v ", aurora.BgBlack(aurora.Gray(string(piece.GetLetter()))))
				} else {
					out += "  "
				}
			}
			out += "\n"
		}
		out += "  A B C D E F G H\n"
	} else if !whiteAI && blackAI {
		fmt.Printf("Player is %v\nAI is %v\n", aurora.BgGray(aurora.Black("White")), aurora.BgBlack(aurora.Gray("Black")))
		out += "  H G F E D C B A\n"
		for y := 7; y > -1; y-- {
			for x := 7; x > -1; x-- {
				piece := b.GetPieceAt(x, y)
				if x == 7 {
					out += fmt.Sprintf("%v ", y+1)
				}
				if piece == nil {
					out += "  "
				} else if piece.IsWhite() {
					out += fmt.Sprintf("%v ", aurora.BgGray(aurora.Black(string(piece.GetLetter()))))
				} else if string(piece.GetLetter()) != " " {
					out += fmt.Sprintf("%v ", aurora.BgBlack(aurora.Gray(string(piece.GetLetter()))))
				} else {
					out += "  "
				}
			}
			out += "\n"
		}
		out += "  H G F E D C B A\n"
	}
	return out
}

//Diff Return the move made between b and b2 if any
func (b Board) Diff(b2 Board) string {
	if blackAI {
		if !b.BlackPieces[0].Position().Equals(b2.BlackPieces[0].Position()) && (!b.BlackPieces[5].Position().Equals(b2.BlackPieces[5].Position()) || !b.BlackPieces[7].Position().Equals(b2.BlackPieces[7].Position())) {
			fX := string(byte(b.BlackPieces[0].Position().X + 97))
			fY := b.BlackPieces[0].Position().Y + 1
			tX := string(byte(b2.BlackPieces[0].Position().X + 97))
			tY := b2.BlackPieces[0].Position().Y + 1
			return fmt.Sprintf("Castling, king %v%v to %v%v", fX, fY, tX, tY)
		}
		for i := 0; i < len(b.BlackPieces); i++ {
			if !b.BlackPieces[i].Position().Equals(b2.BlackPieces[i].Position()) {
				fX := string(byte(b.BlackPieces[i].Position().X + 97))
				fY := b.BlackPieces[i].Position().Y + 1
				tX := string(byte(b2.BlackPieces[i].Position().X + 97))
				tY := b2.BlackPieces[i].Position().Y + 1
				return fmt.Sprintf("%v%v %v%v", fX, fY, tX, tY)
			}
		}
	} else if whiteAI {
		if !b.WhitePieces[0].Position().Equals(b2.WhitePieces[0].Position()) && (!b.WhitePieces[5].Position().Equals(b2.WhitePieces[5].Position()) || !b.WhitePieces[7].Position().Equals(b2.WhitePieces[7].Position())) {
			fX := string(byte(b.WhitePieces[0].Position().X + 97))
			fY := b.WhitePieces[0].Position().Y + 1
			tX := string(byte(b2.WhitePieces[0].Position().X + 97))
			tY := b2.WhitePieces[0].Position().Y + 1
			return fmt.Sprintf("Castling, king %v%v to %v%v", fX, fY, tX, tY)
		}
		for i := 0; i < len(b.WhitePieces); i++ {
			if !b.WhitePieces[i].Position().Equals(b2.WhitePieces[i].Position()) {
				fX := string(byte(b.WhitePieces[i].Position().X + 97))
				fY := b.WhitePieces[i].Position().Y + 1
				tX := string(byte(b2.WhitePieces[i].Position().X + 97))
				tY := b2.WhitePieces[i].Position().Y + 1
				return fmt.Sprintf("%v%v %v%v", fX, fY, tX, tY)
			}
		}
	}
	return ""
}

//Winner Return who won
func (b Board) Winner() string {
	if b.BlackPieces[0].IsTaken() {
		return fmt.Sprintf("%v", aurora.BgGray(aurora.Black("White")))
	} else if b.WhitePieces[0].IsTaken() {
		return fmt.Sprintf("%v", aurora.BgBlack(aurora.Gray("Black")))
	}
	return ""
}

//IsSafe Check if space x;y is safe for the King
func (b *Board) IsSafe(x, y int, white bool) bool {
	if white {
		for _, bp := range b.BlackPieces {
			if bp.CanMove(x, y, b) {
				return false
			}
		}
	} else {
		for _, wp := range b.WhitePieces {
			if wp.CanMove(x, y, b) {
				return false
			}
		}
	}
	return true
}

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/strosel/goinput"
)

var (
	curBoard, lastBoard *Board
	whiteAI             bool
	blackAI             bool
	maxDepth            int

	turn       = 1
	whitesMove = true
	movePiece  = regexp.MustCompile("([a-h])(\\d) ([a-h])(\\d)")
)

func runAI() {
	depth := 0
	if !curBoard.IsDead() && !curBoard.HasWon() {
		if blackAI {
			if !whitesMove {
				_, tmp := maxFunAB(curBoard, -400, 400, depth)
				fmt.Println("Bot>", curBoard.Diff(*tmp))
				curBoard = tmp
				whitesMove = true
			}
		}
		if whiteAI {
			if whitesMove {
				_, tmp := minFunAB(curBoard, -400, 400, depth)
				fmt.Println("Bot>", curBoard.Diff(*tmp))
				curBoard = tmp
				whitesMove = false
			}
		}
	}
}

func runPlayer() {
	cmd := strings.ToLower(goinput.Strin("Player> "))
	if cmd == "help" {
		fmt.Println("")
		fmt.Println("Comands include:")
		fmt.Println("\ta2 b3\tmove piece in a2 to b3")
		fmt.Println("\t\tTo castle move the King on top of the Rook to castle with")
		fmt.Println("\tprint\tprint the current board")
		fmt.Println("\tundo\tundo the last player move")
		fmt.Println("\texit\texit the game")
		fmt.Println("")
	} else if cmd == "exit" {
		fmt.Println("")
		os.Exit(0)
	} else if cmd == "print" {
		fmt.Println("")
		fmt.Println(curBoard)
	} else if cmd == "undo" {
		if lastBoard == nil {
			fmt.Println("Can't Undo")
			return
		}
		curBoard = lastBoard.Clone()
		turn--
		lastBoard = nil
		fmt.Println("Undo")
	} else if movePiece.MatchString(cmd) {
		coords := movePiece.FindAllStringSubmatch(cmd, -1)[0]
		fromX := int(byte(coords[1][0])) - 97
		fromY, _ := strconv.Atoi(coords[2])
		fromY--
		toX := int(byte(coords[3][0])) - 97
		toY, _ := strconv.Atoi(coords[4])
		toY--
		if !curBoard.IsDone() {
			movingPiece := curBoard.GetPieceAt(fromX, fromY)
			if movingPiece == nil || movingPiece.IsWhite() != whitesMove {
				fmt.Println("Not Your Piece")
				return
			}

			if movingPiece.CanMove(toX, toY, curBoard) {
				lastBoard = curBoard.Clone()
				movingPiece.Move(toX, toY, curBoard)
				whitesMove = !whitesMove
				turn++
			} else {
				fmt.Printf("Can't Move Piece from %v%v to %v%v\n", string(coords[1][0]), fromY+1, string(coords[3][0]), toY+1)
			}
		}
	}
}

func main() {
	flag.BoolVar(&whiteAI, "blackPlayer", false, "Set the player to black")
	flag.IntVar(&maxDepth, "depth", 3, "Set the search depth for the AI")
	flag.Parse()

	blackAI = !whiteAI
	if maxDepth < 3 {
		maxDepth = 3
	}

	if whiteAI && !blackAI {
		fmt.Printf("Player is %v\nAI is %v\n\n", aurora.BgBlack(aurora.Gray("Black")), aurora.BgGray(aurora.Black("White")))
	} else if !whiteAI && blackAI {
		fmt.Printf("Player is %v\nAI is %v\n\n", aurora.BgGray(aurora.Black("White")), aurora.BgBlack(aurora.Gray("Black")))
	} else {
		fmt.Println("Someting is wrong, two players or two AI")
		os.Exit(0)
	}

	curBoard = NewBoard()

	for true {
		if curBoard.IsDead() {
			fmt.Println(curBoard.Winner(), "Winns!")
			os.Exit(0)
		} else if (whiteAI && whitesMove) || (blackAI && !whitesMove) {
			runAI()
		} else {
			runPlayer()
		}
	}
}

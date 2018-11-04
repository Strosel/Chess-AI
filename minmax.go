package main

import (
	"fmt"
	"math/rand"
	"os"
)

func minFun(b *Board, depth int) int {
	if depth >= maxDepth {
		b.SetScore()
		return b.Score
	}

	boards := b.GenerateNewBoardsWhitesTurn()
	// lowestBoardNo := 0
	lowestScore := 100000
	for i := 0; i < len(boards); i++ {
		if !boards[i].IsDead() {
			score, _ := maxFun(boards[i], depth+1)
			if score < lowestScore {
				// lowestBoardNo = i
				lowestScore = score
			}
		}
	}
	return lowestScore
}

func maxFun(b *Board, depth int) (int, *Board) {
	if depth >= maxDepth {
		b.SetScore()
		return b.Score, nil
	}

	boards := b.GenerateNewBoardsBlacksTurn()
	topBoardNo := 0
	topScore := -100000
	for i := 0; i < len(boards); i++ {
		score := minFun(boards[i], depth+1)
		if score > topScore {
			topBoardNo = i
			topScore = score
		}
	}

	if depth == 0 {
		if topBoardNo >= len(boards) {
			//fmt.Printf("Index %v out of range of %v\n", topBoardNo, len(boards))
			fmt.Println("No Moves")
			os.Exit(0)
		}
		return 0, boards[topBoardNo]
	}
	return topScore, nil
}

func minFunAB(b *Board, alpha, beta, depth int) (int, *Board) {
	if depth >= maxDepth {
		b.SetScore()
		return b.Score, nil
	}

	if b.IsDead() {
		if whiteAI && whitesMove {
			return 200, nil
		}
		if blackAI && !whitesMove {
			return -200, nil
		}
	}

	if b.HasWon() {
		if whiteAI && whitesMove {
			return -200, nil
		}
		if blackAI && !whitesMove {
			return 200, nil
		}
	}

	boards := b.GenerateNewBoardsWhitesTurn()
	lowestBoardNo := 0
	lowestScore := 300
	for i := 0; i < len(boards); i++ {

		score, _ := maxFunAB(boards[i], alpha, beta, depth+1)
		if score < lowestScore {
			lowestBoardNo = i
			lowestScore = score
		} else {
			if depth == 0 && score == lowestScore {
				if rand.Float64() < 0.3 {
					lowestBoardNo = i
				}
			}
		}
		if score < alpha {
			return lowestScore, nil
		}
		if score < beta {
			beta = score
		}

	}

	if depth == 0 {
		if lowestBoardNo >= len(boards) {
			//fmt.Printf("Index %v out of range of %v\n", lowestBoardNo, len(boards))
			fmt.Println("No Moves")
			os.Exit(0)
		}
		return 0, boards[lowestBoardNo]
	}
	return lowestScore, nil
}

func maxFunAB(b *Board, alpha, beta, depth int) (int, *Board) {
	if depth >= maxDepth {
		b.SetScore()
		return b.Score, nil
	}

	if b.IsDead() {
		if whiteAI && whitesMove {
			return 200, nil
		}
		if blackAI && !whitesMove {
			return -200, nil
		}
	}

	if b.HasWon() {
		if whiteAI && whitesMove {
			return -200, nil
		}
		if blackAI && !whitesMove {
			return 200, nil
		}
	}

	boards := b.GenerateNewBoardsBlacksTurn()
	topBoardNo := 0
	topScore := -300
	for i := 0; i < len(boards); i++ {

		score, _ := minFunAB(boards[i], alpha, beta, depth+1)
		if score > topScore {
			topBoardNo = i
			topScore = score
		} else {
			if depth == 0 && score == topScore {
				if rand.Float64() < 0.3 {
					topBoardNo = i
				}
			}
		}
		if score > beta {
			return topScore, nil
		}
		if score > alpha {
			alpha = score
		}

	}

	if depth == 0 {
		if topBoardNo >= len(boards) {
			// fmt.Printf("Index %v out of range of %v\n", topBoardNo, len(boards))
			fmt.Println("No Moves")
			os.Exit(0)
		}
		return 0, boards[topBoardNo]
	}
	return topScore, nil
}

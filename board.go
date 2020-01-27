package main

import (
	
)

const NUM_TOTAL_SPACES = 64
const ROW_LENGTH = 8
const EMPTY_ID = -1

type Game struct{
	DarwinId string
	Board [64]int16
}

type Options struct{
	OptionBoard [64]int16
}

func ( game Game ) getLegalMoves( playerId int16 ) [64]int16 {
	options := [64]int16{
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
	}
	SURROUNDING_SPACES := [8]int {
		-ROW_LENGTH - 1, -ROW_LENGTH, -ROW_LENGTH + 1,
		-1, +1,
		ROW_LENGTH - 1, ROW_LENGTH, ROW_LENGTH + 1,
	}
	enemyId := ( playerId + 1 ) % 2
	for i := 0; i < len( game.Board ); i++ {
		if game.Board[i] == playerId {
			for _, s_space := range SURROUNDING_SPACES {
				offsetDirected := s_space
				foundCapturablePieces := false
				for game.isIndexOnBoard( i + offsetDirected ) && game.Board[i + offsetDirected] == enemyId {
					offsetDirected += s_space
					foundCapturablePieces = true
				}
				if foundCapturablePieces && game.isIndexOnBoard( i + offsetDirected ) && game.Board[i + offsetDirected] == EMPTY_ID {
					options[i + offsetDirected] = playerId
				}
			}
		}
	}
	return options
}

func ( game Game ) getNextState( playerId int16, move int ) Game {
	SURROUNDING_SPACES := [8]int {
		-ROW_LENGTH - 1, -ROW_LENGTH, -ROW_LENGTH + 1,
		-1, +1,
		ROW_LENGTH - 1, ROW_LENGTH, ROW_LENGTH + 1,
	}
	board := game.Board
	board[move] = playerId
	enemyId := ( playerId + 1 ) % 2
	for _, s_space := range SURROUNDING_SPACES {
		offsetDirected := s_space
		foundCapturablePieces := false
		var capturableIndices []int
		for game.isIndexOnBoard( move + offsetDirected ) && game.Board[move + offsetDirected] == enemyId {
			foundCapturablePieces = true
			capturableIndices = append( capturableIndices, move + offsetDirected )
			offsetDirected += s_space
		}

		if foundCapturablePieces && game.isIndexOnBoard( move + offsetDirected ) && game.Board[move + offsetDirected] == playerId {
			for _, index := range capturableIndices {
				board[index] = playerId
			}
		}
	}
	return Game{ game.DarwinId, board }
}
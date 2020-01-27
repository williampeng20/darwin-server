package main

import (
	"math"
	"log"
)

func getDarwinId() string {
	return "darwin-placeholder-id"
}

func ( game Game ) minimax( depthRemaining int, isMax bool, playerId int16, originalPlayerId int16 ) float64 {
	if depthRemaining > 0 {
		return game.simpleScore( originalPlayerId )
	}

	optionBoard := game.getLegalMoves( playerId )
	var optionIndices []int
	for i := 0; i < len(optionBoard); i++ {
		if optionBoard[i] != EMPTY_ID {
			optionIndices = append( optionIndices, i )
		}
	}

	var rewardMap map[Game]float64
	rewardMap = make( map[Game]float64 )
	for optionIndex := range optionIndices {
		nextState := game.getNextState( playerId, optionIndex )
		reward := nextState.minimax( depthRemaining - 1, !isMax, (playerId + 1) % 2, originalPlayerId )
		rewardMap[nextState] = reward
	}

	if ( isMax ) {
		max := -999999.99
		for _, reward := range rewardMap {
			max = math.Max(max, reward)
		}
		return max
	} else {
		min := 999999.99
		for _, reward := range rewardMap {
			min = math.Min(min, reward)
		}
		return min
	}
}

func ( game Game ) simpleScore( playerId int16 ) float64 {
	score := 0.0
	for i := 0; i < len(game.Board); i++ {
		if game.Board[i] == playerId {
			score += 1
		} else if game.Board[i] != EMPTY_ID {
			score -= 1
		}
	}
	return score
}

const DEFAULT_TREE_DEPTH = 5

func ( game Game ) getDarwinMove( playerId int16 ) Game {
	darwinId := (playerId + 1) % 2
	darwinOptions := game.getLegalMoves( darwinId )
	var optionIndices []int
	for i := 0; i < len( darwinOptions ); i++ {
		if darwinOptions[i] != EMPTY_ID {
			optionIndices = append( optionIndices, i )
		}
	}
	log.Println(optionIndices)
	var rewardMap map[Game]float64
	rewardMap = make( map[Game]float64 )
	nextStateMap := make( map[Game]int )
	for _, optionIndex := range optionIndices {
		nextState := game.getNextState( darwinId, optionIndex )
		nextStateMap[nextState] = optionIndex
		rewardMap[nextState] = nextState.minimax( DEFAULT_TREE_DEPTH, false, darwinId, darwinId )
	}


	var bestGame Game
	max := -999999.99
	for game, reward := range rewardMap {
		if reward > max {
			bestGame = game
			max = reward
		}
	}
	log.Println("DARWIN PLACED INDEX")
	log.Println(nextStateMap[bestGame])
	log.Println("reward of option")
	log.Println(rewardMap[bestGame])
	return bestGame
}
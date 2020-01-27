package main

import (
	"net/http"
	"strconv"
)

func getGameParametersFromRequest( w http.ResponseWriter, r *http.Request ) ( rtnPlayerId int16, rtnDarwinId string, rtnBoard [64]int16, success bool ) {
	playerIdStr, playerIdExists := r.URL.Query()["playerId"]
	if !playerIdExists || len(playerIdStr) < 1 {
		http.Error(w, "Url Param 'playerId' Missing", http.StatusLengthRequired)
		return 0, "", [64]int16{}, false
	}
	playerId, err := strconv.ParseInt( playerIdStr[0], 10, 16)
	if err != nil {
		http.Error(w, "Url Param 'playerId' Malformed", http.StatusNotAcceptable)
		return 0, "", [64]int16{}, false
	}

	darwinId, darwinIdExists := r.URL.Query()["darwinId"]
	if !darwinIdExists || len(darwinId) < 1 {
		http.Error(w, "Url Param 'darwinId' Missing", http.StatusLengthRequired)
		return 0, "", [64]int16{}, false
	}

	boardStr, boardExists := r.URL.Query()["board"]
	if !boardExists || len(boardStr) != NUM_TOTAL_SPACES {
		http.Error(w, "Url Param 'board' Missing", http.StatusLengthRequired)
		return 0, "", [64]int16{}, false
	}
	var board [64]int16
	for i := 0; i < len(boardStr); i++ {
		id, err := strconv.ParseInt( boardStr[i], 10, 16 )
		if err != nil {
			http.Error(w, "Url Param 'board' Malformed", http.StatusNotAcceptable)
			return 0, "", [64]int16{}, false
		}
		board[i] = int16(id)
	}

	return int16(playerId), darwinId[0], board, true
}

func ( game Game ) isIndexOnBoard( index int ) bool {
	return index >= 0 && index < len( game.Board )
}
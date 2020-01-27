package main

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
)

func getPlayerOptions( w http.ResponseWriter, r *http.Request ) {
	playerId, darwinId, board, success := getGameParametersFromRequest( w, r )
	if !success {
		return
	}
	gameState := Game{ darwinId, board }

	options := Options{ gameState.getLegalMoves( playerId ) }
	js, err := json.Marshal( options )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set( "Content-Type", "application/json" )
	w.Header().Set( "Access-Control-Allow-Origin", "*" )
	w.Write(js)
}

func getDarwinMove( w http.ResponseWriter, r *http.Request ) {
	playerId, darwinId, board, success := getGameParametersFromRequest( w, r )
	if !success {
		return
	}
	gameState := Game{ darwinId, board }

	darwinMove := gameState.getDarwinMove( playerId )
	log.Println(darwinMove.Board)
	js, err := json.Marshal( darwinMove )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set( "Content-Type", "application/json" )
	w.Header().Set( "Access-Control-Allow-Origin", "*" )
	w.Write(js)
}

func getPlayerCapturedBoard( w http.ResponseWriter, r *http.Request ) {
	playerId, darwinId, board, success := getGameParametersFromRequest( w, r )
	if !success {
		return
	}
	gameState := Game{ darwinId, board }

	moveStr, moveExists := r.URL.Query()["move"]
	if !moveExists || len(moveStr) < 1 {
		http.Error(w, "Url Param 'move' Missing", http.StatusLengthRequired)
		return
	}
	move, err := strconv.ParseInt( moveStr[0], 10, 16)
	if err != nil {
		http.Error(w, "Url Param 'move' Malformed", http.StatusNotAcceptable)
		return
	}

	nextState := gameState.getNextState( playerId, int(move) )
	js, err := json.Marshal( nextState )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set( "Content-Type", "application/json" )
	w.Header().Set( "Access-Control-Allow-Origin", "*" )
	w.Write(js)
}

func getNewGame( w http.ResponseWriter, r *http.Request ) {
	game := Game{ getDarwinId(), [64]int16{
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1,  1,  0, -1, -1, -1,
		-1, -1, -1,  0,  1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, -1, -1,
	}}

	js, err := json.Marshal( game )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set( "Content-Type", "application/json" )
	w.Header().Set( "Access-Control-Allow-Origin", "*" )
	w.Write(js)
}

func main() {
	http.HandleFunc( "/playerOptions", getPlayerOptions )
	http.HandleFunc( "/darwinMove", getDarwinMove )
	http.HandleFunc( "/playerCaptured", getPlayerCapturedBoard )
	http.HandleFunc( "/newGame", getNewGame )
	log.Fatal(http.ListenAndServe(":8080", nil))
}

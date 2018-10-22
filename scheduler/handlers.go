package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/dbops"
)


func vidDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	fmt.Println("receive vid: ", vid)
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should no be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}
	sendResponse(w, 200, "delete video successfully")
	return
}
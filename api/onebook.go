package handler

import "net/http"

func OneBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("One Book Handler"))
}

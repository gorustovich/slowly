package app

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"time"
)

type SlowHandler struct {
	Processor SlowProcessor
}

func NewSlowHandler(processor SlowProcessor) *SlowHandler {
	return &SlowHandler{
		Processor: processor,
	}
}

type slowRequest struct {
	Timeout *int
}

func (sh *SlowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NewErrorResponse("error with getting request body", http.StatusBadRequest).Write(w)
		return
	}
	if len(body) == 0 {
		NewErrorResponse("post body is required", http.StatusBadRequest).Write(w)
		return
	}
	req := slowRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		NewErrorResponse(fmt.Sprintf("error with marshaling post body: %s", err), http.StatusBadRequest).Write(w)
		return
	}
	if req.Timeout == nil {
		NewErrorResponse("timeout in body is required", http.StatusBadRequest).Write(w)
		return
	}
	err = sh.Processor.Process(time.Millisecond * time.Duration(*req.Timeout))
	if err != nil {
		NewErrorResponse(err.Error(), http.StatusBadRequest).Write(w)
		return
	}
	NewSuccessResponse("ok").Write(w)
}

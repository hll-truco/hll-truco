package worker

import "github.com/hll-truco/hll-truco/hll-dist-http/root/state"

type UpdateRequest struct {
	Gob string `json:"gob"`
}

func SendUpdateRequest(baseURL string, gobString string) {
	url := baseURL + "/update"
	// Create the UpdateRequest struct
	update := UpdateRequest{Gob: gobString}
	sendPOSTJsonData(url, update)
}

func SendReportRequest(baseURL string, report state.WorkerResult) {
	url := baseURL + "/report"
	sendPOSTJsonData(url, report)
}

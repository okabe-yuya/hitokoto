package p

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/okabe-yuya/hitokoto/member"
	"github.com/slack-go/slack"
)


func httpServer(w http.ResponseWriter, r *http.Request) {
	member := member.Get()
	if len(member) > 0 {
		params := &slack.Msg{Text: toResponse(member), ResponseType: "in_channel"}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		w.Write([]byte("おや！何か手違いがありました！ もう一度試してください😭"))
		w.WriteHeader(http.StatusBadRequest)
	}
}

func toResponse(member []string) string {
	resp := ""
	for i, m := range member {
		resp += strconv.Itoa(i+1) + "番目: " + m + "\n"
	}
	return resp
}

package browser

import (
	"log"
	"net/http"
)

type StaticServer struct {
	Addr string
	Data []byte
}

//启动服务端
func (s *StaticServer) Serve() {
	http.HandleFunc("/", s.handle)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		panic(err)
	}
}

func (s *StaticServer) handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(s.Data)
	if err != nil {
		log.Println(err)
	}
}

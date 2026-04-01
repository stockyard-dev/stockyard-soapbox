package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-soapbox/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){q:=r.URL.Query().Get("q");list,_:=s.db.List(q);if list==nil{list=[]store.Question{}};writeJSON(w,200,list)}
func(s *Server)handleAsk(w http.ResponseWriter,r *http.Request){var q store.Question;json.NewDecoder(r.Body).Decode(&q);if q.Title==""{writeError(w,400,"title required");return};s.db.Ask(&q);writeJSON(w,201,q)}
func(s *Server)handleVote(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.VoteQuestion(id);writeJSON(w,200,map[string]string{"status":"voted"})}
func(s *Server)handleListAnswers(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListAnswers(id);if list==nil{list=[]store.Answer{}};writeJSON(w,200,list)}
func(s *Server)handleAnswer(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var a store.Answer;json.NewDecoder(r.Body).Decode(&a);a.QuestionID=id;if a.Body==""{writeError(w,400,"body required");return};s.db.Answer(&a);writeJSON(w,201,a)}
func(s *Server)handleAccept(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Accept(id);writeJSON(w,200,map[string]string{"status":"accepted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}

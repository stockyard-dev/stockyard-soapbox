package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Question struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
	Tags string `json:"tags"`
	Votes int `json:"votes"`
	AnswerCount int `json:"answer_count"`
	Accepted int `json:"accepted"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"soapbox.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS questions(id TEXT PRIMARY KEY,title TEXT NOT NULL,body TEXT DEFAULT '',author TEXT DEFAULT '',tags TEXT DEFAULT '',votes INTEGER DEFAULT 0,answer_count INTEGER DEFAULT 0,accepted INTEGER DEFAULT 0,status TEXT DEFAULT 'open',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Question)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO questions(id,title,body,author,tags,votes,answer_count,accepted,status,created_at)VALUES(?,?,?,?,?,?,?,?,?,?)`,e.ID,e.Title,e.Body,e.Author,e.Tags,e.Votes,e.AnswerCount,e.Accepted,e.Status,e.CreatedAt);return err}
func(d *DB)Get(id string)*Question{var e Question;if d.db.QueryRow(`SELECT id,title,body,author,tags,votes,answer_count,accepted,status,created_at FROM questions WHERE id=?`,id).Scan(&e.ID,&e.Title,&e.Body,&e.Author,&e.Tags,&e.Votes,&e.AnswerCount,&e.Accepted,&e.Status,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Question{rows,_:=d.db.Query(`SELECT id,title,body,author,tags,votes,answer_count,accepted,status,created_at FROM questions ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Question;for rows.Next(){var e Question;rows.Scan(&e.ID,&e.Title,&e.Body,&e.Author,&e.Tags,&e.Votes,&e.AnswerCount,&e.Accepted,&e.Status,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Update(e *Question)error{_,err:=d.db.Exec(`UPDATE questions SET title=?,body=?,author=?,tags=?,votes=?,answer_count=?,accepted=?,status=? WHERE id=?`,e.Title,e.Body,e.Author,e.Tags,e.Votes,e.AnswerCount,e.Accepted,e.Status,e.ID);return err}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM questions WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM questions`).Scan(&n);return n}

func(d *DB)Search(q string, filters map[string]string)[]Question{
    where:="1=1"
    args:=[]any{}
    if q!=""{
        where+=" AND (title LIKE ? OR body LIKE ?)"
        args=append(args,"%"+q+"%");args=append(args,"%"+q+"%");
    }
    if v,ok:=filters["status"];ok&&v!=""{where+=" AND status=?";args=append(args,v)}
    rows,_:=d.db.Query(`SELECT id,title,body,author,tags,votes,answer_count,accepted,status,created_at FROM questions WHERE `+where+` ORDER BY created_at DESC`,args...)
    if rows==nil{return nil};defer rows.Close()
    var o []Question;for rows.Next(){var e Question;rows.Scan(&e.ID,&e.Title,&e.Body,&e.Author,&e.Tags,&e.Votes,&e.AnswerCount,&e.Accepted,&e.Status,&e.CreatedAt);o=append(o,e)};return o
}

func(d *DB)Stats()map[string]any{
    m:=map[string]any{"total":d.Count()}
    rows,_:=d.db.Query(`SELECT status,COUNT(*) FROM questions GROUP BY status`)
    if rows!=nil{defer rows.Close();by:=map[string]int{};for rows.Next(){var s string;var c int;rows.Scan(&s,&c);by[s]=c};m["by_status"]=by}
    return m
}

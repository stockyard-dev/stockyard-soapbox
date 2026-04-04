package server
import "net/http"
func (s *Server) dashboard(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/html"); w.Write([]byte(dashHTML)) }
const dashHTML = `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>Soapbox</title><link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;700&display=swap" rel="stylesheet"><style>:root{--bg:#1a1410;--bg2:#241e18;--bg3:#2e261e;--rust:#e8753a;--leather:#a0845c;--cream:#f0e6d3;--cd:#bfb5a3;--cm:#7a7060;--gold:#d4a843;--green:#4a9e5c;--red:#c94444;--blue:#5b8dd9;--mono:'JetBrains Mono',monospace}*{margin:0;padding:0;box-sizing:border-box}body{background:var(--bg);color:var(--cream);font-family:var(--mono);line-height:1.5}.hdr{padding:1rem 1.5rem;border-bottom:1px solid var(--bg3);display:flex;justify-content:space-between;align-items:center}.hdr h1{font-size:.9rem;letter-spacing:2px}.hdr h1 span{color:var(--rust)}.main{padding:1.5rem;max-width:800px;margin:0 auto}.stats{display:grid;grid-template-columns:repeat(3,1fr);gap:.5rem;margin-bottom:1rem}.st{background:var(--bg2);border:1px solid var(--bg3);padding:.6rem;text-align:center}.st-v{font-size:1.2rem;font-weight:700}.st-l{font-size:.5rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-top:.15rem}.toolbar{display:flex;gap:.5rem;margin-bottom:1rem}.search{flex:1;padding:.4rem .6rem;background:var(--bg2);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}.search:focus{outline:none;border-color:var(--leather)}.q{background:var(--bg2);border:1px solid var(--bg3);padding:.8rem;margin-bottom:.5rem;display:flex;gap:.8rem;transition:border-color .2s}.q:hover{border-color:var(--leather)}.vote-col{display:flex;flex-direction:column;align-items:center;min-width:40px}.vote-btn{background:none;border:none;color:var(--cm);cursor:pointer;font-size:1rem;padding:0;line-height:1}.vote-btn:hover{color:var(--rust)}.vote-count{font-size:.9rem;font-weight:700}.q-content{flex:1}.q-title{font-size:.85rem;font-weight:700}.q-body{font-size:.7rem;color:var(--cd);margin-top:.2rem;display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden}.q-meta{font-size:.55rem;color:var(--cm);margin-top:.3rem;display:flex;gap:.5rem;flex-wrap:wrap;align-items:center}.q-actions{display:flex;gap:.3rem;flex-shrink:0;align-self:flex-start}.badge{font-size:.5rem;padding:.12rem .35rem;text-transform:uppercase;letter-spacing:1px;border:1px solid}.badge.open{border-color:var(--green);color:var(--green)}.badge.answered{border-color:var(--blue);color:var(--blue)}.badge.closed{border-color:var(--cm);color:var(--cm)}.tag{font-size:.45rem;padding:.1rem .25rem;background:var(--bg3);color:var(--cm)}.btn{font-size:.6rem;padding:.25rem .5rem;cursor:pointer;border:1px solid var(--bg3);background:var(--bg);color:var(--cd);transition:all .2s}.btn:hover{border-color:var(--leather);color:var(--cream)}.btn-p{background:var(--rust);border-color:var(--rust);color:#fff}.btn-sm{font-size:.55rem;padding:.2rem .4rem}.modal-bg{display:none;position:fixed;inset:0;background:rgba(0,0,0,.65);z-index:100;align-items:center;justify-content:center}.modal-bg.open{display:flex}.modal{background:var(--bg2);border:1px solid var(--bg3);padding:1.5rem;width:500px;max-width:92vw;max-height:90vh;overflow-y:auto}.modal h2{font-size:.8rem;margin-bottom:1rem;color:var(--rust)}.fr{margin-bottom:.6rem}.fr label{display:block;font-size:.55rem;color:var(--cm);text-transform:uppercase;letter-spacing:1px;margin-bottom:.2rem}.fr input,.fr select,.fr textarea{width:100%;padding:.4rem .5rem;background:var(--bg);border:1px solid var(--bg3);color:var(--cream);font-family:var(--mono);font-size:.7rem}.fr input:focus,.fr textarea:focus{outline:none;border-color:var(--leather)}.row2{display:grid;grid-template-columns:1fr 1fr;gap:.5rem}.acts{display:flex;gap:.4rem;justify-content:flex-end;margin-top:1rem}.empty{text-align:center;padding:3rem;color:var(--cm);font-style:italic;font-size:.75rem}</style></head><body>
<div class="hdr"><h1><span>&#9670;</span> SOAPBOX</h1><button class="btn btn-p" onclick="openForm()">+ Ask Question</button></div>
<div class="main"><div class="stats" id="stats"></div><div class="toolbar"><input class="search" id="search" placeholder="Search questions..." oninput="render()"></div><div id="list"></div></div>
<div class="modal-bg" id="mbg" onclick="if(event.target===this)closeModal()"><div class="modal" id="mdl"></div></div>
<script>
var A='/api',items=[],editId=null;
async function load(){var r=await fetch(A+'/questions').then(function(r){return r.json()});items=r.questions||[];renderStats();render();}
function renderStats(){var t=items.length,open=items.filter(function(q){return q.status==='open'}).length,totalVotes=items.reduce(function(s,q){return s+(q.votes||0)},0);
document.getElementById('stats').innerHTML='<div class="st"><div class="st-v">'+t+'</div><div class="st-l">Questions</div></div><div class="st"><div class="st-v" style="color:var(--green)">'+open+'</div><div class="st-l">Open</div></div><div class="st"><div class="st-v">'+totalVotes+'</div><div class="st-l">Votes</div></div>';}
function render(){var q=(document.getElementById('search').value||'').toLowerCase();var f=items.slice();
if(q)f=f.filter(function(i){return(i.title||'').toLowerCase().includes(q)||(i.body||'').toLowerCase().includes(q)||(i.tags||'').toLowerCase().includes(q)});
f.sort(function(a,b){return(b.votes||0)-(a.votes||0)});
if(!f.length){document.getElementById('list').innerHTML='<div class="empty">No questions yet.</div>';return;}
var h='';f.forEach(function(i){
h+='<div class="q"><div class="vote-col"><button class="vote-btn" onclick="upvote(''+i.id+'')">&#9650;</button><div class="vote-count">'+(i.votes||0)+'</div></div>';
h+='<div class="q-content"><div class="q-title">'+esc(i.title)+'</div>';
if(i.body)h+='<div class="q-body">'+esc(i.body)+'</div>';
h+='<div class="q-meta">';
if(i.status)h+='<span class="badge '+(i.status||'open')+'">'+esc(i.status)+'</span>';
if(i.author)h+='<span>'+esc(i.author)+'</span>';
if(i.answer_count)h+='<span>'+i.answer_count+' answers</span>';
if(i.tags){i.tags.split(',').forEach(function(t){t=t.trim();if(t)h+='<span class="tag">#'+esc(t)+'</span>';});}
h+='<span>'+ft(i.created_at)+'</span>';
h+='</div></div><div class="q-actions"><button class="btn btn-sm" onclick="openEdit(''+i.id+'')">Edit</button><button class="btn btn-sm" onclick="del(''+i.id+'')" style="color:var(--red)">&#10005;</button></div></div>';});
document.getElementById('list').innerHTML=h;}
async function upvote(id){var q=null;for(var j=0;j<items.length;j++){if(items[j].id===id){q=items[j];break;}}if(!q)return;
await fetch(A+'/questions/'+id,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify({votes:(q.votes||0)+1})});load();}
async function del(id){if(!confirm('Delete?'))return;await fetch(A+'/questions/'+id,{method:'DELETE'});load();}
function formHTML(question){var i=question||{title:'',body:'',author:'',tags:'',status:'open'};var isEdit=!!question;
var h='<h2>'+(isEdit?'EDIT':'ASK A')+' QUESTION</h2>';
h+='<div class="fr"><label>Title *</label><input id="f-title" value="'+esc(i.title)+'" placeholder="What do you want to know?"></div>';
h+='<div class="fr"><label>Details</label><textarea id="f-body" rows="4">'+esc(i.body)+'</textarea></div>';
h+='<div class="row2"><div class="fr"><label>Author</label><input id="f-author" value="'+esc(i.author)+'"></div><div class="fr"><label>Tags</label><input id="f-tags" value="'+esc(i.tags)+'" placeholder="comma separated"></div></div>';
h+='<div class="acts"><button class="btn" onclick="closeModal()">Cancel</button><button class="btn btn-p" onclick="submit()">'+(isEdit?'Save':'Post')+'</button></div>';return h;}
function openForm(){editId=null;document.getElementById('mdl').innerHTML=formHTML();document.getElementById('mbg').classList.add('open');}
function openEdit(id){var q=null;for(var j=0;j<items.length;j++){if(items[j].id===id){q=items[j];break;}}if(!q)return;editId=id;document.getElementById('mdl').innerHTML=formHTML(q);document.getElementById('mbg').classList.add('open');}
function closeModal(){document.getElementById('mbg').classList.remove('open');editId=null;}
async function submit(){var title=document.getElementById('f-title').value.trim();if(!title){alert('Title required');return;}
var body={title:title,body:document.getElementById('f-body').value.trim(),author:document.getElementById('f-author').value.trim(),tags:document.getElementById('f-tags').value.trim()};
if(editId){await fetch(A+'/questions/'+editId,{method:'PUT',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}
else{body.votes=0;body.status='open';await fetch(A+'/questions',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(body)});}closeModal();load();}
function ft(t){if(!t)return'';try{return new Date(t).toLocaleDateString('en-US',{month:'short',day:'numeric'})}catch(e){return t;}}
function esc(s){if(!s)return'';var d=document.createElement('div');d.textContent=s;return d.innerHTML;}
document.addEventListener('keydown',function(e){if(e.key==='Escape')closeModal();});load();
</script></body></html>`

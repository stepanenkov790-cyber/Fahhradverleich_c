const api=async(path,options={})=>{const r=await fetch(path,{headers:{"Content-Type":"application/json"},...options});if(!r.ok)throw new Error(await r.text());return r.status===204?null:r.json()};
const $=id=>document.getElementById(id), data=f=>Object.fromEntries(new FormData(f).entries()), esc=s=>String(s).replace(/[&<>"']/g,c=>({"&":"&amp;","<":"&lt;",">":"&gt;",'"':"&quot;","'":"&#39;"}[c]));
let bikes=[],customers=[];
function table(head,rows){return rows.length?`<table><thead><tr>${head.map(x=>`<th>${x}</th>`).join("")}</tr></thead><tbody>${rows.join("")}</tbody></table>`:"<p>Noch keine Einträge vorhanden.</p>"}
async function refresh(){const [d,b,c,r,v]=await Promise.all([api("/api/dashboard"),api("/api/bikes"),api("/api/customers"),api("/api/rentals"),api("/api/reservations")]);bikes=b;customers=c;
$("available").textContent=d.availableBikes;$("rented").textContent=d.rentedBikes;$("total").textContent=d.totalBikes;$("customersCount").textContent=d.customers;
$("bikeList").innerHTML=table(["ID","Name","Typ","Größe","Preis/Tag","Status",""],b.map(x=>`<tr><td>${x.id}</td><td>${esc(x.name)}</td><td>${esc(x.type)}</td><td>${esc(x.size)}</td><td>€ ${x.dailyPrice.toFixed(2)}</td><td>${esc(x.status)}</td><td><button class="delete" onclick="deleteBike(${x.id})">Löschen</button></td></tr>`));
$("customerList").innerHTML=table(["ID","Name","E-Mail",""],c.map(x=>`<tr><td>${x.id}</td><td>${esc(x.name)}</td><td>${esc(x.email)}</td><td><button class="delete" onclick="deleteCustomer(${x.id})">Löschen</button></td></tr>`));
const opts=a=>a.map(x=>`<option value="${x.id}">${esc(x.name)} (#${x.id})</option>`).join("");
$("rentalBike").innerHTML=opts(b.filter(x=>x.status==="available"));$("reservationBike").innerHTML=opts(b);$("rentalCustomer").innerHTML=opts(c);$("reservationCustomer").innerHTML=opts(c);
$("rentalList").innerHTML=table(["ID","Fahrrad-ID","Kunden-ID","Datum","Tage","Preis/Tag","Gesamt"],r.map(x=>`<tr><td>${x.id}</td><td>${x.bikeId}</td><td>${x.customerId}</td><td>${esc(x.date)}</td><td>${x.days}</td><td>€ ${x.dailyPrice.toFixed(2)}</td><td>€ ${x.totalPrice.toFixed(2)}</td></tr>`));
$("reservationList").innerHTML=table(["ID","Fahrrad-ID","Kunden-ID","Datum","Status"],v.map(x=>`<tr><td>${x.id}</td><td>${x.bikeId}</td><td>${x.customerId}</td><td>${esc(x.date)}</td><td>${esc(x.status)}</td></tr>`));}
function bind(id,path,convert){$(id).addEventListener("submit",async e=>{e.preventDefault();try{const v=data(e.target);if(convert)convert(v);await api(path,{method:"POST",body:JSON.stringify(v)});e.target.reset();await refresh()}catch(err){alert(err.message)}})}
bind("bikeForm","/api/bikes",v=>v.dailyPrice=Number(v.dailyPrice));bind("customerForm","/api/customers");
bind("rentalForm","/api/rentals",v=>{v.bikeId=Number(v.bikeId);v.customerId=Number(v.customerId);v.days=Number(v.days)});
bind("reservationForm","/api/reservations",v=>{v.bikeId=Number(v.bikeId);v.customerId=Number(v.customerId)});
window.deleteBike=async id=>{if(confirm("Fahrrad wirklich löschen?")){try{await api(`/api/bikes/${id}`,{method:"DELETE"});await refresh()}catch(e){alert(e.message)}}};
window.deleteCustomer=async id=>{if(confirm("Kunde wirklich löschen?")){try{await api(`/api/customers/${id}`,{method:"DELETE"});await refresh()}catch(e){alert(e.message)}}};
refresh().catch(e=>alert("Backend nicht erreichbar: "+e.message));

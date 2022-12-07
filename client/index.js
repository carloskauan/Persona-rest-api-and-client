const express = require("express")
const app = express();
const axios = require('axios').default;
const port = process.env.PORT || 8080

app.set("view engine", "ejs");
app.use(express.json())
app.use(express.urlencoded({ extended: true}))

app.get("/", (req, res)=>{
  axios.get("http://localhost:7070/api/persona")
        .then(response=>{
          res.render("index.ejs",{data: response.data})
        }).catch(err=>{
          console.error("Erro ao relizar requisição", err);
        });
});

app.post("/find", (req, res)=>{
  let id = req.body.id
  axios.get(`http://localhost:7070/api/persona/${req.body.id}`)
        .then(persona=>{
          res.render("find.ejs",{persona: persona.data});
        });
});

app.get("/register", (req, res)=>{res.render("register.ejs")});
app.post("/sendregister", (req, res)=>{
  axios.post("http://localhost:7070/api/persona",{
    name: req.body.name, history: req.body.history
  }).then(response=>{
    console.log("Personalidade registrada com sucesso");
    res.render("find.ejs", {persona: response.data});
  }).catch(err=>{console.error("Erro ao registar Personalidade", err);})
})

app.get("/edit/:id", (req, res)=>{
  axios.get(`http://localhost:7070/api/persona/${req.params.id}`)
        .then(response=>{
          res.render("edit.ejs",{persona: response.data})
        }).catch(err=>{
          console.error("Erro ao relizar requisição", err);
        });
});
app.post("/sendedit", (req, res)=>{
  axios.put(`http://localhost:7070/api/persona/${req.body.id}`, {
    name: req.body.name, history: req.body.history
  }).then(response=>{
    console.log("Registro editado com sucesso");
    res.render("find.ejs", {persona: response.data})
  }).catch(err=>{console.error("Erro ao editar registr", err);})
});

app.post("/delete", (req, res)=>{
  axios.delete(`http://localhost:7070/api/persona/${req.body.id}`)
        .then(()=>{res.redirect("/");})
});

app.listen(port, ()=>{console.log("Cliente rodando na porta 8080")});

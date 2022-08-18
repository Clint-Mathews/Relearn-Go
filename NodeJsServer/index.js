// import express (after npm install express)
const express = require('express');

// create new express app and save it as "app"
const app = express();

// server configuration
const PORT = 9088;

app.use(express.json())
app.use(express.urlencoded({extended:true}))

// create a route for the app
app.get('/', (req, res) => {
  res.status(200).send('Hello World');
});
app.get('/get', (req, res) => {
    res.status(200).send({message:"Hello"});
  });

app.post('/post', (req, res) => {
    res.status(200).send(req.body);
  });

  app.post('/postForm', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
  });

// make the server listen to requests
app.listen(PORT, () => {
  console.log(`Server running at: http://localhost:${PORT}/`);
});

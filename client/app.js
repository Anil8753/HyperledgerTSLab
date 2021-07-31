const express = require('express')
const bodyParser = require('body-parser');

const rgGet= require("./registeration_data_get");
const rgSet= require("./registeration_data_set");

const svGet= require("./service_data_get");
const svSet= require("./service_data_set");

const insGet= require("./insurance_data_get");
const insSet= require("./insurance_data_set");


const app = express()
const port = 3000

app.use(bodyParser.json());
app.use(express.static('www'))

//----------------------------------------------------------------------
app.get('/api/registration/:id', (req, res) => {
  console.log(req.params['id'])
  rgGet.get_registration_data(req.params['id'], false)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.get('/api/registration_history/:id', (req, res) => {
  console.log(req.params['id'])
  rgGet.get_registration_data(req.params['id'], true)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.post('/api/registration', (req, res) => {
  console.log(req.body);
  const record = req.body;
  rgSet.set_registration_data(record)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});
//----------------------------------------------------------------------
//----------------------------------------------------------------------

app.get('/api/service/:id', (req, res) => {
  console.log(req.params['id'])
  svGet.get_service_data(req.params['id'], false)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.get('/api/service_history/:id', (req, res) => {
  console.log(req.params['id'])
  svGet.get_service_data(req.params['id'], true)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.post('/api/service', (req, res) => {
  console.log(req.body);
  const record = req.body;
  svSet.set_service_data(record)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});
//----------------------------------------------------------------------
//----------------------------------------------------------------------

//----------------------------------------------------------------------
//----------------------------------------------------------------------

app.get('/api/insurance/:id', (req, res) => {
  console.log(req.params['id'])
  insGet.get_insurance_data(req.params['id'], false)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.get('/api/insurance_history/:id', (req, res) => {
  console.log(req.params['id'])
  insGet.get_insurance_data(req.params['id'], true)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});

app.post('/api/insurance', (req, res) => {
  console.log(req.body);
  const record = req.body;
  insSet.set_insurance_data(record)
  .then(resp=>{
    res.send(resp);
  }).catch(err=>{ 
    next(err)
  });
});
//----------------------------------------------------------------------
//----------------------------------------------------------------------


app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
});



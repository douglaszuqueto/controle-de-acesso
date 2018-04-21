'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  CDA_API_HOST : process.env.CDA_API_HOST ? '"http://' + process.env.CDA_API_HOST +  '/api/"' : '"http://127.0.0.1:3000/api/"',
  CDA_MQTT_HOST : process.env.CDA_API_HOST ? '"ws://' + process.env.CDA_MQTT_HOST + '"' : '"ws://127.0.0.1:8083"'
})

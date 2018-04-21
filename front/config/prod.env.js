'use strict'
module.exports = {
  NODE_ENV: '"production"',
  CDA_API_HOST : process.env.CDA_API_HOST ? '"http://' + process.env.CDA_API_HOST +  '/api/"' : '"http://127.0.0.1:3000/api/"',
  CDA_MQTT_HOST : process.env.CDA_API_HOST ? '"ws://' + process.env.CDA_MQTT_HOST + '"' : '"ws://127.0.0.1:8083"'
}

<template>
  <div class="row">
    <div class="col-12 col-md-10">
      <div class="row card-deck">
        <div class="col-12 col-sm-6 col-md-4" v-for="(device) in devices" v-bind:key="device.ID">
          <div :id="'device-' + device.ChipID" class="card" v-bind:class="device.Class.border"
               style="max-width: 18rem;">
            <div class="card-header text-white" v-bind:class="device.Class.bg">
              {{device.Name}}
              <router-link :to="{name: 'device.update', params: {id: device.ID}}" class="float-right text-white">
                <i class="fas fa-server"></i>
              </router-link>
            </div>
            <div class="card-body text-center">
              <img :src="device.Gravatar">
            </div>
            <div class="card-footer text-center text-dark bg-transparent">{{device.User}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="col-12 col-md-2">
      <div class="row text-white">
        <h6 class="col-12 bg-dark pad">
          <i class="fas fa-check"></i> Desativado</h6>
        <h6 class="col-12 bg-success pad">
          <i class="fas fa-check"></i> Ativo e Conectado</h6>
        <h6 class="col-12 bg-warning pad">
          <i class="fas fa-check"></i> Desconectado</h6>
        <h6 class="col-12 bg-info pad">
          <i class="fas fa-check"></i> Acesso autorizado</h6>
        <h6 class="col-12 bg-danger pad">
          <i class="fas fa-check"></i> Acesso negado</h6>
      </div>
    </div>

  </div>
</template>

<script>
import {http} from '@/plugins/http'
import mqtt from 'mqtt'

const mqttStateTopic = '/device/state/#'
const mqttTagTopic = '/device/verify'
const defaultGravatar = '/static/img/gravatar.png'
const defaultUser = '-'

export default {
  beforeMount: function () {
    this.initMQTT()
  },
  beforeDestroy: function () {
    this.closeMQTT()
  },
  mounted: function () {
    this.getDevices()
  },
  data: () => {
    return {
      devices: [],
      classes: {
        'success': {border: 'border-success', bg: 'bg-success'},
        'danger': {border: 'border-danger', bg: 'bg-danger'},
        'warning': {border: 'border-warning', bg: 'bg-warning'},
        'dark': {border: 'border-dark', bg: 'bg-dark'},
        'info': {border: 'border-info', bg: 'bg-info'}
      }
    }
  },
  methods: {
    getDevices: async function () {
      const devices = await http.get('/device')
        .then((data) => data.data.Data)

      Object.keys(devices).forEach((key) => {
        const item = devices[key]
        item.ChipID = item.ChipID.String
        item.User = defaultUser
        item.Gravatar = defaultGravatar
        if (item.State === 1) {
          item.Class = this.classes['success']
        } else {
          item.Class = this.classes['dark']
        }
      })
      this.devices = devices
    },
    changeState: async function (chipId, state) {
      let self = this
      Object.keys(this.devices).forEach((key) => {
        const item = self.devices[key]
        if (item.ChipID !== chipId) {
          return
        }
        if (state === 'online') {
          item.Class = this.classes['success']
          item.State = 1
        } else {
          item.Class = this.classes['warning']
          item.State = 0
        }
      })
    },
    changeStateByTag: async function (payload) {
      let self = this
      Object.keys(this.devices).forEach((key) => {
        const item = self.devices[key]
        if (item.ChipID !== payload.chip_id || item.State === 0) {
          return
        }
        item.Gravatar = payload.gravatar
        item.User = payload.user
        if (payload.state === '1') {
          item.Class = self.classes['info']
        } else {
          item.Class = self.classes['danger']
        }
        setTimeout(() => {
          item.Class = self.classes['success']
          item.Gravatar = defaultGravatar
          item.User = defaultUser
        }, 1000)
      })
    },
    initMQTT: function () {
      const self = this
      console.log('[MQTT] Iniciando conexão...')
      this.client = mqtt.connect(process.env.CDA_MQTT_HOST, {
        username: 'dashboard',
        password: 'dashboard'
      })

      this.client.on('connect', function () {
        self.client.subscribe(mqttStateTopic)
        self.client.subscribe(mqttTagTopic)
      })

      this.client.on('message', function (topic, message) {
        const msg = message.toString()
        if (topic === mqttTagTopic) {
          const payload = JSON.parse(msg)
          self.changeStateByTag(payload)
        } else {
          const t = topic.split('/')
          const chipID = t[t.length - 1]
          self.changeState(chipID, msg)
        }
      })
    },
    closeMQTT: function () {
      console.log('[MQTT] Fechando conexão')
      this.client.end()
    }
  }
}
</script>

<style scoped>
  img {
    border-radius: 100px;
    width: 100px;
  }
  .pad{
    padding: 10px
  }
</style>

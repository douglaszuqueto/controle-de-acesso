<template>
  <div class="">
    <h3 class="">Create device</h3>
    <hr>

    <form class="" method="post">
      <div class="form-group row">
        <label for="name" class="col-sm-1 col-form-label">Name</label>
        <div class="col-sm-5">
          <input type="text" class="form-control" v-model="form.name" id="name" placeholder="Name">
        </div>
        <label for="chip_id" class="col-sm-1 col-form-label">Chip ID</label>
        <div class="col-sm-5">
          <input type="text" class="form-control" v-model="form.chip_id" id="chip_id" placeholder="Chip ID">
        </div>
      </div>

      <div class="form-group row">
        <label for="description" class="col-sm-1 col-form-label">Description</label>
        <div class="col-sm-5">
          <textarea class="form-control" v-model="form.description" id="description" rows="3"></textarea>
        </div>

        <div class="col-sm-1">State</div>
        <div class="col-sm-5">
          <div class="custom-control custom-checkbox">
            <input class="custom-control-input" type="checkbox" v-model="form.state" id="state" value="0">
            <label class="custom-control-label" for="state">
            </label>
          </div>
        </div>
      </div>

      <div class="form-group text-right">
        <button type="button" class="btn btn-primary" onclick="history.go(-1)" title="Return">
          <i class="fas fa-arrow-circle-left"></i> Return
        </button>
        <button type="submit" class="btn btn-success" @click.prevent="save()">
          <i class="fa fa-save"></i> Save
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { http } from '@/plugins/http'
import swal from 'sweetalert2'
import mqtt from 'mqtt'

export default {
  beforeMount: function () {
    this.initMQTT()
  },
  beforeDestroy: function () {
    this.closeMQTT()
  },
  data: () => {
    return {
      form: {
        name: 'Device 02',
        description: 'Descricao device 02',
        state: 1,
        chip_id: ''
      },
      client: null
    }
  },
  methods: {
    save: async function () {
      try {
        const {data} = await http.post('/device', this.form)
        const {Error} = data
        if (Error.Error) {
          throw Error.Message
        }
        this.reset()
        swal('Sucesso!', 'Device cadastrado com sucesso', 'success')
      } catch (err) {
        if (err instanceof Object) {
          swal('Erro!', err.message, 'error')
          return
        }
        console.log(err)
        swal('Erro!', err, 'error')
      }
    },
    reset: function () {
      this.form = {
        name: '',
        description: '',
        state: 1
      }
    },
    initMQTT: function () {
      const self = this
      console.log('[MQTT] Iniciando conexão...')
      this.client = mqtt.connect(process.env.CDA_MQTT_HOST)

      this.client.on('connect', function () {
        self.client.subscribe('/help/chip_id')
      })

      this.client.on('message', function (topic, message) {
        const msg = message.toString()
        console.log('[MQTT] ' + msg)
        self.form.chip_id = msg
      })
    },
    closeMQTT: function () {
      console.log('[MQTT] Fechando conexão')
      this.client.end()
    }
  }
}
</script>

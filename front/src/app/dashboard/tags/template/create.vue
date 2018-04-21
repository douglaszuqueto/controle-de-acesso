<template>
  <div class="">
    <h3 class="">Create tag</h3>
    <hr>

    <form class="" method="post">
      <div class="form-group row">
        <label for="tag" class="col-sm-1 col-form-label">Tag</label>
        <div class="col-sm-5">
          <input type="text" class="form-control" v-model="form.tag" id="tag" placeholder="Tag">
        </div>
        <label for="user" class="col-sm-1 col-form-label">User:</label>
        <div class="col-sm-5">
          <select name="id_user" v-model="form.id_user" id="user" class="form-control">
            <option value="" selected></option>
            <option v-for="(user) in users" :key="user.ID" :value="user.ID">{{user.Name}}</option>
          </select>
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
  mounted: function () {
    this.getUsers()
    this.form.id_user = this.$route.params.id_user
  },
  data: () => {
    return {
      form: {
        tag: '',
        description: 'asd',
        state: 1
      },
      users: []
    }
  },
  methods: {
    save: async function () {
      try {
        const {data} = await http.post('/tag', this.form)
        const {Error} = data
        if (Error.Error) {
          throw Error.Message
        }
        this.reset()
        swal('Sucesso!', 'Cadastrado com sucesso', 'success')
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
        tag: '',
        id_user: 'e5be054d-37fe-499d-9e70-78dea410bd17',
        description: '',
        state: 1
      }
    },
    getUsers: async function () {
      let self = this
      http.get(`/user`)
        .then((data) => {
          self.users = data.data.Data
        })
    },
    initMQTT: function () {
      const self = this
      console.log('[MQTT] Iniciando conexão...')
      this.client = mqtt.connect(process.env.CDA_MQTT_HOST)

      this.client.on('connect', function () {
        self.client.subscribe('/help/tag')
      })

      this.client.on('message', function (topic, message) {
        const msg = message.toString()
        console.log('[MQTT] ' + msg)
        self.form.tag = msg
      })
    },
    closeMQTT: function () {
      console.log('[MQTT] Fechando conexão')
      this.client.end()
    }
  }
}
</script>

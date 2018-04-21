<template>
  <div class="">
    <h3 class="">Atualização de device</h3>
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
          <i class="fa fa-save"></i> Salvar
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { http } from '@/plugins/http'
import swal from 'sweetalert2'

export default {
  data: () => {
    return {
      form: {}
    }
  },
  mounted () {
    this.getDevice()
  },
  methods: {
    getDevice: function () {
      let self = this
      const {id} = this.$route.params
      http.get(`/device/${id}`)
        .then((data) => data.data)
        .then((data) => data.Data)
        .then((data) => {
          self.form = {
            name: data.Name,
            description: data.Description,
            state: data.State,
            chip_id: data.ChipID.String
          }
        })
        .catch((err) => console.log(err))
    },
    save: async function () {
      const {id} = this.$route.params
      const payload = this.form
      payload.state = (payload.state === true || payload.state === 1) ? 1 : 0
      http.put(`/device/${id}`, payload)
        .then((data) => data.data)
        .then((data) => {
          const {Error} = data
          if (Error.Error) {
            throw Error.Message
          }
          swal('Sucesso!', 'Device atualizado com sucesso', 'success')
          this.getDevice()
        })
        .catch((err) => {
          if (err instanceof Object) {
            swal('Erro!', err.message, 'error')
            return
          }
          console.log(err)
          swal('Erro!', err, 'error')
        })
    }
  }
}
</script>

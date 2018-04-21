<template>
  <div class="">
    <h3 class="">Update user</h3>
    <hr>

    <form class="" method="post">
      <div class="form-group row">
        <label for="name" class="col-sm-1 col-form-label">Name</label>
        <div class="col-sm-5">
          <input type="text" class="form-control" v-model="form.name" id="name" placeholder="Name">
        </div>
        <label for="email" class="col-sm-1 col-form-label">Email</label>
        <div class="col-sm-5">
          <input type="email" class="form-control" v-model="form.email" id="email" placeholder="Email">
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

export default {
  data: () => {
    return {
      form: {}
    }
  },
  mounted () {
    this.get()
  },
  methods: {
    get: function () {
      let self = this
      const {id} = this.$route.params
      http.get(`/user/${id}`)
        .then((data) => data.data)
        .then((data) => data.Data)
        .then((data) => {
          self.form = {
            name: data.Name,
            email: data.Email
          }
        })
        .catch((err) => console.log(err))
    },
    save: async function () {
      const {id} = this.$route.params
      http.put(`/user/${id}`, this.form)
        .then((data) => data.data)
        .then((data) => {
          const {Error} = data
          if (Error.Error) {
            throw Error.Message
          }
          swal('Sucesso!', 'Atualizado com sucesso', 'success')
          this.get()
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

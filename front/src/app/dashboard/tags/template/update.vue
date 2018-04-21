<template>
  <div class="">
    <h3 class="">Update tag</h3>
    <hr>

    <form class="" method="post">
      <div class="form-group row">
        <label for="tag" class="col-sm-1 col-form-label">Tag</label>
        <div class="col-sm-5">
          <input type="text" class="form-control" v-model="form.tag" id="tag" placeholder="Tag" disabled>
        </div>
        <label for="user" class="col-sm-1 col-form-label">User</label>
        <div class="col-sm-5">
          <select name="id_user" v-model="form.id_user" id="user" class="form-control">
            <option value="" selected></option>
            <option v-for="(user) in users" :key="user.ID" v-bind:value="user.ID">{{user.Name}}</option>
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

export default {
  data: () => {
    return {
      form: {},
      users: []
    }
  },
  mounted () {
    this.get()
    this.getUsers()
  },
  methods: {
    get: function () {
      let self = this
      const {id} = this.$route.params
      http.get(`/tag/${id}`)
        .then((data) => data.data)
        .then((data) => data.Data)
        .then((data) => {
          self.form = {
            tag: data.Tag,
            id_user: data.IDUser,
            description: data.Description,
            state: data.State
          }
        })
        .catch((err) => console.log(err))
    },
    save: async function () {
      const {id} = this.$route.params
      const payload = this.form
      payload.state = (payload.state === true || payload.state === 1) ? 1 : 0
      http.put(`/tag/${id}`, payload)
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
    },
    getUsers: async function () {
      let self = this
      http.get(`/user`)
        .then((data) => {
          self.users = data.data.Data
        })
    }
  }
}
</script>

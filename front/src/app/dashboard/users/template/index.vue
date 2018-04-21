<template>
  <div class="">
    <div class="row">
      <div class="col-6">
        <h3>Users</h3>
      </div>
      <div class="col-6 text-right">
        <router-link class="btn btn-success" :to="{name:'user.new'}">
          <i class="fa fa-plus"></i> <span class="d-none d-sm-inline">New user</span>
        </router-link>
        <button class="btn btn-info" @click="getAll()" data-toggle="tooltip" title="Update list">
          <i class="fa fa-sync-alt"></i>
        </button>
      </div>
    </div>
    <br>
    <div class="row">
      <div class="col-12">
        <table class="table table-responsive">
          <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Created</th>
            <th class="text-center" width="5%">#</th>
            <th class="text-center" width="5%">#</th>
            <th class="text-center" width="5%">#</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(user) in users" :key="user.ID">
            <td>{{user.Name}}</td>
            <td>{{user.Email}}</td>
            <td><nobr>{{user.CreatedAt}}</nobr></td>
            <td class="text-center">
              <router-link :to="{name: 'user.tags', params: {id: user.ID}}" class="btn btn-sm btn-link text-warning" data-toggle="tooltip" title="Tags">
                <i class="fa fa-tags"></i>
              </router-link>
            </td>
            <td class="text-center">
              <router-link :to="{name: 'user.update', params: {id: user.ID}}" class="btn btn-sm btn-link text-info" data-toggle="tooltip" title="Edit">
                <i class="fa fa-edit"></i>
              </router-link>
            </td>
            <td class="text-center">
              <button class="btn btn-sm btn-link text-danger" @click="remove(user.ID)" data-toggle="tooltip" title="Remove">
                <i class="fa fa-trash-alt"></i>
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>

  </div>
</template>

<script>
import {http} from '@/plugins/http'
import swal from 'sweetalert2'

export default {
  mounted: function () {
    this.getAll()
  },
  data: () => {
    return {
      users: {}
    }
  },
  methods: {
    getAll: async function () {
      let self = this
      http.get('/user')
        .then((data) => {
          self.users = data.data.Data
        })
    },
    remove: async function (id) {
      await http.delete(`/user/${id}`)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            return swal('Erro!', data.Error.Message, 'error')
          }
          swal('Sucesso!', data.Error.Message, 'success')
        })
        .catch((err) => swal('Opps...', err, 'error'))

      this.getAll()
    }
  }
}
</script>

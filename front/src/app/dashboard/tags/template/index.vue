<template>
  <div class="">
    <div class="row">
      <div class="col-6">
        <h3>Tags</h3>
      </div>
      <div class="col-6 text-right">
        <router-link class="btn btn-success" :to="{name:'tag.new'}">
          <i class="fa fa-plus"></i> <span class="d-none d-sm-inline">New tag</span>
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
            <th>Tag</th>
            <th>User</th>
            <th>Created_at</th>
            <th class="text-center">State</th>
            <th class="text-center" width="5%">#</th>
            <th class="text-center" width="5%">#</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(tag) in tags" :key="tag.ID">
            <td>{{tag.Tag}}</td>
            <td>{{tag.User.Name}}</td>
            <td><nobr>{{tag.CreatedAt}}</nobr></td>
            <td class="text-center">
              <div v-if="tag.State === 1">
                <span class="badge badge-success">Active</span>
              </div>
              <div v-else-if="tag.State === 0">
                <span class="badge badge-danger">Disabled</span>
              </div>
            </td>
            <td class="text-center">
              <router-link :to="{name: 'tag.update', params: {id: tag.ID}}" class="btn btn-sm btn-link text-info" data-toggle="tooltip" title="Edit">
                <i class="fa fa-edit"></i>
              </router-link>
            </td>
            <td class="text-center">
              <button class="btn btn-sm btn-link text-danger" @click="remove(tag.ID)" data-toggle="tooltip" title="Remove">
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
      tags: {}
    }
  },
  methods: {
    getAll: async function () {
      let self = this
      http.get('/tag')
        .then((data) => {
          self.tags = data.data.Data
        })
    },
    remove: async function (id) {
      await http.delete(`/tag/${id}`)
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

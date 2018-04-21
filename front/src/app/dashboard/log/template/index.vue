<template>
  <div class="">
    <div class="row">
      <div class="col-6">
        <h3>Log</h3>
      </div>
      <div class="col-6 text-right">
        <button class="btn btn-info" @click="getAll()" data-toggle="tooltip" title="Update list">
          <i class="fa fa-sync-alt"></i>
        </button>
      </div>
    </div>
    <br>
    <div class="row">
      <div class="col-12">
        <table class="table table-responsive-sm">
          <thead>
          <tr>
            <th>Device</th>
            <th>User</th>
            <th>Tag</th>
            <th>Created_at</th>
            <th>State</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(log) in logs" :key="log.ID">
            <td>{{log.Device}}</td>
            <td>{{log.User}}</td>
            <td>{{log.Tag}}</td>
            <td><nobr>{{log.CreatedAt}}</nobr></td>
            <td>
              <div v-if="log.State === 1">
                <span class="badge badge-success">Acesso autorizado</span>
              </div>
              <div v-else-if="log.State === 0">
                <span class="badge badge-danger">Acesso negado</span>
              </div>
              <div v-else-if="log.State === 2">
                <span class="badge badge-warning">Acesso não autorizado</span>
              </div>
              <div v-else-if="log.State === 3">
                <span class="badge badge-dark">Acesso inativo</span>
              </div>
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

export default {
  mounted: function () {
    this.getAll()
  },
  data: () => {
    return {
      logs: {}
    }
  },
  methods: {
    getAll: async function () {
      let self = this
      http.get('/log')
        .then((data) => {
          self.logs = data.data.Data
        })
    }
  }
}
</script>

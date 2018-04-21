<template>
  <div class="">
    <div class="row">
      <div class="col-6">
        <h3>Devices</h3>
      </div>
      <div class="col-6 text-right">
        <router-link class="btn btn-success" :to="{name:'device.new'}">
          <i class="fa fa-plus"></i> <span class="d-none d-sm-inline">New device</span>
        </router-link>
        <button class="btn btn-info" @click="getDevices()" data-toggle="tooltip" title="Update list">
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
            <th>ChipID</th>
            <th>Created</th>
            <th class="text-center">State</th>
            <th class="text-center" width="5%">#</th>
            <th class="text-center" width="5%">#</th>
            <th class="text-center" width="5%">#</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(device) in devices" :key="device.ID">
            <td>{{device.Name}}</td>
            <td>{{device.ChipID.String}}</td>
            <td><nobr>{{device.CreatedAt}}</nobr></td>
            <td class="text-center">
              <div v-if="device.State === 1">
                <span class="badge badge-success">Active</span>
              </div>
              <div v-else-if="device.State === 0">
                <span class="badge badge-danger">Disabled</span>
              </div>
            </td>
            <td class="text-center">
              <router-link :to="{name: 'device.tags', params: {id: device.ID}}" class="btn btn-sm btn-link text-warning" data-toggle="tooltip" title="Tags">
                <i class="fa fa-tags"></i>
              </router-link>
            </td>
            <td class="text-center">
              <router-link :to="{name: 'device.update', params: {id: device.ID}}" class="btn btn-sm btn-link text-info" data-toggle="tooltip" title="Edit">
                <i class="fa fa-edit"></i>
              </router-link>
            </td>
            <td class="text-center">
              <button class="btn btn-sm btn-link text-danger" @click="remove(device.ID)" data-toggle="tooltip" title="Remove">
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
    this.getDevices()
  },
  data: () => {
    return {
      devices: {}
    }
  },
  methods: {
    getDevices: async function () {
      let self = this
      http.get('/device')
        .then((data) => {
          self.devices = data.data.Data
        })
    },
    remove: async function (id) {
      await http.delete(`/device/${id}`)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            return swal('Erro!', data.Error.Message, 'error')
          }
          swal('Sucesso!', data.Error.Message, 'success')
        })
        .catch((err) => swal('Opps...', err, 'error'))

      this.getDevices()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
</style>

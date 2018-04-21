<template>
  <div class="">
    <div class="row">
      <div class="col-6">
        <h3>{{device.Name}} - Tags</h3>
      </div>
      <div class="col-6">
        <div class="text-right">
          <button class="btn btn-success text-white" data-toggle="modal" data-target="#modal-attach">
            <i class="fa fa-plus"></i> <span class="d-none d-sm-inline">Attach new tag(user)</span>
          </button>
          <button class="btn btn-info" @click="getAll()" data-toggle="tooltip" title="Update list">
            <i class="fa fa-sync-alt"></i>
          </button>
        </div>
      </div>
    </div>
    <br>
    <table class="table table-responsive">
      <thead>
      <tr>
        <th>Tag</th>
        <th>Created</th>
        <th class="text-center">State</th>
        <th class="text-center" width="5%">#</th>
        <th class="text-center" width="5%">#</th>
        <th class="text-center" width="5%">#</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(tag) in tags" :key="tag.ID">
        <td>{{tag.Tag}}</td>
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
          <div v-if="tag.State === 1">
            <button class="btn btn-sm btn-link text-success" @click="changeState(tag.ID)" data-toggle="tooltip"
                    title="Deactivate access">
              <i class="fa fa-check-square"></i>
            </button>
          </div>
          <div v-else-if="tag.State === 0">
            <button class="btn btn-sm btn-link text-dark" @click="changeState(tag.ID)" data-toggle="tooltip"
                    title="Activate access">
              <i class="fa fa-check-square"></i>
            </button>
          </div>
        </td>
        <td class="text-center">
          <router-link :to="{name: 'tag.update', params: {id: tag.ID}}" class="btn btn-sm btn-link text-warning"
                       data-toggle="tooltip" title="Edit tag">
            <i class="fa fa-tag"></i>
          </router-link>
        </td>
        <td class="text-center">
          <button class="btn btn-sm btn-link text-danger" @click="dettach(tag.ID)" data-toggle="tooltip"
                  title="Dettach tag">
            <i class="fa fa-trash-alt"></i>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
    <button type="button" class="btn btn-primary float-right" onclick="history.go(-1)" title="Return">
      <i class="fas fa-arrow-circle-left"></i> Return
    </button>
    <form action="">
      <div class="modal fade" id="modal-attach">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="exampleModalLabel">Attach a tag</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <div class="form-group row">
                <label for="user" class="col-sm-2 col-form-label">Usuário:</label>
                <div class="col-sm-10">
                  <select name="id_user" v-model="form.id_user" id="user" class="form-control"
                          @change="onUserSelected()">
                    <option value=""></option>
                    <option v-for="(user) in users" :key="user.ID" :value="user.ID">{{user.Name}}</option>
                  </select>
                </div>

              </div>
              <div class="form-group row">
                <label for="tag" class="col-sm-2 col-form-label">Tag:</label>
                <div class="col-sm-10">
                  <select name="id_tag" v-model="form.id_tag" id="tag" class="form-control" :disabled="!form.id_user">
                    <option value=""></option>
                    <option v-for="(tag) in user_tags" :key="tag.id" :value="tag.ID">{{tag.Tag}}</option>
                  </select>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              <button type="button" class="btn btn-success" @click.prevent="attach()">Attach</button>
            </div>
          </div>
        </div>
      </div>
    </form>

  </div>
</template>

<script>
import { http } from '@/plugins/http'
import $ from 'jquery'
import swal from 'sweetalert2'

export default {
  mounted: function () {
    let self = this
    this.getAll()
    this.getUsers()
    $('#modal-attach').on('hidden.bs.modal', function () {
      self.getAll()
      self.user_tags = []
      self.form.id_user = null
    })
  },
  data: () => {
    return {
      device: {
        Name: ''
      },
      tags: {},
      form: {},
      users: [],
      user_tags: []
    }
  },
  methods: {
    getAll: async function () {
      let self = this
      http.get(`/device/${this.$route.params.id}/tags`)
        .then((data) => {
          const {Device, Tags} = data.data.Data
          self.device = Device
          self.tags = Tags
        })
    },
    getUsers: async function () {
      let self = this
      http.get(`/user`)
        .then((data) => {
          self.users = data.data.Data
        })
    },
    onUserSelected: function () {
      let self = this
      this.form.id_tag = ''
      http.get(`/user/${this.form.id_user}/tags`)
        .then((data) => {
          self.user_tags = data.data.Data
        })
    },
    attach: function () {
      const payload = {
        id_device: this.$route.params.id,
        id_tag: this.form.id_tag
      }
      http.post('/tag/attach', payload)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            throw data.Error.Message
          }
          swal('Sucesso!', data.Error.Message, 'success')
        })
        .catch((err) => {
          swal('Erro!', err, 'error')
          console.log(err)
        })
    },
    dettach: function (idTag) {
      console.log(this.form)
      const payload = {
        id_device: this.$route.params.id,
        id_tag: idTag
      }
      http.post('/tag/dettach', payload)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            throw data.Error.Message
          }
          swal('Sucesso!', data.Error.Message, 'success')
          this.getAll()
        })
        .catch((err) => {
          swal('Erro!', err, 'error')
          console.log(err)
        })
    },
    changeState: function (idTag) {
      const payload = {
        id_device: this.$route.params.id,
        id_tag: idTag
      }
      http.put(`/device/${payload.id_device}/tags/${payload.id_tag}`)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            throw data.Error.Message
          }
          swal('Sucesso!', data.Error.Message, 'success')
          this.getAll()
        })
        .catch((err) => {
          swal('Erro!', err, 'error')
          console.log(err)
        })
    }
  }
}
</script>

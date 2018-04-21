<template>
  <div class="row card-deck">
    <div class="col-12 col-md-4 col-lg-4">
      <div class="card">
        <div class="card-header">
          <div class="row">
            <div class="col-6">Devices</div>
            <div class="col-6">
              <router-link :to="{name: 'devices.index'}" class="float-right">
                <i class="fas fa-arrow-right"></i>
              </router-link>
            </div>
          </div>
        </div>
        <div class="card-body">
          <h3 class="card-text text-center">
            <i class="fas fa-server"></i>
          </h3>
          <h3 class="card-text text-center">{{statistics.devices}}</h3>
        </div>
      </div>
    </div>

    <div class="col-12 col-md-4 col-lg-4">
      <div class="card">
        <div class="card-header">
          <div class="row">
            <div class="col-6">Users</div>
            <div class="col-6">
              <router-link :to="{name: 'user.index'}" class="float-right">
                <i class="fas fa-arrow-right"></i>
              </router-link>
            </div>
          </div>
        </div>
        <div class="card-body">
          <h3 class="card-text text-center">
            <i class="fas fa-user"></i>
          </h3>
          <h3 class="card-text text-center">{{statistics.users}}</h3>
        </div>
      </div>
    </div>

    <div class="col-12 col-md-4 col-lg-4">
      <div class="card">
        <div class="card-header">
          <div class="row">
            <div class="col-6">Tags</div>
            <div class="col-6">
              <router-link :to="{name: 'tag.index'}" class="float-right">
                <i class="fas fa-arrow-right"></i>
              </router-link>
            </div>
          </div>
        </div>
        <div class="card-body">
          <h3 class="card-text text-center">
            <i class="fas fa-tags"></i>
          </h3>
          <h3 class="card-text text-center">{{statistics.tags}}</h3>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { http } from '@/plugins/http'

export default {
  mounted: function () {
    this.getDevices()
    this.getTags()
    this.getUsers()
  },
  data: () => {
    return {
      statistics: {
        users: 0,
        tags: 0,
        devices: 0
      }
    }
  },
  methods: {
    getDevices: async function () {
      await http.get('/device')
        .then((data) => { this.statistics.devices = data.data.Data.length })
    },
    getTags: async function () {
      await http.get('/tag')
        .then((data) => { this.statistics.tags = data.data.Data.length })
    },
    getUsers: async function () {
      await http.get('/user')
        .then((data) => { this.statistics.users = data.data.Data.length })
    }
  }
}
</script>

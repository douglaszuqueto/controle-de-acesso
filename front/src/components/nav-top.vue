<template>
  <div class="">
    <nav class="navbar fixed-top navbar-light">
      <button class="navbar-toggler text-white d-md-none" type="button" data-toggle="collapse" data-target="#navbarSupportedContent">
        <i class="fas fa-bars"></i>
      </button>

      <router-link :to="{name: 'dashboard'}" class="navbar-brand text-light d-none d-sm-inline" >
        <i class="fa fa-unlock-alt fa-lg" style="margin-left: 5px"></i>
        <span class="" style="margin-left: 20px">Controle de acesso</span>
      </router-link>

      <div class="float-right text-right">
        <div class="d-none d-lg-inline ">
          <img :src="user.gravatar" alt="" style="border-radius: 50px" width="40px">
          <span class="text-white text-right" style="margin-top: 8px;margin-right: 10px">Seja bem vindo(a) {{user.name}}</span>
        </div>
        <a href="#" class="text-white float-right" style="margin-top: 2%;margin-left: 10px" title="Sair" @click.prevent="logout()">
          <i class="fas fa-lg fa-sign-out-alt"></i>
        </a>
      </div>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav">
          <li class="nav-item ">
            <router-link :to="{name: 'dashboard'}" class="nav-link text-white">
              <i class="fas fa-home fa-lg"></i> Home
            </router-link>
          </li>
          <li class="nav-item ">
            <router-link :to="{name: 'panel'}" class="nav-link text-white">
              <i class="fas fa-tachometer-alt fa-lg"></i> Dashboard
            </router-link>
          </li>
          <li class="nav-item ">
            <router-link :to="{name: 'devices.index'}" class="nav-link text-white">
              <i class="fas fa-server fa-lg"></i> Devices
            </router-link>
          </li>
          <li class="nav-item ">
            <router-link :to="{name: 'tag.index'}" class="nav-link text-white">
              <i class="fas fa-tags fa-lg"></i> Tags
            </router-link>
          </li>
          <li class="nav-item ">
            <router-link :to="{name: 'user.index'}" class="nav-link text-white">
              <i class="fas fa-user fa-lg"></i> Users
            </router-link>
          </li>
          <li class="nav-item ">
            <router-link :to="{name: 'log.index'}" class="nav-link text-white">
              <i class="fas fa-chart-line fa-lg"></i> History
            </router-link>
          </li>
          <!--<li class="nav-item ">-->
            <!--<a class="nav-link text-white" href="#">-->
              <!--<i class="fas fa-cog fa-lg"></i> Config-->
            <!--</a>-->
          <!--</li>-->
        </ul>
      </div>
    </nav>
  </div>
</template>

<script>
import $ from 'jquery'

$(function () {
  const navMobile = $('#navbarSupportedContent')
  $('.nav-link').on('click', function (e) {
    e.preventDefault()
    navMobile.removeClass('show').addClass('collapsing')
  })
})
export default {
  mounted: function () {
    this.loadUser()
  },
  data: () => {
    return {
      user: {
        name: '',
        email: '',
        gravatar: 'https://secure.gravatar.com/avatar'
      }
    }
  },
  methods: {
    loadUser: function () {
      const u = JSON.parse(localStorage.getItem('user'))
      if (u === null) {
        return
      }
      this.user.name = u.Name ? u.Name : ''
      this.user.email = u.Email ? u.Email : ''
      this.user.gravatar = u.Gravatar ? u.Gravatar : 'https://secure.gravatar.com/avatar'
    },
    logout: function () {
      localStorage.removeItem('user')
      this.$router.push('/login')
    }
  }
}
</script>

<style>
  .fixed-top {
    background-color: #8973cb;
  }

  .nav-item:hover{
    background-color: #765ecb;
  }
</style>

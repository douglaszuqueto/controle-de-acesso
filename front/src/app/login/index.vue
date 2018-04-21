<template>
  <div class="container login-page">
    <div class="row">
      <div class="align-self-center col-12 col-md-6 offset-md-3 col-lg-4 offset-lg-4">
        <h1 class="text-center" style="color: #444051">
          <i class="fas fa-lg fa-user-circle"></i>
        </h1>
        <h2 class="text-center" style="color: #444051">Login</h2>
        <br>
        <form method="post">
          <div class="form-row">
            <div class="form-group col-12">
              <label for="email">Email</label>
              <input type="email" class="form-control" v-model="user.email" id="email" placeholder="Email">
            </div>
          </div>
          <div class="form-row">
            <div class="form-group col-12">
              <label for="password">Password</label>
              <input type="password" class="form-control" v-model="user.password" id="password" placeholder="Password">
            </div>
          </div>
          <button type="submit" class="btn btn-success col-12" @click.prevent="login()">
            Sign in<br/>
            <i class="fas fa-lg fa-sign-in-alt"></i>
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { http } from '@/plugins/http'
import swal from 'sweetalert2'

export default {
  data: () => {
    return {
      user: {
        email: 'douglas.zuqueto@gmail.com',
        password: 'admin'
      }
    }
  },
  methods: {
    login: async function () {
      http.post('/auth', this.user)
        .then((data) => data.data)
        .then((data) => {
          if (data.Error.Error) {
            throw new Error(data.Error.Message)
          }
          localStorage.setItem('user', JSON.stringify(data.Data))
          window.location = '/'
        })
        .catch((err) => {
          console.log(err)
          swal('Erro!', err.message, 'error')
        })
    }
  }
}
</script>

<style scoped>
  .login-page {
    margin-top: 50px;
  }

  @media (min-width: 720px) {
    .login-page{
      margin-top: 125px;
    }
  }
</style>

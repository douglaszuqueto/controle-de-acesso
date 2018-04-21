<template>
  <div class="">
    <h3 class="">Create user</h3>
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

      <div class="form-group text-right row">
        <div class="col">
          <button type="button" class="btn btn-primary" onclick="history.go(-1)"
                  title="Return">
            <i class="fas fa-arrow-circle-left"></i> Return
          </button>
          <button type="submit" class="btn btn-success" @click.prevent="save()">
            <i class="fa fa-save"></i> Save
          </button>
        </div>
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
      form: {
        name: 'Douglas',
        email: 'douglas.zuqueto@gmail.com'
      }
    }
  },
  methods: {
    save: async function () {
      try {
        const {data} = await http.post('/user', this.form)
        const {Error} = data
        if (Error.Error) {
          throw Error.Message
        }
        this.reset()
        swal('Sucesso!', 'Cadastrado com sucesso', 'success')
      } catch (err) {
        if (err instanceof Object) {
          swal('Erro!', err.message, 'error')
          return
        }
        console.log(err)
        swal('Erro!', err, 'error')
      }
    },
    reset: function () {
      this.form = {
        name: '',
        email: ''
      }
    }
  }
}
</script>

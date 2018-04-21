<template>
  <div id="app">
    <nav-left/>
    <nav-top/>

    <div class="container">
      <transition name="fade" mode="out-in" appear>
        <router-view class="" style="margin-top: 80px;"/>
      </transition>
    </div>
  </div>
</template>

<script>
import NavTop from '../../components/nav-top'
import NavLeft from '../../components/nav-left'
//
export default {
  components: {
    NavTop, NavLeft
  }
}
</script>

<style>
  .container-fluid-app{
    padding-left: 90px!important;
    padding-right: 76px!important;
  }
  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.1s
  }

  .fade-enter, .fade-leave-active {
    opacity: 0
  }
</style>

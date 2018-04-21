<template>
  <div>
    <div class="nav-left d-none d-xl-block">
      <ul class="navbar-nav">
        <li class="nav-item active" data-toggle="tooltip" data-placement="right" title="Home">
          <router-link :to="{name: 'dashboard'}" class="nav-link text-white" exact>
            <i class="fas fa-home fa-lg"></i>
          </router-link>
        </li>
        <li class="nav-item">
          <router-link :to="{name: 'panel'}" class="nav-link text-white" data-toggle="tooltip" data-placement="right"
                       title="Dashboard">
            <i class="fas fa-tachometer-alt fa-lg"></i>
          </router-link>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Devices">
          <router-link :to="{name: 'devices.index'}" class="nav-link text-white">
            <i class="fas fa-server fa-lg"></i>
          </router-link>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Tags">
          <router-link :to="{name: 'tag.index'}" class="nav-link text-white">
            <i class="fas fa-tags fa-lg"></i>
          </router-link>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Usuários">
          <router-link :to="{name: 'user.index'}" class="nav-link text-white">
            <i class="fas fa-user fa-lg"></i>
          </router-link>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="History">
          <router-link :to="{name: 'log.index'}" class="nav-link text-white">
            <i class="fas fa-chart-line fa-lg"></i>
          </router-link>
        </li>
        <!--<li class="nav-item" data-toggle="tooltip" data-placement="right" title="Configurações">-->
          <!--<a class="nav-link text-white" href="#">-->
            <!--<i class="fas fa-cog fa-lg"></i>-->
          <!--</a>-->
        <!--</li>-->
      </ul>
    </div>
  </div>
</template>

<style>
  .nav-left {
    width: 64px;
    position: fixed;
    top: 56px;
    bottom: 0;
    left: 0;
    z-index: 9999;
    padding: 0;
    border-right: none;
    min-height: calc(100vh - 50px);
    background-color: #414150;
    overflow: visible;
  }

  .nav-item {
    padding: 15px 0 0 0;
    text-align: center;
    height: 64px;
  }

  .nav-item.active {
    background-color: #322D37;
  }

  .nav-item. > li > a.active {
    background-color: #322D37;
  }

  .nav-item:hover {
    cursor: pointer;
    background-color: #322D37;
  }
</style>

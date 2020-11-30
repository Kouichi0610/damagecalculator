<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">対象選択</router-link> |
      <router-link to="/attacker0">攻撃調整A</router-link> |
      <router-link to="/attacker1">攻撃調整B</router-link> |
      <router-link to="/attacker2">攻撃調整C</router-link> |
      <router-link to="/attacker3">攻撃調整D</router-link> |
      <router-link to="/defender">耐久調整</router-link> |
      <router-link to="/speed">速度調整</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class App extends Vue {
}
</script>

<style scoped>
#app {
  width: 1000px;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>

<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/current">ポケモン選択(HOME)</router-link> |
      <router-link to="/attacker">攻撃調整</router-link> |
      <router-link to="/defender">耐久調整</router-link> |
      <router-link to="/speed">速度調整</router-link> |
      <router-link to="/sandboxts">SandBox(TypeScript)</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <names v-bind:tag="pokemon" v-bind:name="name" v-bind:candidates="candidates"></names>
    <router-view/>
  </div>
</template>

<script>
import Names from './components/names.vue'

export default {
  name: 'App',
  components: {
    'names': Names,
  },
  created: function() {
    /*
    // TODO:最初の通信でエラーが出てない場合ローカルモードに切り替える
    this.axios.get('get_names')
    .then((response) => {
      // TODO:絞り込み
      this.candidates = response.data;
      console.log('pokeLen:' + this.candidates.length);
    })
    .catch((e) => {
      alert('failed:' + e);
    })
    */
  },
  methods: {
  },
  data: function() {
    return {
      pokemon: 'ポケモン',
      name: 'ピカチュウ',
      candidates: [],
    }
  }


}
</script>

<style>
#app {
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

<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/select">ポケモン選択</router-link> |
      <router-link to="/selectold">ポケモン選択(旧)</router-link> |
      <router-link to="/attacker">攻撃調整</router-link> |
      <router-link to="/defender">耐久調整</router-link> |
      <router-link to="/speed">速度調整</router-link> |
      <router-link to="/sandboxts">SandBox</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
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

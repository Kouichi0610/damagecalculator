<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">調整対象選択</router-link> |
      <router-link to="/attacker0">{{ moveA }}</router-link> |
      <router-link to="/attacker1">{{ moveB }}</router-link> |
      <router-link to="/attacker2">{{ moveC }}</router-link> |
      <router-link to="/attacker3">{{ moveD }}</router-link> |
      <router-link to="/defender">耐久調整</router-link> |
      <router-link to="/speed">速度調整</router-link>
    </div>
    <router-view/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from "vuex-class";

const namespace: string = "attackerState";

@Component
export default class App extends Vue {
  @Getter('selectedMoves', { namespace })
  private selectedMoves!: string[];
  readonly tag: string = '攻撃調整:';

  get moveA(): string {
    return this.tag + this.selectedMoves[0];
  }
  get moveB(): string {
    return this.tag + this.selectedMoves[1];
  }
  get moveC(): string {
    return this.tag + this.selectedMoves[2];
  }
  get moveD(): string {
    return this.tag + this.selectedMoves[3];
  }

}
</script>

<style scoped>
#app {
  width: 1200px;
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

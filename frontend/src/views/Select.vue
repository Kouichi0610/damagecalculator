/*
  調整対象選択
*/
<template>
  <div class="select">
    <target :targetName="name" @changeTarget="onTarget" :show="false"></target>
    <target-select @selectPokemon="onSelect"></target-select>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TargetSelect from '../components/targetSelect/targetSelect.vue'
import Target from '../components/target/target.vue'

import router from '../router/index'

@Component({
  components: {
    TargetSelect,
    Target,
  }
})
export default class Select extends Vue {
  private name: string = ""

  onTarget(name: string) {
    router.push({path: '/attacker0'})
  }

  onSelect(name: string) {
    this.name = name;
  }
}
</script>

<style scoped>
</style>

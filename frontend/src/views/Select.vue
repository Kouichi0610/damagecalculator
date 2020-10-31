/*
  調整対象選択
*/
<template>
  <div class="select">
    <target :targetName="name" :show="false"></target>
    <target-select @selectPokemon="onSelect"></target-select>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TargetSelect from '../components/targetSelect/targetSelect.vue'
import Target from '../components/target/target.vue'
import router from "../router"

@Component({
  components: {
    TargetSelect,
    Target,
  }
})
export default class Select extends Vue {
  private name: string = ""

  onSelect(name: string) {
    this.name = name;
    router.push({path: 'attacker'});
  }
}
</script>

<style scoped>
</style>

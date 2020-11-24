/*
  調整対象選択
*/
<template>
  <div class="select">
    <species-loader :target="name" @change="onComplete"></species-loader>
    <target-select @selectPokemon="onSelect"></target-select>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import TargetSelect from '../components/targetSelect/targetSelect.vue'
import SpeciesLoader from '../components/species/speciesLoader.vue'

import { PokeData } from '../store/species/types'

import router from '../router/index'

@Component({
  components: {
    TargetSelect,
    SpeciesLoader,
  }
})
export default class Select extends Vue {
  private name: string = ""
  private target: PokeData = PokeData.default();

  // 対象ポケモン選択時
  onSelect(name: string) {
    this.name = name;
  }

  // 種族読み込み完了時
  onComplete(data: PokeData) {
    this.target = data;
  }

  @Watch('target', { deep: true })
  targetChanged() {
    if (!this.target.hasTarget()) return;
    router.push({path: '/attacker0'})
  }
}
</script>

<style scoped>
</style>

/*
  調整対象選択
  TODO:Currentと置き換え
*/
<template>
  <div class="select">
    <target-select @selectPokemon="onSelect"></target-select>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import TargetSelect from '../components/targetSelect/targetSelect.vue'

@Component({
  components: {
    TargetSelect,
  }
})
export default class Select extends Vue {

  created() {
  }

  /*
    ポケモン名
    ->種族情報
      Species
      Types
      Weight
      Ability
      Moves
  
  */
  onSelect(name: string) {
    console.log('selected:' + name);
  }
}
</script>

<style scoped>
</style>

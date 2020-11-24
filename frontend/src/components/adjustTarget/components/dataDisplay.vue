<template>
  <div class="data-display">
    <div class="row mb-1">
      <div class="col-3">{{ name }}</div>
      <div class="col-3">{{ types }}</div>
      <div class="col-3">{{ weight }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';
import { PokeData } from '../../../store/species/types'
/*
  名前、タイプ、重さ
*/
@Component
export default class DataDisplay extends Vue {
  @Prop()
  private data!: PokeData;

  get name(): string {
    return this.data.name;
  }
  get types(): string {
    let res = '';
    let len = this.data.types.length;
    for (var i = 0; i < len; i++) {
      res += this.data.types[i];
      if (i < len-1) {
        res += '/';
      }
    }
    return res;
  }
  get weight(): string {
    return '' + this.data.weight + 'kg';
  }
 
}
</script>

<style scoped>
</style>

<template>
  <div class="input-group mb-3">
    <p>種族値合計</p>
    <input type="number" v-model="total" class="form-control" aria-label="Text input with checkbox">
    <p>以上</p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';

// 種族値合計
@Component
export default class SpeciesTotal extends Vue {
  @Prop()
  private initialTotal: number;

  private total: number = 0;
  created() {
    this.total = this.initialTotal;
  }

  @Watch('total')
  private onTotalChanged(after: number, before: number) {
    if (after == before) {
      return;
    }
    this.$emit('changed', after);
  }
  
}
</script>
<style scoped>
</style>

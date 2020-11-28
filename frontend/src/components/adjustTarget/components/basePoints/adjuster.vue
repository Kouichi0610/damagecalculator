<template>
  <div class="adjuster">
    <span class="row mb-1">
        <div class="col-sm-7 pt-1">
            <input type="range" class="form-control-change" id="slider" v-model.number="current" step="4" min="0" max="252" :disabled="!enable">
        </div>
        <div class="col-1">{{ current }}</div>
    </span>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { IBasePoint } from '../../../../store/basePoints/types';

/*
  基礎ポイント調整
*/
@Component
export default class Adjuster extends Vue {
  @Prop()
  private basePoint!: IBasePoint;
  @Prop()
  private enable!: boolean;

  private current: number = 0;

  created() {
    this.current = this.basePoint.value;
  }

  @Watch('current')
  private currentChanged() {
    if (this.current > this.basePoint.limit()) {
      this.current = this.basePoint.limit();
    }
    this.$emit('change', this.basePoint.index, this.current);
  }

  @Watch('basePoint', { deep: true })
  private basePointChanged() {
    this.current = this.basePoint.value;
  }
}
</script>

<style scoped>
</style>

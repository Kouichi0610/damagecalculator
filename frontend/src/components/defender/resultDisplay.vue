<template>
  <div class="result-display">
    <div class="row mb-1">
      <template v-if="isDetermine1">
        <b-progress class="col-6" :max="100" show-value animated>
          <b-progress-bar :value="100" :label="alpha" variant="danger"></b-progress-bar>
        </b-progress>
      </template>
      <template v-else>
        <b-progress class="col-6" :max="100" show-value animated>
          <b-progress-bar :value="rateMin" :label="alpha" variant="danger"></b-progress-bar>
          <b-progress-bar :value="rateMax" variant="warning"></b-progress-bar>
        </b-progress>
      </template>
      <div class="col-4">
        {{ result.toString() }}
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { Result } from '../../store/defender/receiveDamages'

@Component
export default class ResultDisplay extends Vue {
  @Prop()
  result!: Result;

  get rateMin(): number {
    return this.result.rateMin;
  }
  get rateMax(): number {
    if (this.result.rateMax > 100) {
      return 100 - this.result.rateMin;
    }
    return this.result.rateMax - this.result.rateMin;
  }

  get isDetermine1(): boolean {
    return this.result.determineCount == 1;
  }
  
}
</script>
<style scoped>
</style>

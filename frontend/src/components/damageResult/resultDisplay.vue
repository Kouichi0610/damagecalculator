<template>
  <div class="result-display">
    <div class="row">
      <template v-if="isDetermine1">
        <b-progress class="col-md-2" :max="100" show-value animated>
          <b-progress-bar :value="100" :label="alpha" variant="danger"></b-progress-bar>
        </b-progress>
      </template>
      <template v-else>
        <b-progress class="col-md-2" :max="100" show-value animated>
          <b-progress-bar :value="rateMin" :label="alpha" variant="danger"></b-progress-bar>
          <b-progress-bar :value="rateMax" variant="warning"></b-progress-bar>
        </b-progress>
      </template>
      <div class="col-md-8">
        {{ result.toString() }}
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { IDamageResult } from '../../store/damageResult/damageResult'

@Component
export default class ResultDisplay extends Vue {
  @Prop()
  result!: IDamageResult;

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

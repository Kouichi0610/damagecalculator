<template>
  <div class="damageRate">
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
import { Vue, Prop, Component, Watch } from "vue-property-decorator";
import { DamageRateInfo } from '../target/store/damageRateInfo'

@Component
export default class DamageRate extends Vue {
  @Prop() private result!: DamageRateInfo;
  
  get rateMin(): number {
    return this.result.RateMin;
  }
  get rateMax(): number {
    if (this.result.RateMax > 100) {
      return 100 - this.result.RateMin;
    }
    return this.result.RateMax - this.result.RateMin;
  }

  get isDetermine1(): boolean {
    return this.result.DetermineCount == 1;
  }
}
</script>

<style scoped>
</style>

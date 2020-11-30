// 耐久調整
<template>
  <div class="defender">
    <adjust-target :speedLock="false" @targetCondition="changeTargetCondition"></adjust-target>
    <list-display :results="results"></list-display>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import { Getter } from "vuex-class";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'
import ListDisplay from '../components/damageResult/listDisplay.vue'

import { ReceiveDamages, Result } from '../store/defender/receiveDamages'
import { TargetCondition } from '../store/target/targetCondition'

const namespace: string = "defenderState";

@Component({
  components: {
    AdjustTarget,
    ListDisplay,
  }
})
export default class Defender extends Vue {
  private targetCondition: TargetCondition = TargetCondition.default();
  private results: Result[] = [];

  changeTargetCondition(targetCondition: TargetCondition) {
    this.targetCondition = targetCondition;
  }

  @Watch('targetCondition', {deep: true})
  calcDamages() {
    if (!this.targetCondition.enable()) return;
    new ReceiveDamages().receiveDamages(this.targetCondition)
    .then((results) => {
      this.results = results;
    })
  }

}
</script>

<style scoped>

</style>

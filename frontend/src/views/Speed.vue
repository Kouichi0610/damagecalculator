<template>
  <div class="speed-adjuster">
    <adjust-target :speedLock="true" @targetCondition="changeTargetCondition"></adjust-target>
    <b-form-checkbox id="check-trickroom" v-model="trickroom">{{ trickRoomMessage }}</b-form-checkbox>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'

import Target from "../components/target/target.vue";
import SpeedOrder from "../components/speedOrder/speedOrder.vue";
import { TargetCondition } from "../store/target/targetCondition";

// TODO:targetAdjuster 速度以外動かさない
// TODO:targetConditionをAPIに渡す
// TODO:補正後の能力値を表示

@Component({
  components: {
    AdjustTarget,
  },
})
export default class Speed extends Vue {
  private targetCondition: TargetCondition = TargetCondition.default();

  private trickroom: boolean = false;

  get trickRoomMessage(): string {
    if (this.trickroom) {
      return 'トリックルームON';
    }
      return 'トリックルームOFF';
  }

  @Watch('targetCondition', {deep: true})
  calcSpeed() {
    if (!this.targetCondition.enable()) return;
    console.log('calcSpeed:');
  }

changeTargetCondition(targetCondition: TargetCondition) {
    this.targetCondition = targetCondition;
  }
}
</script>
<style scoped>
</style>

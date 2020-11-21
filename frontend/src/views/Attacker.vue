// 攻撃側
<template>
  <div class="attacker">
    <adjust-target></adjust-target>
    <target :show="true" @target="targetChanged"></target>
    <move-result :index="index" :damages="defenderDamages"></move-result>
    <div class="row mb-1">
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import { Getter } from "vuex-class";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'

import Target from '../components/target/target.vue'
import MoveResult from '../components/attacker/moveResult.vue'
import { MoveInfo } from '../components/moves/store/types'
import { TargetState } from '../components/target/store/types'

import { DefenderDamages } from '../components/target/store/defenderDamages'

const namespace: string = "target";

@Component({
  components: {
    Target,
    MoveResult,
    AdjustTarget,
  }
})
export default class Attacker extends Vue {
  @Prop() private index!: number;
  @Getter("defenderDamages", { namespace })
  private defenderDamages!: DefenderDamages;

}
</script>

<style scoped>

</style>

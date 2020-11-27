// 攻撃側
<template>
  <div class="attacker">
    <adjust-target :speedLock="false" @targetCondition="changeTargetCondition"></adjust-target>
    <move-selector :physicals="physicals" :specials="specials" @select="selectMove"></move-selector>
    <p>わざ:{{currentMove}}</p>
    <list-display :results="results"></list-display>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { Action, Getter, Mutation } from "vuex-class";

import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'
import MoveSelector from '../components/moves/moveSelector.vue'
import ListDisplay from '../components/attacker/listDisplay.vue'

import { SendDamages, Result } from '../store/attacker/sendDamage'
import { TargetCondition } from '../store/target/targetCondition'
import { Move, MoveLoader } from '../store/attacker/types'

const namespace: string = "attackerState";

@Component({
  components: {
    AdjustTarget,
    MoveSelector,
    ListDisplay,
  }
})
export default class Attacker extends Vue {
  @Prop() private index!: number;

  @Action('loadMoves', { namespace })
  private loadMoves!: (attacker: string) => void;
  @Getter('physicals', { namespace })
  private physicals!: Move[];
  @Getter('specials', { namespace })
  private specials!: Move[];
  @Getter('currentMove', { namespace })
  private currentMove!: string;
  @Mutation('setCurrent', { namespace })
  private setCurrent!: (current: number) => void;
  @Mutation('setMove', { namespace })
  private setMove!: (move: string) => void;

  private targetCondition: TargetCondition = TargetCondition.default();
  private results: Result[] = [];

  @Watch('index')
  changeTab() {
    this.setCurrent(this.index);
  }

  @Watch('targetCondition', {deep: true})
  @Watch('currentMove', {deep: true})
  calcDamages() {
    if (!this.targetCondition.enable()) return;
    if (this.currentMove.length == 0) return;
    new SendDamages().sendDamages(this.targetCondition, this.currentMove)
    .then((results) => {
      this.results = results;
    });
  }

  @Watch('targetCondition.target')
  targetChanged(attacker: string, before: string) {
    this.loadMoves(attacker);
  }

  changeTargetCondition(targetCondition: TargetCondition) {
    this.targetCondition = targetCondition;
  }

  selectMove(move: Move) {
    this.setMove(move.name);
  }
}
</script>

<style scoped>

</style>

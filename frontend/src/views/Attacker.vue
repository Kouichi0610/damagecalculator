// 攻撃側
<template>
  <div class="attacker">
    <adjust-target @sendDamage="changeSendDamage"></adjust-target>
    <move-selector :physicals="physicals" :specials="specials" @select="selectMove"></move-selector>
    <p>わざ:{{currentMove}}</p>
    <list-display :results="results"></list-display>
    <p>resulsssst.</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { Action, Getter, Mutation } from "vuex-class";

import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'
import MoveSelector from '../components/moves/moveSelector.vue'
import ListDisplay from '../components/attacker/listDisplay.vue'

import { SendDamages, Result } from '../store/attacker/sendDamage'
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

  private sendDamages: SendDamages = SendDamages.default();
  private results: Result[] = [];

  @Watch('index')
  changeTab() {
    this.setCurrent(this.index);
  }

  @Watch('sendDamages', {deep: true})
  @Watch('currentMove', {deep: true})
  calcDamages() {
    if (!this.sendDamages.enable()) return;
    if (this.currentMove.length == 0) return;

    this.sendDamages.sendDamages(this.currentMove)
    .then((results) => {
      console.log('result:' + results.length);
      this.results = results;
    });
  }

  @Watch('sendDamages.attacker')
  targetChanged(attacker: string, before: string) {
    this.loadMoves(attacker);
  }

  changeSendDamage(sendDamages: SendDamages) {
    this.sendDamages = sendDamages;
  }

  selectMove(move: Move) {
    this.setMove(move.name);
  }
}
</script>

<style scoped>

</style>

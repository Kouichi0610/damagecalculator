// 攻撃側
<template>
  <div class="attacker">
    <adjust-target @sendDamage="changeSendDamage"></adjust-target>
    <move-selector :physicals="physicals" :specials="specials" @select="selectMove"></move-selector>
    <div class="row mb-1">
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { Action, Getter, Mutation } from "vuex-class";

import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'
import MoveSelector from '../components/moves/moveSelector.vue'

import { SendDamages, Result } from '../store/attacker/sendDamage'
import { Move, MoveLoader } from '../store/attacker/types'

const namespace: string = "attackerState";

// TODO:&結果一覧

@Component({
  components: {
    AdjustTarget,
    MoveSelector,
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
  @Getter('results', { namespace })
  private results!: Result[];
  @Getter('move', { namespace })
  private move!: string;
  @Mutation('setCurrent', { namespace })
  private setCurrent!: (current: number) => void;
  @Mutation('setMove', { namespace })
  private setMove!: (move: string) => void;
  @Mutation('setResults', { namespace })
  private setResults!: (results: Result[]) => void;

  private sendDamages: SendDamages = SendDamages.default();

  @Watch('index')
  changeTab() {
    this.setCurrent(this.index);
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

  @Watch('results')
  resultsTest() {
    console.log('ダメージ結果' + this.results.length);
  }

}
</script>

<style scoped>

</style>

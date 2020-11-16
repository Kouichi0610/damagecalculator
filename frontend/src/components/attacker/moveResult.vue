<template>
  <div class="move-result">
    Index:{{index}} Attacker:{{damages.attacker}}
    <moves :target="damages.attacker" @select="onSelect"></moves>
    わざ：{{ currentMove.name }}
  </div>
</template>

<script lang="ts">
/*
  ダメージ結果
*/
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
import { State, Action, Getter, Mutation } from "vuex-class";

import { Species } from '../target/store/species'
import { DefenderDamages, DefendersResult } from '../target/store/defenderDamages'
import { MoveInfo } from '../moves/store/types'
import Moves from '../moves/moves.vue'

const namespace: string = "attacker";

@Component({
  components:{
    Moves,
  },
})
export default class MoveResult extends Vue {
  @Prop() private index!: number;
  @Prop() private damages!: DefenderDamages;

  @Mutation("setCurrent", { namespace })
  private setCurrent!: (number) => void;
  @Action("setMove", { namespace })
  private setMove!: (MoveInfo) => void;

  @Getter("currentMove", { namespace })
  private currentMove!: MoveInfo;
  @Getter("moves", { namespace })
  private moves!: MoveInfo[];

  @Watch("damages", {deep: true})
  @Watch("currentMove", {deep: true})
  damageChanged() {
    if (this.currentMove.name.length == 0) {
      return;
    }
    this.damages.defenderDamages(this.currentMove.name)
    .then((results) => {
      console.log('change Move.' + this.currentMove.name);
    });
  }

  @Watch("index")
  indexChanged() {
    this.setCurrent(this.index);
  }

  created() {
    this.setCurrent(this.index);
  }

  onSelect(move: MoveInfo) {
    this.setMove(move);
  }
}
</script>

<style scoped>
</style>

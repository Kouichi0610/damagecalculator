<template>
  <div class="move-result">
    Index:{{index}} Attacker:{{damages.attacker}}
    <moves :target="damages.attacker" @select="onSelect"></moves>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
//import { State, Action, Getter, Mutation } from "vuex-class";

import { Species } from '../target/store/species'
import { DefenderDamages } from '../target/store/defenderDamages'
import { MoveInfo } from '../moves/store/types'
import Moves from '../moves/moves.vue'

@Component({
  components:{
    Moves,
  },
})
export default class MoveResult extends Vue {
  @Prop() private index!: number;
  @Prop() private damages!: DefenderDamages;
  private move: MoveInfo = MoveInfo.empty();

  @Watch("damages", {deep: true})
  @Watch("move", {deep: true})
  damageChanged() {
    if (this.move.name.length == 0) return;
    this.damages.defenderDamages(this.move.name);
  }

  onSelect(move: MoveInfo) {
    this.move = move;
  }

  
}
</script>

<style scoped>
</style>

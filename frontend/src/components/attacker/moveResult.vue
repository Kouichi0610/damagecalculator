<template>
  <div class="move-result">
    <moves :target="damages.attacker" @select="onSelect"></moves>
    わざ：{{ currentMove.name }}
    <div v-for="res in results" :key="res.Target">
      <div class="sample">Move:{{ res.toString() }} </div>
    </div>
    候補:{{ results.length }}
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
import DamageInfo from './components/damageInfo.vue'

const namespace: string = "attacker";

@Component({
  components:{
    Moves,
    DamageInfo,
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

  private results: DefendersResult[] = [];

  @Watch("damages", {deep: true})
  @Watch("currentMove", {deep: true})
  damageChanged() {
    if (this.currentMove.name.length == 0) {
      this.results = [];
      return;
    }
    this.damages.defenderDamages(this.currentMove.name)
    .then((results) => {
      this.results = results;
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

// 攻撃側
<template>
  <div class="attacker">
    Name:[{{ targetName }}]
    <target :show="true" @target="targetChanged"></target>
    <moves :target="targetName" @select="moveChanged"></moves>
    <move-result :index="index" :target="target.species"></move-result>
    <div class="row mb-1">
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import { State } from "vuex-class";
import Target from '../components/target/target.vue'
import Moves from '../components/moves/moves.vue'
import MoveResult from '../components/attacker/moveResult.vue'
import { MoveInfo } from '../components/moves/store/types'
import { TargetState } from '../components/target/store/types'

const namespace: string = "target";

@Component({
  components: {
    Target,
    Moves,
    MoveResult,
  }
})
export default class Attacker extends Vue {
  @State("target") target: TargetState;
  @Prop() private index!: number;
  private targetName: string = '';

  targetChanged(name: string) {
    this.targetName = name;
  }

  moveChanged(info: MoveInfo) {
    //this.target.species
    console.log('Move:' + info.name + ' ' + info.power);
  }
}
</script>

<style scoped>

</style>

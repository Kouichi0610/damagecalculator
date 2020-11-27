<template>
  <div class="speed-adjuster">
    <adjust-target :speedLock="true" @targetCondition="changeTargetCondition"></adjust-target>
    <b-form-checkbox id="check-trickroom" v-model="trickroom">{{ trickRoomMessage }}</b-form-checkbox>
    <p>素早さ{{targetSpeed}} TODO:コンポーネント</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { Action, Getter } from "vuex-class";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'

import { TargetCondition } from "../store/target/targetCondition";

//import { SpeedOrder } from '../store/speed/speedOrder'

const namespace: string = "speedState";

/*
  特性変更時に補正取得
  ぬめぬめ発動時
*/
@Component({
  components: {
    AdjustTarget,
  },
})
export default class Speed extends Vue {
  @Action('loadList', { namespace })
  private loadList!: (level: number) => void;
  @Action('getSpeed', { namespace })
  private getSpeed!: (target: TargetCondition) => void;

  @Getter('targetSpeed', { namespace })
  private targetSpeed!: number;

  private targetCondition: TargetCondition = TargetCondition.default();
  private trickroom: boolean = false;

  get trickRoomMessage(): string {
    if (this.trickroom) {
      return 'トリックルームON';
    }
    return 'トリックルームOFF';
  }

  @Watch('targetCondition.level')
  levelChanged() {
    this.loadList(this.targetCondition.level);
  }

  @Watch('targetCondition.ability')
  abilityChanged() {
    console.log('TODO:' + this.targetCondition.ability);
  }

  @Watch('targetCondition', {deep: true})
  calcSpeed() {
    if (!this.targetCondition.enable()) return;
    this.getSpeed(this.targetCondition);

    /*
    // 試作
    console.log('calcSpeed:');
    let s = new SpeedOrder();
    // 最初に保存
    s.speedList(this.targetCondition);
    // 特性変更時にthis.targetCondition.ability
    s.abilityOwner(this.targetCondition);
    s.abilityOther(this.targetCondition);
    */
  }

  changeTargetCondition(targetCondition: TargetCondition) {
    this.targetCondition = targetCondition;
  }
}
</script>
<style scoped>
</style>

<template>
  <div class="container">
    <div class="row mb-1">
    <adjust-target :speedLock="true" @targetCondition="changeTargetCondition"></adjust-target>
    </div>
    <div class="row mb-1">
      <b-form-checkbox id="check-trickroom" v-model="trickRoomFlag">{{ trickRoomMessage }}</b-form-checkbox>
    </div>
    <div class="row mb-1">
      素早さ{{targetSpeed}}
    </div>
    <div class="row mb-1">
      <div class="col-md-4">
        <speed-ranking-display :infos="nearRanking"></speed-ranking-display>
      </div>
      <div class="col-md-4">
        <speed-ranking-display :infos="speedRanking"></speed-ranking-display>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { Action, Getter, Mutation } from "vuex-class";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'
import SpeedRankingDisplay from '../components/speed/speedRankingDisplay.vue'

import { TargetCondition } from "../store/target/targetCondition";
import { SpeedInfo } from '../store/speed/types'

const namespace: string = "speedState";

/*
  すばやさ調整
  TODO:とくせい発動時の相手に影響を与える
*/
@Component({
  components: {
    AdjustTarget,
    SpeedRankingDisplay,
  },
})
export default class Speed extends Vue {
  @Action('initialize', { namespace })
  private initialize!: (target: TargetCondition) => void;
  @Action('loadList', { namespace })
  private loadList!: (level: number) => void;
  @Action('getSpeed', { namespace })
  private getSpeed!: (target: TargetCondition) => void;
  @Mutation('setTrickRoom', { namespace })
  private setTrickRoom!: (trickRoom: boolean) => void;

  @Getter('trickRoom', { namespace })
  private trickRoom!: boolean;
  @Getter('targetSpeed', { namespace })
  private targetSpeed!: number;

  @Getter('speedRanking', { namespace })
  private speedRanking!: SpeedInfo[];

  private targetCondition: TargetCondition = TargetCondition.default();
  private trickRoomFlag: boolean = false;

  created() {
    this.trickRoomFlag = this.trickRoom;
  }

  get trickRoomMessage(): string {
    if (this.trickRoom) {
      return 'トリックルーム使用中';
    }
    return 'トリックルーム未使用';
  }

  // 調整対象に近いところを表示
  get nearRanking(): SpeedInfo[] {
    let res: SpeedInfo[] = [];
    var i = 0;
    var first = 0;
    for (i = 0; i < this.speedRanking.length; i++) {
      if (this.speedRanking[i].target) {
        first = Math.max(0, i-2);
        break;
      }
    }
    var last = Math.min(first+5, this.speedRanking.length);
    for (i = first; i < last; i++) {
      res.push(this.speedRanking[i]);
    }
    return res;
  }

  @Watch('targetCondition.level')
  levelChanged() {
    this.loadList(this.targetCondition.level);
  }

  @Watch('targetCondition', {deep: true})
  changeCondition() {
    if (!this.targetCondition.enable()) return;
    this.initialize(this.targetCondition);
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

  @Watch('trickRoomFlag')
  changeTrickRoom() {
    this.setTrickRoom(this.trickRoomFlag);
  }

  changeTargetCondition(targetCondition: TargetCondition) {
    this.targetCondition = targetCondition;
  }
}
</script>
<style scoped>
</style>

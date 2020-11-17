<template>
  <div class="speed-order">
    <p>速度一覧</p>
    <template v-if="hasList">
      <template v-if="abilityOwnerCorrector.effective()">
        <b-form-checkbox id="check-ability-owner" v-model="useAbility">{{ ability }} {{ abilityOwnerCorrector.comment }}</b-form-checkbox>
      </template>
      <template v-if="abilityOtherCorrector.effective()">
        <b-form-checkbox id="check-ability-other" v-model="useAbility">{{ ability }} {{ abilityOtherCorrector.comment }}</b-form-checkbox>
      </template>
      <!-- TODO:items -->
      <div class="row mb-1">
        <b-dropdown class="col-1" id="item-dropdown" text="もちもの">
          <b-dropdown-item v-for="(item, index) in itemCorrectors" :key="item.comment" @click="usedItem = index">{{ item.comment }}</b-dropdown-item>
        </b-dropdown>
        <div class="col-1">
          {{itemCorrectors[usedItem].comment}}
        </div>
      </div>
 
      <div class="environment">
        <b-form-checkbox id="check-trickroom" v-model="trickRoom" name="check-trickroom">トリックルーム</b-form-checkbox>
      </div>

      <div class="row mb-1">
        <div class="col-4">
          <speed-ranking :ranking="nearDisplay"></speed-ranking>
        </div>
        <div class="col-4">
          <speed-ranking :ranking="display"></speed-ranking>
        </div>
      </div>
    </template>
    <template v-else>
    </template>
  </div>
</template>

<script lang="ts">
// 速度一覧
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';

import SpeedRanking from "./components/SpeedRanking.vue"
import { SpeedOrderState, SpeedInfo, SpeedCorrector } from './store/types';

const namespace: string ="speedOrder"

class InfoImpl implements SpeedInfo {
  info: string;     // 同速のポケモン
  speed: number;  // 素早さ(実数)
  constructor(info: string, speed: number) {
    this.info = info;
    this.speed = speed;
  }
}

@Component({
  components: {
    SpeedRanking,
  }
})
export default class SpeedOrder extends Vue {
  @Prop() private targetSpeed!: number;
  @Prop() private ability!: string;

  @State('speedOrder') speedOrder!: SpeedOrderState;

  @Action('getOrders', { namespace })
  private getOrders!: (level: number) => void;
  @Action('getAbilityEffect', { namespace })
  private getAbilityEffect!: (ability: string) => void;
  @Getter('hasList', { namespace })
  private hasList!: boolean;
  @Getter('list', { namespace })
  private list!: SpeedInfo[];

  @Getter('abilityOwnerCorrector', { namespace })
  private abilityOwnerCorrector!: SpeedCorrector;
  @Getter('abilityOtherCorrector', { namespace })
  private abilityOtherCorrector!: SpeedCorrector;
  @Getter('itemCorrectors', { namespace })
  private itemCorrectors!: SpeedCorrector[];

  private trickRoom: boolean = false;
  private useAbility: boolean = false;
  private usedItem: number = 0;

  created() {
    // TODO:level
    this.getOrders(50);
  }

  @Watch('ability')
  abilityChanged() {
    this.getAbilityEffect(this.ability);
    this.useAbility = false;
  }

  correctedTargetSpeed(): number {
    let res = this.targetSpeed;
    if (this.useAbility) {
      res = this.abilityOwnerCorrector.correct(res);
    }
    let item = this.itemCorrectors[this.usedItem];
    res = item.correct(res);
    return res;
  }

  get nearDisplay(): SpeedInfo[] {
    let empty: SpeedInfo[] = [];
    let disp: SpeedInfo[] = empty.concat(this.display);

    var idx = 0;
    for (var i = 0; i < disp.length; i++) {
      if (disp[i].info == '') {
        idx = i;
        break;
      }
    }

    const count = 5
    let start = (idx-count < 0) ? 0 : idx-count;
    let last = start + count*2;
    let end = (last < this.display.length) ? last : this.display.length-1;
    let res: SpeedInfo[] = [];
    for (i = start; i <= end; i++) {
      res.push(disp[i]);
    }
    return res;
  }
  get display(): SpeedInfo[] {
    let a: SpeedInfo[] = [new InfoImpl('', this.correctedTargetSpeed())];
    let disp:SpeedInfo[] = a.concat(this.correctedList());

    let decending = function(a: SpeedInfo, b: SpeedInfo) {
        if (a.speed > b.speed) return -1;
        if (a.speed < b.speed) return 1;
        return 0;
    }
    let ascending = function(a: SpeedInfo, b: SpeedInfo) {
        if (a.speed < b.speed) return -1;
        if (a.speed > b.speed) return 1;
        return 0;
    }
    let order = this.trickRoom ? ascending : decending;
    disp.sort(order);

    return disp;
  }

  correctedList(): SpeedInfo[] {
    let res: SpeedInfo[] = [];
    for (var i = 0; i < this.list.length; i++) {
      let info = this.list[i].info;
      let speed = this.list[i].speed;
      if (this.useAbility) {
        speed = this.abilityOtherCorrector.correct(speed);
      }
      res.push(new InfoImpl(info, speed));
    }
    return res;
  }

}

</script>

<style scoped>
</style>


<template>
  <div class="speed-order">
    <p>速度一覧x {{ targetSpeed }}</p>
    <p>サーバーから速度一覧を取得</p>
    <p>TODO:とくせい、もちものをサーバから取得</p>
    <p>TODO:表示</p>
    <template v-if="hasList">
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
import { Component, Vue, Prop } from 'vue-property-decorator';
import { State, Action, Getter } from 'vuex-class';

import SpeedRanking from "./components/SpeedRanking.vue"

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
  @Prop() private targetSpeed: number;

  @State('speedOrder') speedOrder: SpeedOrderState;

  @Action('getOrders', { namespace })
  private getOrders!: (number) => void;
  @Getter('hasList', { namespace })
  private hasList!: boolean;
  @Getter('list', { namespace })
  private list!: SpeedInfo[];

  private trickRoom: boolean = false;

  created() {
    // TODO:level
    this.getOrders(50);
  }

  get nearDisplay(): SpeedInfo[] {
    let disp = [].concat(this.display);

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
    let a: SpeedInfo[] = [new InfoImpl('', this.targetSpeed)];
    let disp:SpeedInfo[] = a.concat(this.list);

    let decending = function(a, b) {
        if (a.speed > b.speed) return -1;
        if (a.speed < b.speed) return 1;
        return 0;
    }
    let ascending = function(a, b) {
        if (a.speed < b.speed) return -1;
        if (a.speed > b.speed) return 1;
        return 0;
    }
    let order = this.trickRoom ? ascending : decending;
    disp.sort(order);

    return disp;
  }

}


</script>

<style scoped>
</style>


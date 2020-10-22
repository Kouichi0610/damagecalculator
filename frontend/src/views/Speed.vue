// 速度調整
<template>
  <div class="speed">
    <h1>速度調整</h1>
    <StatsEditor @speed="speedChanged"></StatsEditor>
    <template v-if="hasList && display.length > 1">
      <div class="row mb-1">
        <div class="col-4">
          <ul class="list-group" v-for="info in display" :key="info.info">
            <template v-if="info.info == ''">
              <li class="list-group-item list-group-item-primary">
                {{ speed }}
              </li>
            </template>
            <template v-else>
              <li class="list-group-item">
                {{ info.speed }} {{ info.info }}
              </li>
            </template>
          </ul>
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { State, Action, Getter } from 'vuex-class';
import Component from 'vue-class-component';

import StatsEditor from '../components/stats/statsEditor.vue'

const namespace: string = 'speedList';

class InfoImpl implements SpeedInfo {
  info: string;     // 同速のポケモン
  speed: number;  // 素早さ(実数)
  constructor(info: string, speed: number) {
    this.info = info;
    this.speed = speed;
  }
}

// TODO:トリックルーム(反転)
// TODO:すいすいようりょくそ(特性選択する必要が)
// TODO:スカーフ
@Component({
      components: {
          StatsEditor,
    },
})
export default class Speed extends Vue {
    @State('speedList') speedList: SpeedListState;

    @Getter('hasList', { namespace })
    private hasList!: boolean;
    @Getter('list', { namespace })
    private list!: SpeedInfo[];

    @Action('getList', { namespace })
    private getList!: () => Promise<boolean>;

    private speed: number = 0;

    created() {
      if (this.hasList) {
        return;
      }
      this.getList();
    }

    // TODO:↑5つ&↓5つあたり
    get display(): SpeedInfo[] {
      let a: SpeedInfo[] = [new InfoImpl('', this.speed)];
      let disp:SpeedInfo[] = a.concat(this.list);
      disp.sort(
        function(a, b) {
          if (a.speed > b.speed) return -1;
          if (a.speed < b.speed) return 1;
          return 0;
        }
      );
      return disp;
    }

    speedChanged(val: number) {
      this.speed = val;
    }

}
</script>

<style scoped>
.list-group-item {
  height:28px;
  font-size: 12px;
  text-align: center;
}

</style>

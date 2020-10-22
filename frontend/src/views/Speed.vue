// 速度調整
<template>
  <div class="speed">
    <h1>速度調整</h1>
    <StatsEditor @speed="speedChanged"></StatsEditor>
    <template v-if="hasList">
      <div class="row mb-1">
        <div class="col-6">
          <ul class="list-group" v-for="info in list" :key="info.speed">
            <li class="list-group-item">
              {{ info.speed }} {{ info.info }}
            </li>
          </ul>
        </div>
        <div class="col-2">
          <ul class="list-group" v-for="info in display" :key="info.speed">
            <li class="list-group-item">
              {{ info.speed }} {{ info.info }}
            </li>
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

    private display: InfoImpl[]

    created() {
      if (this.hasList) {
        for (var i = 0; i < this.list.length; i++) {
          let l = this.list[i];
          console.log('' + l.speed + ' info:' + l.info);
        }
        return;
      }
      this.getList();

    }

    speedChanged(val: number) {
      console.log('Speed:' + val);

      this.display = [];
      for (var i = 0; i < 10; i++) {
        this.display.push(new InfoImpl('', 0));
      }
      this.display.push(new InfoImpl('TODO:Mine', val));

      /*
      let a: SpeedInfo[] = [new InfoImpl('TODO:Mine', val)];

      this.display = a.concat(this.list);
      this.display.sort(
        function(a, b) {
          if (a.speed < b.speed) return -1;
          if (a.speed > b.speed) return 1;
          return 0;
        }
      );
      for (var i = 0; i < this.display.length; i++) {
        let d = this.display[i];
        console.log('' + d.speed + ':' + d.info);
      }
      */
    }

}
</script>

<style scoped>

</style>

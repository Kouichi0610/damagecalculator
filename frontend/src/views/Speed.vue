// 速度調整
<template>
  <div class="speed">
    <h1>速度調整</h1>
    <StatsEditor @speed="speedChanged"></StatsEditor>
    <template v-if="hasList">
      {{ list[1].info }}
    </template>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { State, Action, Getter } from 'vuex-class';
import Component from 'vue-class-component';

import StatsEditor from '../components/stats/statsEditor.vue'

const namespace: string = 'speedList';

// TODO:速度の仮想敵一覧
// TODO:vuex経由で取得しておく。
// TODO:トリックルーム(反転)
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

    created() {
      if (this.hasList) {
        for (var i = 0; i < this.list.length; i++) {
          let l = this.list[i];
          console.log('' + l.species + ' info:' + l.info);
        }
        return;
      }
      this.getList();

    }

    private get Speed(): number {
      return 0;
    }

    speedChanged(val: number) {
      console.log('Speed:' + val);
    }

}
</script>

<style scoped>

</style>

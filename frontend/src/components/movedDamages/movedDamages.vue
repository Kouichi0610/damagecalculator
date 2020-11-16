<template>
  <div class="moved-damages">
    ダメージ一覧
    {{ move }}
    <div class="row mb-1">
      <b-progress class="col-6" :max="100" show-value animated>
        <b-progress-bar :value="min" :label="alpha" variant="danger"></b-progress-bar>
        <b-progress-bar :value="max - min" variant="warning"></b-progress-bar>
      </b-progress>
      <div class="col-4">
        {{ min }}～{{ max }}%
      </div>
    </div>
  </div>
</template>

<script lang="ts">
// class="col-2"
import { Vue, Component } from "vue-property-decorator";

import { DefenderDamages } from '../target/store/defenderDamages'

/*
  わざと仮想敵へのダメージ一覧

  name (types) damegepercent 確定数
*/
@Component
export default class MovedDamages extends Vue {
  private move: string = "はかいこうせん";
  private min: number = 43;
  private max: number = 50;

  created() {
    //this.getResult();
  }
  /*
  async getResult() {
    let d = new DefenderDamages('デンジュモク', 50, 'Max', [0,0,0,252,0,252], 'ひかえめ', 'ビーストブースト', 'なし', '正常', '', 'エレキフィールド');
    let res = await d.defenderDamages('かみなりパンチ');
    for (var i = 0; i < res.length; i++) {
      let x = res[i]
      console.log('' + x.Target + ' Damage:' + x.DamageMin + '-' + x.DamageMax + ' Rate:' + x.RateMin + '-' + x.RateMax + ' 確定数:' + x.DetermineCount);
    }
  }
  */
}
</script>

<style scoped>
</style>


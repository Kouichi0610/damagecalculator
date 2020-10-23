/*
  素早さに関係ある持ち物、とくせい、天候など

  すばやさ関係ある特性
  かそく  (ターン経過で1ランクずつ上昇)
  すいすい (あめの時素早さ2倍)
  すなかき (すなあらしの時素早さ2倍)
  スロースタート(5ターンこうげき、素早さ半分)
  でんきエンジン(電気無効、電気タイプのわざをうけると素早さ1ランク上昇)
  ぬめぬめ(直接攻撃を受けると相手のすばやさランク1下げる)
  ゆきかき(あられのときすばやさ2倍)
  ようりょくそ(はれのとき素早さ2倍)
  わたげ(こうげきを受けると、自分以外の全てのポケモンの素早さランク1下げる)

*/
<template>
  <div class="speed-environment">
    <p>元の素早さ:{{ baseSpeed }} => 持ち物、天候補正:{{ correctedSpeed }}</p>
    <label><input type="radio" v-model.number="item" value="0">持ち物なし</label>
    <label><input type="radio" v-model.number="item" value="1">こだわりスカーフ</label>
    <label><input type="radio" v-model.number="item" value="2">くろいてっきゅう</label>
    <!-- TODO:とくせいが有効か取得する -->
    <template v-if="true">
      <b-form-checkbox id="check-ability" v-model="ability" name="check-ability">すいすい、ようりょくそ</b-form-checkbox>
    </template>
  </div>
 
</template>

<script lang="ts">
import { Prop, Vue } from 'vue-property-decorator';
import Component from 'vue-class-component';

@Component
export default class SpeedEnvironment extends Vue {
  item: number = 0;
  ability: boolean = false;

  @Prop()
  private baseSpeed!: number;

  get correctedSpeed(): number {
    var res = this.baseSpeed;
    res = res * this.itemCorrect();
    if (this.ability) {
      res *= 2;
    }
    res = Math.floor(res);
    console.log('notice speed' + res);
    this.$emit('correctedSpeed', res);
    return res;
  }

  private itemCorrect(): number {
    if (this.item == 1) {
      return 1.5;
    }
    if (this.item == 2) {
      return 0.5;
    }
    return 1.0;
  }

  created() {
  }

}
</script>

<style scoped>
</style>

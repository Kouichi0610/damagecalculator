<template>
  <div class="move-result">
    耐久調整
    <div v-for="res in results" :key="res.Target">
      <damage-rate :result="res"></damage-rate>
    </div>
  </div>
</template>

<script lang="ts">
// 耐久調整 ダメージ結果
import { Vue, Prop, Component, Watch } from "vue-property-decorator";

import { AttackerDamages, AttackersResult } from '../target/store/attackerDamages'
import DamageRate from './components/damageRate.vue'

//const namespace: string = "defender";

@Component({
  components: {
    DamageRate,
  },
})
export default class MoveResult extends Vue {
  @Prop() private damages!: AttackerDamages;

  private results: AttackersResult[] = [];

  @Watch('damages', { deep: true })
  damagesChanged() {
    this.damages.attackerDamages()
    .then((results) => {
      this.results = results;
    });
  }

  created() {
    this.damagesChanged();
  }
}
</script>

<style scoped>
</style>

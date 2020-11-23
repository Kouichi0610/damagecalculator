<template>
  <div class="adjust-target">
    <template v-if="hasTarget">
      <p>試作 {{ nature.name }}</p>
      <p>{{ individuals.toString() }}</p>
      <p>{{ basePoints.toString() }}</p>
      <nature-selector :target="target" @change="changeNature"></nature-selector>
      <individuals-selector :target="target" @change="changeIndividuals"></individuals-selector>
      <base-points-adjuster :target="target" @change="changeBasePoints"></base-points-adjuster>
    </template>
    <template v-else>
      No Target.
    </template>
  </div>
</template>

<script lang="ts">
/*
  調整対象
  TODO:target.vueを改修してこっちを使用
  TODO:components分割
  TODO:target更新時、性格、個体値などリセット(各コンポーネントにWatchさせるのが無難か)
  TODO:肥大化しているstoreを細分化
  TODO:attackerDamages, defenderDamages, 速度調整のクラス作成の役割を持つ(extend試作)
    名前
    特性
    タイプ
    性格  DONE
    個体値 DONE
    種族値
    基礎ポイント DONE
    能力値
    天候＆フィールド
    技一覧
    もちもの
    状態異常
*/
import { Vue, Component, Watch } from 'vue-property-decorator';

import NatureSelector from './components/natureSelector.vue'
import IndividualsSelector from './components/individualsSelector.vue'
import BasePointsAdjuster from './components/basePointsAdjuster.vue'

import { Nature } from '../../store/nature/types'
import { Individuals } from '../../store/individuals/types'
import { IBasePoints, defaultBasePoints } from '../../store/basePoints/types'

@Component({
  components: {
    NatureSelector,
    IndividualsSelector,
    BasePointsAdjuster,
  }
})
export default class AdjustTarget extends Vue {

  private target: string = 'フカマル';

  private nature: Nature = Nature.default();
  private individuals: Individuals = Individuals.default();
  private basePoints: IBasePoints = defaultBasePoints();

  get hasTarget(): boolean {
    return true;
  }

  changeNature(nature: Nature) {
    this.nature = nature;
  }
  changeIndividuals(individuals: Individuals) {
    this.individuals = individuals;
  }
  changeBasePoints(basePoints: IBasePoints) {
    this.basePoints = basePoints;
  }

  created() {
    //this.$emit('sample', "AdjustSample");

  }
}
</script>

<style scoped>
</style>

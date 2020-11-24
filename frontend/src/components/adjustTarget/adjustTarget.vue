<template>
  <div class="adjust-target">
    <template v-if="hasTarget">
      <p>対象:{{ nature.name }}</p>
      <nature-selector :target="target" @change="changeNature"></nature-selector>
      <individuals-selector :target="target" @change="changeIndividuals"></individuals-selector>
      <div class="row mb-1">
        <species-display class="col-3" :species="species"></species-display>
        <stats-display class="col-3" :stats="stats"></stats-display>
        <base-points-adjuster class="col-3" :target="target" @change="changeBasePoints"></base-points-adjuster>
      </div>
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
  TODO:target更新時、性格、個体値などリセット(各コンポーネントにWatchさせるのが無難か)
  TODO:肥大化しているstoreを細分化
    名前
    特性
    タイプ
    おもさ
    性格  DONE
    個体値 DONE
    種族値 DONE
    基礎ポイント DONE
    能力値 DONE
    天候＆フィールド
    技一覧
    もちもの
    状態異常
*/
import { Vue, Component, Watch } from 'vue-property-decorator';

import NatureSelector from './components/natureSelector.vue'
import IndividualsSelector from './components/individualsSelector.vue'
import BasePointsAdjuster from './components/basePointsAdjuster.vue'
import SpeciesDisplay from './components/speciesDisplay.vue'
import StatsDisplay from './components/statsDisplay.vue'

import { Nature } from '../../store/nature/types'
import { Individuals } from '../../store/individuals/types'
import { IBasePoints, defaultBasePoints } from '../../store/basePoints/types'
import { Species } from '../../store/species/types'
import { StatsPatterns, StatsPatternsLoader } from '../../store/stats/types'

@Component({
  components: {
    NatureSelector,
    IndividualsSelector,
    BasePointsAdjuster,
    SpeciesDisplay,
    StatsDisplay,
  }
})
export default class AdjustTarget extends Vue {
  // TODO:store
  private target: string = 'ガブリアス';
  private nature: Nature = Nature.default();
  private individuals: Individuals = Individuals.default();
  private basePoints: IBasePoints = defaultBasePoints();
  private species: Species = new Species(100, 80, 60, 70, 130, 50);
  private statsPatterns: StatsPatterns = StatsPatterns.default();

  get hasTarget(): boolean {
    return true;
  }

  get stats(): number[] {
    return this.statsPatterns.statsArray(this.basePoints);
  }

  @Watch('target')
  @Watch('nature', { deep: true })
  @Watch('individuals', { deep: true })
  @Watch('basePoints', { deep: true })
  async loadStatsPatterns() {
    let loader = new StatsPatternsLoader();
    this.statsPatterns = await loader.load(50, this.target, this.individuals, this.nature);
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

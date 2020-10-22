// 
<template>
  <div class="statsEditor">
    <template v-if="hasTarget">
      <nature :selected="targetNature" @changed="updateNature"></nature>
      <individuals
        :params="targetIndividuals"
        @changed="updateIndividuals"
      ></individuals>
      <div class="row mb-1">
        <stats class="col-4" :params="stats"></stats>
        <species class="col-2" :params="species"></species>
        <base-points
          class="col-2"
          :params="targetBasePoints"
          @changed="updateBasePoints"
        ></base-points>
      </div>
    </template>
    <template v-else>
      <p>対象を選択してください</p>
      <button type="button" class="btn btn-warning" @click="toSelect">対象選択</button>
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import Nature from "./components/nature.vue";
import BasePoints from "./components/basePoints.vue";
import Individuals from "./components/individuals.vue";
import Species from "./components/species.vue";
import Stats from "./components/stats.vue";
import { State, Getter, Mutation } from "vuex-class";

//import * as nature from "../../domain/nature";
//import { INature } from "../../domain/nature"
import * as stats from "../../domain/stats";
import router from '../../router'

const namespace: string = "targets";

// TODO: $emit stats
// TODO:名前被るのでSpeciesComponent当たりに(変えなくても動作はする)
@Component({
  components: {
    Nature,
    BasePoints,
    Individuals,
    Species,
    Stats,
  },
})
export default class StatsEditor extends Vue {
  @State('targets') targets: TargetsState;

  @Getter('hasTarget', { namespace })
  private hasTarget!: boolean;
  @Getter('targetNature', { namespace })
  private targetNature!: INature;
  @Getter('targetSpecies', { namespace })
  private species!: number[];
  @Getter('targetIndividuals', { namespace })
  private targetIndividuals!: number[];
  @Getter('targetBasePoints', { namespace })
  private targetBasePoints!: number[];

  @Mutation('setTargetNature', { namespace })
  private setTargetNature!: (INature) => void;
  @Mutation('setTargetIndividuals', { namespace })
  private setTargetIndividuals!: (Individuals) => void;
  @Mutation('setTargetBasePoints', { namespace })
  private setTargetBasePoints!: (BasePoints) => void;

  private stats: number[] = [0, 0, 0, 0, 0, 0];
  private calculator: stats.StatsCalculator = new stats.StatsCalculator();

  created() {
    this.calcStats();
  }

  private updateNature(nature: INature) {
    this.setTargetNature(nature);
    this.calcStats();
  }
  private updateIndividuals(params: number[]) {
    this.setTargetIndividuals({hp: params[0], attack: params[1], defense: params[2], spAttack: params[3], spDefense: params[4], speed: params[5]});
    this.calcStats();
  }
  private updateBasePoints(params: number[]) {
    this.setTargetBasePoints({hp: params[0], attack: params[1], defense: params[2], spAttack: params[3], spDefense: params[4], speed: params[5]});
    this.calcStats();
  }

  private calcStats(): number[] {
    let oldSpeed = this.stats[5];

    this.stats = this.calculator.Calculate(
      this.targetNature,
      50,
      this.species,
      this.targetIndividuals,
      this.targetBasePoints
    );

    let newSpeed = this.stats[5];
    if (oldSpeed != newSpeed) {
      this.$emit('speed', this.stats[5]);
    }
   return this.stats;
  }

  private toSelect() {
    router.push({ path: 'select' });
  }
}
</script>

<style>
</style>
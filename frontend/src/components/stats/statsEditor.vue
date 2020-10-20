// 
<template>
  <div class="statsEditor">
    <p>Indi{{ targetIndividuals }}</p>
    <p>Base{{ targetBasePoints }}</p>
    <nature :selected="nature" @changed="updateNature"></nature>
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
  </div>
</template>

/*
*/

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import Nature from "./components/nature.vue";
import BasePoints from "./components/basePoints.vue";
import Individuals from "./components/individuals.vue";
import Species from "./components/species.vue";
import Stats from "./components/stats.vue";
import { State, Getter, Mutation } from "vuex-class";

import * as nature from "../../domain/nature";
import * as stats from "../../domain/stats";

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

  @Getter('targetSpecies', { namespace })
  private species: number[];
  @Getter('targetIndividuals', { namespace })
  private targetIndividuals: number[];
  @Getter('targetBasePoints', { namespace })
  private targetBasePoints: number[];

  @Mutation('setTargetIndividuals', { namespace })
  private setTargetIndividuals: (Individuals) => void;
  @Mutation('setTargetBasePoints', { namespace })
  private setTargetBasePoints: (BasePoints) => void;


  private stats: number[] = [0, 0, 0, 0, 0, 0];
  private nature: nature.INature = nature.NatureFactory('てれや');
  private calculator: stats.StatsCalculator = new stats.StatsCalculator();

  created() {
    this.calcStats();
  }

  private get natureName(): string {
    return this.nature.Name();
  }

  private updateNature(nature: nature.INature) {
    this.nature = nature;
    console.log("next:" + this.nature.Name());
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
    this.stats = this.calculator.Calculate(
      this.nature,
      50,
      this.species,
      this.targetIndividuals,
      this.targetBasePoints
    );
    console.log("Stats:" + this.stats);
    return this.stats;
  }
}
</script>

<style>
</style>
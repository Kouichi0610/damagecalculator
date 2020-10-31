/*
  対象となるポケモンの情報を表示する
  種族値
  とくせい
  わざ
*/
<template>
  <div class="target">
    <template v-if="hasTarget">
      <div>対象:{{ name }}</div>
      <div>{{ types }}</div><!-- TODO:component -->
      <div>{{ weight }}kg</div><!-- TODO:component -->
      <nature @changed="changeNature"></nature>
      <individuals-adjuster :individuals="individuals" @slowest="changeSlowest" @weakest="changeWeakest"></individuals-adjuster>
      <div class="row mb-1">
        <species-display class="col-2" :species="species"></species-display>
        <base-points-adjuster class="col-2" :basePoints="basePoints" @changed="changeBasePoints"></base-points-adjuster>
        <stats-display class="col-4" :hp="hp" :attack="attack" :defense="defense" :spAttack="spAttack" :spDefense="spDefense" :speed="speed"></stats-display>
      </div>
      <div></div>
    </template>
    <template v-else>
      <div>No Target.</div>
    </template>
  </div>
 </template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from 'vuex-class';

//import {Species} from './store/types'
import IndividualsAdjuster from './components/individualsAdjuster.vue'
import SpeciesDisplay from './components/speciesDisplay.vue'
import BasePointsAdjuster from './components/basePointsAdjuster.vue'
import Nature from '../nature/nature.vue';
import StatsDisplay from './components/statsDisplay.vue'

import {StatsPatternArgs} from "./store/types"

const namespace: string = "target";

@Component({
  components: {
    IndividualsAdjuster,
    SpeciesDisplay,
    BasePointsAdjuster,
    Nature,
    StatsDisplay,
  }
})
export default class Target extends Vue {
  @Prop() private targetName: string;

  @State('target') target: TargetState;

  @Action('getSpecies', { namespace })
  private getSpecies!: (string) => Promise<boolean>;
  @Action('getStatsPattern', { namespace })
  private getStatsPattern!: (StatsPatternArgs) => void;

  @Getter('level', { namespace })
  private level!: number;
  @Getter('hasTarget', { namespace })
  private hasTarget!: boolean;
  @Getter('name', { namespace })
  private name!: string;
  @Getter('nature', { namespace })
  private nature!: string;
  @Getter('species', { namespace })
  private species!: Species;
  @Getter('types', { namespace })
  private types!: string[];
  @Getter('weight', { namespace })
  private weight: number;
  @Getter('individuals', { namespace })
  private individuals: Individuals;
  @Getter('basePoints', { namespace })
  private basePoints: BasePoints;
  @Getter('hp', { namespace })
  private hp: number;
  @Getter('attack', { namespace })
  private attack: number;
  @Getter('defense', { namespace })
  private defense: number;
  @Getter('spAttack', { namespace })
  private spAttack: number;
  @Getter('spDefense', { namespace })
  private spDefense: number;
  @Getter('speed', { namespace })
  private speed: number;


  @Mutation('changeSlowest', { namespace })
  private changeSlowest!: (boolean) => void;
  @Mutation('changeWeakest', { namespace })
  private changeWeakest!: (boolean) => void;
  @Mutation('changeBasePoints', { namespace })
  private changeBasePoints!: (BasePoints) => void;
  @Mutation('changeNature', { namespace })
  private changeNature!: (string) => void;

  @Watch('name')
  @Watch('individuals')
  @Watch('nature')
  private getCalculate() {
    console.log('this.individuals:' + this.individuals);
    if (this.name.length == 0) {
      return;
    }
    if (this.nature.length == 0) {
      return;
    }
    let args = new StatsPatternArgs(this.level, this.name, this.nature,
     this.individuals.hp,
     this.individuals.at,
     this.individuals.df,
     this.individuals.sa,
     this.individuals.sd,
     this.individuals.sp);
    this.getStatsPattern(args);
  }

  @Watch('targetName')
  private nameChanged() {
    if (this.targetName == '') {
      return;
    }
    this.getSpecies(this.targetName);
  }
}
</script>

<style scoped>
</style>

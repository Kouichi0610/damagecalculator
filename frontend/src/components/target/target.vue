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
      <div>{{ species }}</div><!-- TODO:component -->
      <div>{{ types }}</div><!-- TODO:component -->
      <div>{{ weight }}</div><!-- TODO:component -->
      <individuals-adjuster :individuals="individuals" @slowest="changeSlowest" @weakest="changeWeakest"></individuals-adjuster>
      <div class="row mb-1">
        <species-display class="col-2" :species="species"></species-display>
        <base-points-adjuster class="col-2" :basePoints="basePoints" @changed="changeBasePoints"></base-points-adjuster>
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

const namespace: string = "target";

@Component({
  components: {
    IndividualsAdjuster,
    SpeciesDisplay,
    BasePointsAdjuster,
  }
})
export default class Target extends Vue {
  @State('target') target: TargetState;
  @Action('getSpecies', { namespace })
  private getSpecies!: (string) => Promise<boolean>;

  // この名前の情報を取ってくる
  @Prop() private targetName: string;

  @Getter('hasTarget', { namespace })
  private hasTarget!: boolean;
  @Getter('name', { namespace })
  private name!: string;
  @Getter('species', { namespace })
  private species!: Species;
  @Getter('types', { namespace })
  private types!: string[];
  @Getter('weight', { namespace })
  private weight: number;
  //@Getter('nature', { namespace })// TODO:性格リストサーバーから取る
  //private nature: INature;
  @Getter('individuals', { namespace })
  private individuals: Individuals;
  @Getter('basePoints', { namespace })
  private basePoints: BasePoints;

  @Mutation('changeSlowest', { namespace })
  private changeSlowest!: (boolean) => void;
  @Mutation('changeWeakest', { namespace })
  private changeWeakest!: (boolean) => void;
  @Mutation('changeBasePoints', { namespace })
  private changeBasePoints!: (BasePoints) => void;
  

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

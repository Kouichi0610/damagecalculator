/*
  対象となるポケモンの情報を表示する
  種族値
  とくせい
  わざ
*/
<template>
  <div class="target">
    <template v-if="show == false"> </template>
    <template v-else>
      <template v-if="species.hasTarget()">
        <div>対象:{{ name }}</div>
        <ability-selector :abilities="species.abilities()" :current="currentAbility" @changed="setCurrentAbility"></ability-selector>
        <!-- TODO:component -->
        <div>{{ species.types() }}</div>
        <!-- TODO:component -->
        <div>{{ species.weight() }}kg</div>
        <nature @changed="changeNature"></nature>
        <individuals-adjuster
          :individuals="individuals"
          @slowest="changeSlowest"
          @weakest="changeWeakest"
        ></individuals-adjuster>
        <div class="row mb-1">
          <species-display class="col-2" :species="species"></species-display>
          <base-points-adjuster
            class="col-2"
            :basePoints="basePoints"
            @changed="changeBasePoints"
          ></base-points-adjuster>
          <stats-display
            class="col-4"
            :hp="hp"
            :attack="attack"
            :defense="defense"
            :spAttack="spAttack"
            :spDefense="spDefense"
            :speed="speed"
          ></stats-display>
        </div>
        <div></div>
      </template>
      <template v-else>
        <div>No Target.</div>
      </template>
    </template>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
import { State, Action, Getter, Mutation } from "vuex-class";

//import {Species} from './store/types'
import IndividualsAdjuster from "./components/individualsAdjuster.vue";
import SpeciesDisplay from "./components/speciesDisplay.vue";
import BasePointsAdjuster from "./components/basePointsAdjuster.vue";
import Nature from "../nature/nature.vue";
import StatsDisplay from "./components/statsDisplay.vue";
import AbilitySelector from './components/abilitySelector.vue'

import { StatePatternsLoader } from './store/statePattern'
import { Species, SpeciesLoader } from './store/species'
import { Individuals } from './store/individuals'
//import { TargetState } from './store/types'

const namespace: string = "target";

@Component({
  components: {
    IndividualsAdjuster,
    SpeciesDisplay,
    BasePointsAdjuster,
    Nature,
    StatsDisplay,
    AbilitySelector,
  },
})
export default class Target extends Vue {
  @Prop() private targetName: string;
  @Prop() private show: boolean;
  //@State("target") target: TargetState;

  @Action("getSpecies", { namespace })
  private getSpecies!: (SpeciesLoader) => void;
  @Action("getStatsPattern", { namespace })
  private getStatsPattern!: (StatePatternsLoader) => void;
  @Action("setCurrentAbility", { namespace })
  private setCurrentAbility!: (string) => void;

  @Getter("level", { namespace })
  private level!: number;
  @Getter("species", { namespace })
  private species!: Species;
  @Getter("currentAbility", { namespace })
  private currentAbility!: string;

  @Getter("nature", { namespace })
  private nature!: string;
  @Getter("individuals", { namespace })
  private individuals: Individuals;
  @Getter("basePoints", { namespace })
  private basePoints: BasePoints;
  @Getter("hp", { namespace })
  private hp: number;
  @Getter("attack", { namespace })
  private attack: number;
  @Getter("defense", { namespace })
  private defense: number;
  @Getter("spAttack", { namespace })
  private spAttack: number;
  @Getter("spDefense", { namespace })
  private spDefense: number;
  @Getter("speed", { namespace })
  private speed: number;

  @Mutation("changeSlowest", { namespace })
  private changeSlowest!: (boolean) => void;
  @Mutation("changeWeakest", { namespace })
  private changeWeakest!: (boolean) => void;
  @Mutation("changeBasePoints", { namespace })
  private changeBasePoints!: (BasePoints) => void;
  @Mutation("changeNature", { namespace })
  private changeNature!: (string) => void;

  @Watch("name")
  @Watch("individuals", { deep: true })
  @Watch("nature")
  private getCalculate() {
    if (!this.species.hasTarget()) {
      return;
    }
    let args = new StatePatternsLoader(
      this.level,
      this.species.name(),
      this.nature,
      this.individuals
    );
    
    this.getStatsPattern(args);
  }

  @Watch("currentAbility")
  private abilityChanged() {
    this.$emit('ability', this.currentAbility);
  }

  @Watch("targetName")
  private async nameChanged() {
    if (this.targetName == "") {
      return;
    }
    await this.getSpecies(new SpeciesLoader(this.targetName));
    this.$emit('changeTarget', this.targetName);
  }

  @Watch("hp")
  private hpChanged() {
    this.$emit('hp', this.hp);
  }
  @Watch("attack")
  private attackChanged() {
    this.$emit('attack', this.attack);
  }
  @Watch("defense")
  private defenseChanged() {
    this.$emit('defense', this.defense);
  }
  @Watch("spAttack")
  private spAttackChanged() {
    this.$emit('spAttack', this.spAttack);
  }
  @Watch("spDefense")
  private spDefenseChanged() {
    this.$emit('spDefense', this.spDefense);
  }
  @Watch("speed")
  private speedChanged() {
    this.$emit('speed', this.speed);
  }

  created() {
    this.$emit('target', this.species.name());
    this.$emit('hp', this.hp);
    this.$emit('attack', this.attack);
    this.$emit('defense', this.defense);
    this.$emit('spAttack', this.spAttack);
    this.$emit('spDefense', this.spDefense);
    this.$emit('speed', this.speed);
    this.$emit('ability', this.currentAbility);
  }
}
</script>

<style scoped>
</style>

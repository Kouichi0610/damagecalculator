<template>
  <div class="container">
    <species-loader @target="onTarget"></species-loader>
    <template v-if="hasTarget">
      <data-display :data="data"></data-display>
      <ability-selector :current="ability" :abilities="abilities" @change="changeAbility"></ability-selector>
      <nature-selector :target="target" @change="changeNature"></nature-selector>
      <individuals-selector :target="target" @change="changeIndividuals"></individuals-selector>
      <div class="row">
        <species-display class="col-md-4" :species="species"></species-display>
        <stats-display class="col-md-4" :stats="stats"></stats-display>
        <base-points-adjuster class="col-md-4" :target="target" :speedLock="speedLock" @change="changeBasePoints"></base-points-adjuster>
      </div>
      <weather-field-selector :target="target" @weather="changeWeather" @field="changeField"></weather-field-selector>
      <item-selector :target="target" @change="changeItem"></item-selector>
    </template>
    <template v-else>
      No Target.
    </template>
  </div>
</template>

<script lang="ts">
/*
  調整対象の能力値表示
  TODO:AdjustTarget -> TargetAdjuster
  TODO:状態異常
*/
import { Vue, Component, Prop, Watch } from 'vue-property-decorator';

import SpeciesLoader from '../species/speciesLoader.vue'
import NatureSelector from './components/natureSelector.vue'
import IndividualsSelector from './components/individualsSelector.vue'
import BasePointsAdjuster from './components/basePointsAdjuster.vue'
import SpeciesDisplay from './components/speciesDisplay.vue'
import StatsDisplay from './components/statsDisplay.vue'
import AbilitySelector from './components/abilitySelector.vue'
import DataDisplay from './components/dataDisplay.vue'
import WeatherFieldSelector from './components/weatherFieldSelector.vue'
import ItemSelector from './components/itemSelector.vue'

import { Nature } from '../../store/nature/types'
import { Individuals } from '../../store/individuals/types'
import { IBasePoints, defaultBasePoints, BasePoints } from '../../store/basePoints/types'
import { ISpecies, PokeData } from '../../store/species/types'
import { StatsPatterns, StatsPatternsLoader } from '../../store/stats/types'
import { Item } from '../../store/items/types'
import { Weather, Field } from '../../store/weatherFields/types'
import { TargetCondition } from '../../store/target/targetCondition'

@Component({
  components: {
    DataDisplay,
    SpeciesLoader,
    NatureSelector,
    IndividualsSelector,
    BasePointsAdjuster,
    SpeciesDisplay,
    StatsDisplay,
    AbilitySelector,
    WeatherFieldSelector,
    ItemSelector,
  }
})
export default class AdjustTarget extends Vue {
  @Prop()
  speedLock!: boolean;
  
  private data: PokeData = PokeData.default();
  private nature: Nature = Nature.default();
  private individuals: Individuals = Individuals.default();
  private basePoints: IBasePoints = defaultBasePoints();
  private statsPatterns: StatsPatterns = StatsPatterns.default();
  private ability: string = '';
  private item: Item = Item.default();
  private weather: Weather = Weather.default();
  private field: Field = Field.default();

  get targetCondition(): TargetCondition {
    return new TargetCondition(
      this.data.name,
      50,
      this.individuals,
      this.basePoints,
      this.ability,
      this.nature,
      this.item,
      'なし', // condition
      this.weather,
      this.field
    );
  }

  @Watch('targetCondition', { deep: true })
  changeCondition(after: TargetCondition, before: TargetCondition) {
    this.$emit('targetCondition', after);
  }

  get abilities(): string[]{
    return this.data.abilities;
  }
  get target(): string {
    return this.data.name;
  }
  get hasTarget(): boolean {
    return this.data.hasTarget();
  }
  get species(): ISpecies {
    return this.data.species;
  }
  get stats(): number[] {
    return this.statsPatterns.statsArray(this.basePoints);
  }

  @Watch('loader', { deep: true })
  async loaderChanged() {
    if (!this.loader.enable()) return;
    this.statsPatterns = await this.loader.load();
  }

  get loader(): StatsPatternsLoader {
    return new StatsPatternsLoader(50, this.data.name, this.individuals, this.nature);
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
  changeAbility(ability: string) {
    this.ability = ability;
  }
  changeItem(item: Item) {
    this.item = item;
  }
  changeWeather(weather: Weather) {
    this.weather = weather;
  }
  changeField(field: Field) {
    this.field = field;
  }

  onTarget(data: PokeData) {
    this.data = data;
  }

}
</script>

<style scoped>
</style>

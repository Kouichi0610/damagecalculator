<template>
  <div class="species">
    <p>種族値</p>
    <div v-for="param in params" :key="param.key" class="row mb-1">
        <div class="col-sm-6 pt-1">
          <b-progress :max="255" show-value animated>
            <b-progress-bar :value="param.value" variant="warning"></b-progress-bar>
          </b-progress>
        </div>
        <div class="col-1">{{ param.value }}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { ISpecies } from '../../../store/species/types'

class Params {
  readonly key: number;
  readonly value: number;
  constructor(key: number, value: number) {
    this.key = key;
    this.value = value;
  }
}

@Component
export default class SpeciesDisplay extends Vue {
  @Prop()
  private species!: ISpecies;

  private params: Params[] = [];

  created() {
    this.speciesChanged();
  }

  @Watch('species', { deep: true})
  speciesChanged() {
    this.params = [];
    this.params.push(new Params(0, this.species.hp));
    this.params.push(new Params(1, this.species.attack));
    this.params.push(new Params(2, this.species.defense));
    this.params.push(new Params(3, this.species.spAttack));
    this.params.push(new Params(4, this.species.spDefense));
    this.params.push(new Params(5, this.species.speed));
  }
}
</script>

<style scoped>
</style>

<template>
  <div class="species-loader"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { State, Getter, Mutation, Action } from 'vuex-class';

import { PokeData } from '../../store/species/types'

const namespace: string = "speciesState";

@Component
export default class SpeciesLoader extends Vue {
  @Prop()
  private target!: string;

  @Action('load', { namespace })
  private load!: (name: string) => void;
  @Getter('pokeData', { namespace })
  private pokeData!: PokeData;

  @Watch('target')
  targetChanged() {
    this.load(this.target);
  }

  created() {
    this.$emit('target', this.pokeData);
  }

  @Watch('pokeData', { deep: true })
  dataChanged() {
    this.$emit('change', this.pokeData);
  }
}
</script>

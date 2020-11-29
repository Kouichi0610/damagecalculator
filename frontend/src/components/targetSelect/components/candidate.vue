<template>
  <div class="candidate">
    <li class="list-group-item">
      <b-button class="TypeButton" v-on:click="onButton">{{ candidate.name }}</b-button>
      {{ candidate.types }} HP:{{candidate.hp}} AT:{{candidate.attack}} DF:{{candidate.defense}} SA:{{candidate.spAttack}} SD:{{candidate.spDefense}} SP:{{candidate.speed}}
    </li>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import { Species } from '../../../store/targetSelect/types'

@Component
export default class Candidate extends Vue {
  @Prop()
  private candidate!: Species;

  onButton() {
    this.$emit('selected', this.candidate.name);
  }
}
</script>

<style scoped>
</style>

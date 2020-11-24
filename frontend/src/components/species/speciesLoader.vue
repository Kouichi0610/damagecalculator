<template>
  <div class="species-loader"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { State, Getter, Mutation, Action } from 'vuex-class';

@Component
export default class SpeciesLoader extends Vue {
  @Prop()
  private target!: string;
  
}
</script>

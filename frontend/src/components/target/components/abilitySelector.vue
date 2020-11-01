<template>
  <div class="ability-selector">
    <b-dropdown class="col-1" id="ability-dropdown" text="特性">
      <b-dropdown-item v-for="ability in abilities" :key="ability" @click="change(ability)">{{ ability }}</b-dropdown-item>
    </b-dropdown>
    {{ current }}
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class AbilitySelector extends Vue {
  @Prop() private abilities!: string[];
  @Prop() private current!: string;

  change(ability: string) {
    this.$emit('changed', ability);
  }
}
</script>

<style scoped>
</style>
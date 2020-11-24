<template>
  <div class="ability-selector">
    <div class="row mb-1">
      <b-dropdown class="col-1" id="ability-dropdown" text="特性">
        <b-dropdown-item v-for="ability in abilities" :key="ability" @click="change(ability)">{{ ability }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{ current }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator';

@Component
export default class AbilitySelector extends Vue {
  @Prop()
  private abilities!: string[];
  @Prop()
  private current!: string;

  created() {
    if (this.current.length == 0) {
      this.$emit('change', this.abilities[0]);
    }
  }

  change(ability: string) {
    this.$emit('change', ability);
  }
}
</script>

<style scoped>
</style>
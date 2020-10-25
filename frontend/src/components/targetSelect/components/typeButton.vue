// タイプボタン
<template>
  <div>
    <b-button class="TypeButton" v-on:click="onButton">{{ type }}</b-button>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class TypeButton extends Vue {
  @Prop()
  private type!: string;

  onButton() {
    this.$emit('selected', this.type);
  }
}
</script>

<style scoped>
.TypeButton {
  width: 120px;
}
</style>

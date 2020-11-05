<template>
  <div class="speed-ranking">
    <ul class="list-group" v-for="info in ranking" :key="info.info">
      <template v-if="info.info == ''">
        <li class="list-group-item list-group-item-primary">
          {{ info.speed }}
        </li>
      </template>
      <template v-else>
        <li class="list-group-item">
          {{ info.speed }} {{ info.info }}
        </li>
      </template>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class SpeedRanking extends Vue {
  @Prop() ranking!: SpeedInfo[];
}
</script>

<style scoped>
</style>

<template>
  <div class="speed-ranking-item">
    <template v-if="info.target">
      <li class="list-group-item list-group-item-primary">
          {{ info.info }} {{ info.speed }}
      </li>
    </template>
    <template v-else>
      <li class="list-group-item">
          {{ info.info }} {{ info.speed }}
      </li>
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { SpeedInfo } from '../../store/speed/types'

@Component
export default class SpeedRankingItem extends Vue {
  @Prop()
  private info!: SpeedInfo;
 
}
</script>
<style scoped>
</style>

<template>
  <div class="speed-ranking-display">
    <ul class="list-group" v-for="info in infos" :key="info.index">
      <speed-ranking-item :info="info"></speed-ranking-item>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { SpeedInfo } from '../../store/speed/types'
import SpeedRankingItem from './SpeedRankingItem.vue'

@Component({
  components: {
    SpeedRankingItem,
  }
})
export default class SpeedRankingDisplay extends Vue {
  @Prop()
  private infos!: SpeedInfo[];
 
}
</script>
<style scoped>
</style>

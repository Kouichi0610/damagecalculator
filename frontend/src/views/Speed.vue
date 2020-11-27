<template>
  <div class="speed-adjuster">
    <target @ability="abilityChanged" @speed="speedChanged"></target>
    <speed-order :ability="ability" :targetSpeed="speed"></speed-order>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

import Target from "../components/target/target.vue";
import SpeedOrder from "../components/speedOrder/speedOrder.vue";

@Component({
  components: {
    Target,
    SpeedOrder,
  },
})
export default class Speed extends Vue {
  private speed: number = 0;
  private ability: string = ""

  speedChanged(speed: number) {
    this.speed = speed;
  }

  abilityChanged(ability: string) {
    this.ability = ability;
  }
}
</script>
<style scoped>
</style>

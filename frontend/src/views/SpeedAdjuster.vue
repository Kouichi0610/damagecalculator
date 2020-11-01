<template>
  <div class="speed-adjuster">
    TODO:Speed.vue をこっちに置き換え
    <target @ability="abilityChanged"></target>
    <speed-order></speed-order>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Target from '../components/target/target.vue'
import SpeedOrder from '../components/speedOrder/speedOrder.vue'

@Component({
      components: {
          Target,
          SpeedOrder,
    },
})
export default class SpeedAdjuster extends Vue {


  abilityChanged(ability: string) {
    console.log('' + ability);

  }
}
</script>
<style scoped>
</style>

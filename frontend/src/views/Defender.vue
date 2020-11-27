// 耐久調整
<template>
  <div class="defender">
    <adjust-target @sendDamage="changeSendDamage"></adjust-target>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from "vuex-class";
import AdjustTarget from '../components/adjustTarget/adjustTarget.vue'

import { ReceiveDamages } from '../store/defender/receiveDamages'

const namespace: string = "defenderState";

@Component({
  components: {
    AdjustTarget,
  }
})
export default class Defender extends Vue {
  // TODO:receiveDamages.
}
</script>

<style scoped>

</style>

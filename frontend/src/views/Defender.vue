// 耐久調整
<template>
  <div class="defender">
    <target :show="true"></target>
    <move-result></move-result>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Getter } from "vuex-class";

import Target from '../components/target/target.vue'
import MoveResult from '../components/defender/moveResult.vue'

import { AttackerDamages } from '../components/target/store/attackerDamages'

const namespace: string = "target";

@Component({
  components: {
    Target,
    MoveResult,
  }
})
export default class Defender extends Vue {
  @Getter("attackerDamages", { namespace })
  private attackerDamages!: AttackerDamages;
/*
  @Watch('attackerDamages', {deep: true})
  changed() {
    this.attackerDamages.attackerDamages();
  }
  */
}
</script>

<style scoped>

</style>

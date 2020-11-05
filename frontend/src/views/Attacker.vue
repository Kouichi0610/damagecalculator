// 攻撃側
<template>
  <div class="attacker">
    <target :show="true"></target>
    <div class="row mb-1">
      <moved-damages class="col-5"></moved-damages>
      <moved-damages class="col-5"></moved-damages>
    </div>
  </div>
</template>
/*
  { 技、仮想敵 } -> DamageResult

*/
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Target from '../components/target/target.vue'
import MovedDamages from '../components/movedDamages/movedDamages.vue'

@Component({
  components: {
    Target,
    MovedDamages,
  }
})
export default class Attacker extends Vue {
}
</script>

<style scoped>

</style>

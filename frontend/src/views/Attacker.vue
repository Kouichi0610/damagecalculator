// 攻撃側
<template>
  <div class="attacker">
    <target :show="true"></target>
  </div>
</template>
/*
  { 技、仮想敵 } -> DamageResult

*/
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Target from '../components/target/target.vue'

@Component({
  components: {
    Target,
  }
})
export default class Attacker extends Vue {
}
</script>

<style scoped>

</style>

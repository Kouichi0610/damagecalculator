// 攻撃側
<template>
  <div class="attacker">
    <target :show="true"></target>
    <move-result :index="index" :target="sample"></move-result>
    <div class="row mb-1">
    </div>
  </div>
</template>
/*
  { 技、仮想敵 } -> DamageResult

*/
<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
import Target from '../components/target/target.vue'
import MoveResult from '../components/attacker/moveResult.vue'

@Component({
  components: {
    Target,
    MoveResult,
  }
})
export default class Attacker extends Vue {
  @Prop() private index!: number;
}
</script>

<style scoped>

</style>

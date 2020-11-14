// 攻撃側
<template>
  <div class="attacker">
    Name:[{{ targetName }}]
    <target :show="true" @target="targetChanged"></target>
    <move-result :index="index" :target="targetName"></move-result>
    <div class="row mb-1">
    </div>
  </div>
</template>

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
  private targetName: string = '';

  targetChanged(name: string) {
    this.targetName = name;
  }
}
</script>

<style scoped>

</style>

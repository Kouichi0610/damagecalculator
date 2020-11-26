<template>
  <div class="result-list-display">
    <div v-for="result in results" :key="result.target">
      <result-display :result="result"></result-display>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { Result } from '../../store/attacker/sendDamage'
import ResultDisplay from './resultDisplay.vue'

@Component({
  components: {
    ResultDisplay,
  }
})
export default class ListDisplay extends Vue {
  @Prop()
  results!: Result[];
}
</script>
<style scoped>
</style>

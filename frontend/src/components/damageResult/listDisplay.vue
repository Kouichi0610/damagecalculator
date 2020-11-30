<template>
  <div class="container">
    <div v-for="result in results" :key="result.target">
      <result-display :result="result"></result-display>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator';
import { IDamageResult } from '../../store/damageResult/damageResult'
import ResultDisplay from './resultDisplay.vue'

@Component({
  components: {
    ResultDisplay,
  }
})
export default class ListDisplay extends Vue {
  @Prop()
  results!: IDamageResult[];
}
</script>
<style scoped>
</style>

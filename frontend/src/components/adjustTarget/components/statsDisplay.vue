<template>
  <div class="stats">
    <p>能力値</p>
    <div v-for="param in params" :key="param.key" class="row mb-1">
        <div class="col-sm-7 pt-1">
          <b-progress :max="400" show-value animated>
            <b-progress-bar :value="param.value" variant="info"></b-progress-bar>
          </b-progress>
        </div>
        <div class="col-1">{{ param.value }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';

class Params {
  readonly key: number;
  readonly value: number;
  constructor(key: number, value: number) {
    this.key = key;
    this.value = value;
  }
}

@Component
export default class StatsDisplay extends Vue {
  @Prop()
  private stats!: number[];

  get params(): Params[] {
    let res: Params[] = [];
    for (var i = 0; i < this.stats.length; i++) {
      res.push(new Params(i, this.stats[i]));
    }
    return res;
  }
}
</script>
<style scoped>
</style>

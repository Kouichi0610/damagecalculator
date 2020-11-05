<template>
  <div class="stats">
      <p>能力値</p>
      <div v-for="param in statsValues" :key="param.key" class="row mb-1">
          <div class="col-sm-3">{{ param.label }}{{ param.value }}</div>
          <div class="col-sm-8 pt-1">
            <b-progress :max="255" show-value animated>
              <b-progress-bar :value="param.value" variant="info"></b-progress-bar>
            </b-progress>
          </div>
      </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

interface statsValue {
  label: string;
  value: number;
}

@Component
export default class StatsDisplay extends Vue {
  @Prop()
  private hp!: number;
  @Prop()
  private attack!: number;
  @Prop()
  private defense!: number;
  @Prop()
  private spAttack!: number;
  @Prop()
  private spDefense!: number;
  @Prop()
  private speed!: number;

  private get statsValues(): statsValue[] {
    let res: statsValue[] = [];
    res.push({label: "HP", value: this.hp});
    res.push({label: "攻撃", value: this.attack});
    res.push({label: "防御", value: this.defense});
    res.push({label: "特攻", value: this.spAttack});
    res.push({label: "特御", value: this.spDefense});
    res.push({label: "素早さ", value: this.speed});
    return res;
  }

  
}
</script>

<style scoped>
</style>

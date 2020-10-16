// 種族値コンポーネント
<template>
    <div class="species">
        <p>種族値</p>
        <div v-for="param in speciesValues" :key="param.key" class="row mb-1">
            <div class="col-sm-1">{{ param.value }}</div>
            <div class="col-sm-6 pt-1">
              <b-progress :max="255" show-value animated>
                <b-progress-bar :value="param.value" variant="warning"></b-progress-bar>
              </b-progress>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

class SpeciesParam {
  public key: number;
  public value: number;

  constructor(key: number, value: number) {
    this.key = key;
    this.value = value;
  }
}

@Component
export default class Species extends Vue {
    @Prop()
    private params!: number[];

    created(): void {
    }

    // computed
    private get speciesValues(): SpeciesParam[] {
      let res: SpeciesParam[] = [
        new SpeciesParam(0, this.params[0]),
        new SpeciesParam(1, this.params[1]),
        new SpeciesParam(2, this.params[2]),
        new SpeciesParam(3, this.params[3]),
        new SpeciesParam(4, this.params[4]),
        new SpeciesParam(5, this.params[5]),
      ];
      return res;
    }
}
</script>

<style scoped>
</style>

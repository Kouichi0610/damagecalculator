// 種族値コンポーネント
<template>
    <div class="species">
        <p>種族値</p>
        <div v-for="param in values" :key="param.key" class="row mb-1">
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
import { Species } from '../store/species'

class SpeciesParam {
  public key: number;
  public value: number;

  constructor(key: number, value: number) {
    this.key = key;
    this.value = value;
  }
}

@Component
export default class SpeciesDisplay extends Vue {
    @Prop()
    private species!: Species;

    created(): void {
    }

    // computed
    private get values(): SpeciesParam[] {
      let res: SpeciesParam[] = [
        new SpeciesParam(0, this.species.hp()),
        new SpeciesParam(1, this.species.attack()),
        new SpeciesParam(2, this.species.defense()),
        new SpeciesParam(3, this.species.spAttack()),
        new SpeciesParam(4, this.species.spDefense()),
        new SpeciesParam(5, this.species.speed()),
      ];
      return res;
    }
}
</script>

<style scoped>
</style>

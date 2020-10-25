/*
  対象ポケモンを選択する
  TODO:components/targetSelect.vueと関連クラス削除

  TODO:構成変えてみる
  stats
   +---components
   +---store
   +---logic
*/
<template>
  <div class="targetSelect">
    <TypeButtons :initialType="initialType" @selected="onSelectType"></TypeButtons>
    <species-total :initialTotal="initialTotal" @changed="onChangeTotal"></species-total>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-property-decorator';
import { State, Getter, Mutation } from 'vuex-class';
import Component from 'vue-class-component';

import TypeButtons from './components/typeButtons.vue'
import SpeciesTotal from './components/speciesTotal.vue'

const namespace: string = 'targetSelect';

@Component({
  components: {
    TypeButtons,
    SpeciesTotal,
  }
})
export default class TaretSelect extends Vue {
  @State('targetSelect') targetSelect: TargetSelectState;
  @Getter('initialTotal', { namespace })
  private initialTotal!: number;
  @Getter('initialType', { namespace })
  private initialType!: string;

  @Mutation('setInitialTotal', { namespace })
  private setInitialTotal!: (number) => void;
  @Mutation('setInitialType', { namespace })
  private setInitialType!: (string) => void;

  // TODO:candidates
  created() {
  }

  onSelectType(typeName: string) {
    console.log('st');
    this.setInitialType(typeName);
  }
  onChangeTotal(total: number) {
    console.log('stt');
    this.setInitialTotal(total);
  }
}
</script>

<style scoped>
</style>

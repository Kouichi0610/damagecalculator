/*
  対象ポケモンを選択する
  TODO:components/targetSelect.vueと関連クラス削除
  stats
   +---components
   +---store
   +---logic
*/
<template>
  <div class="targetSelect">
    <TypeButtons :initialType="initialType" @selected="onSelectType"></TypeButtons>
    <species-total :initialTotal="initialTotal" @changed="onChangeTotal"></species-total>
    <candidates :candidates="candidates" @selected="onSelect"></candidates>
  </div>
</template>

<script lang="ts">
import { Vue, Watch } from 'vue-property-decorator';
import { State, Getter, Mutation, Action } from 'vuex-class';
import Component from 'vue-class-component';

import TypeButtons from './components/typeButtons.vue'
import SpeciesTotal from './components/speciesTotal.vue'
import Candidates from './components/candidates.vue'
import { CandidatesFilter, Species } from '../../store/targetSelect/types'

const namespace: string = 'targetSelectState';

@Component({
  components: {
    TypeButtons,
    SpeciesTotal,
    Candidates,
  }
})
export default class TaretSelect extends Vue {
  @Getter('initialTotal', { namespace })
  private initialTotal!: number;
  @Getter('initialType', { namespace })
  private initialType!: string;
  @Getter('candidates', { namespace })
  private candidates!: Species[];

  @Action('getCandidates', { namespace })
  private getCandidates!: (filter: CandidatesFilter) => Promise<boolean>;

  @Mutation('setInitialTotal', { namespace })
  private setInitialTotal!: (total: number) => void;
  @Mutation('setInitialType', { namespace })
  private setInitialType!: (type: string) => void;

  created() {
  }

  @Watch('initialTotal')
  totalChanged() {
    this.getCandidates(new CandidatesFilter(this.initialType, this.initialTotal));
  }
  @Watch('initialType')
  typeChanged() {
    this.getCandidates(new CandidatesFilter(this.initialType, this.initialTotal));
  }

  @Watch('candidates')
  candidatesChanged() {
  }

  onSelectType(typeName: string) {
    this.setInitialType(typeName);
  }
  onChangeTotal(total: number) {
    this.setInitialTotal(total);
  }
  onSelect(name: string) {
    this.$emit('selectPokemon', name);
  }
}
</script>

<style scoped>
</style>

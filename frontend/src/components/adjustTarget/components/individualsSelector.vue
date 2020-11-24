<template>
  <div class="individuals-selector">
    <p>個体値 {{ individuals.toString() }}</p>
    <b-form-checkbox id="check-slowest" v-model="isSlowest" @change="changeSlowest" name="check-slowest">最遅</b-form-checkbox>
    <b-form-checkbox id="check-weakest" v-model="isWeakest" @change="changeWeakest" name="check-weakest">最弱</b-form-checkbox>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { Individuals } from '../../../store/individuals/types'

const namespace: string = "individualsState";

@Component
export default class IndividualsSelector extends Vue {
  @Prop() private target!: string;

  @Getter('individuals', { namespace })
  private individuals!: Individuals;
  @Mutation('reset', { namespace })
  private reset!: () => void;
  @Mutation('changeSlowest', { namespace })
  private changeSlowest!: () => void;
  @Mutation('changeWeakest', { namespace })
  private changeWeakest!: () => void;

  get isSlowest(): boolean {
    return this.individuals.isSlowest();
  }
  get isWeakest(): boolean {
    return this.individuals.isWeakest();
  }

  created() {
    this.individualsChanged();
  }

  @Watch('target')
  clear() {
    this.reset();
  }

  @Watch('individuals', { deep: true })
  individualsChanged() {
    this.$emit('change', this.individuals);
  }
}
</script>

<style scoped>
</style>

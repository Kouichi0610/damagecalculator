<template>
  <div class="nature-selector">
    <div class="row mb-1">
      <b-dropdown class="col-md-2" id="nature-dropdown" text="性格">
        <b-dropdown-item v-for="item in natures" :key="item.name" @click="setCurrent(item)">{{ item.name }} {{ item.description }}</b-dropdown-item>
      </b-dropdown>
      <div class="col-md-4">
        性格:{{ current.name }} {{ current.description }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
/*
  性格設定
*/
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { Nature } from '../../../store/nature/types'

const namespace: string = "natureState";

@Component
export default class NatureSelector extends Vue {
  @Prop() private target!: string;

  @Action('loadNatures', { namespace })
  private loadNatures!: () => void;
  @Action('initialize', { namespace })
  private initialize!: (target: string) => void;

  @Getter('isLoaded', { namespace })
  private isLoaded!: boolean;
  @Getter('natures', { namespace })
  private natures!: Nature[];
  @Getter('current', { namespace })
  private current!: Nature;
  @Mutation('setCurrent', { namespace })
  private setCurrent!: (current: Nature) => void;
  @Mutation('clearCurrent', { namespace })
  private clearCurrent!: () => void;

  created() {
    this.initialize(this.target);
    if (this.isLoaded) {
      this.currentChanged();
      return;
    }
    this.loadNatures();
  }

  @Watch('target')
  reset() {
    this.clearCurrent();
  }

  @Watch('current', { deep: true })
  currentChanged() {
    this.$emit('change', this.current);
  }

}
</script>

<style scoped>
</style>

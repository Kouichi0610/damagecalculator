// 性格一覧、選択
<template>
  <div class="nature">
    <div class="row mb-1">
      <div class="col-1"></div>
      <b-dropdown class="col-1" id="nature-dropdown" text="性格">
        <b-dropdown-item v-for="item in natures" :key="item.name" @click="change(item.name)">{{ item.name }} {{item.description}}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{current}}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Watch } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from 'vuex-class';
//import { NatureState } from './store/types'

const namespace: string = "nature";

@Component
export default class Nature extends Vue {
  //@State('nature') nature!: NatureState;
  @Action('initialize', { namespace })
  private initialize!: () => Promise<boolean>;

@Getter('natures', { namespace })
  private natures!: Nature[];
  @Getter('current', { namespace })
  private current!: string;
  @Mutation('change', { namespace })
  private change!: (natureName: string) => void;

  created() {
    if (this.natures.length > 0) {
      return;
    }
    this.initialize();
  }

  @Watch('current')
  private currentChanged() {
    this.$emit('changed', this.current);
  }
}
</script>

<style scoped>
</style>

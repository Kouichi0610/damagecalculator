<template>
  <div class="basepoints-adjuster">
    基礎ポイント
    {{ basePointString }}
    TOTAL: {{ total }}
    <div v-for="item in basePointsArray" :key="item.index">
      <adjuster :basePoint="item" @change="change"></adjuster>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { IBasePoint, BasePoints } from '../../../store/basePoints/types';
import Adjuster from './basePoints/adjuster.vue';

const namespace: string = "basePointsState";

@Component({
  components: {
    Adjuster,
  }
})
export default class BasePointsAdjuster extends Vue {
  //@Getter('basePoints', { namespace })
  private basePoints: BasePoints = new BasePoints();

  @Mutation('reset', { namespace })
  private reset!: () => void;

  get total(): number {
    return this.basePoints.total();
  }

  get basePointString(): string {
    return this.basePoints.toString();
  }

  @Watch('basePointString')
  changed() {
    console.log('Changed?');
  }

  get basePointsArray(): IBasePoint[] {
    return this.basePoints.array();
  }

  @Watch('basePointsArray')
  basePointsArrayChanged() {
    console.log('array changed.');
  }

  @Watch('basePoints', { deep: true })
  basePointsChanged() {
    console.log('BSChangeed:' + this.basePoints.toString());
  }
  

  created() {
  }

  change(index: number, value: number) {
    console.log('Change:' + index + ' -> ' + value);
    this.basePoints.set(index, value);
    console.log('changed.' + this.basePoints.toString());
  }
}
</script>

<style scoped>
</style>

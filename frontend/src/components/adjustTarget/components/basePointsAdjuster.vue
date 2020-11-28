<template>
  <div class="basepoints-adjuster">
    基礎ポイント
    <div v-for="item in basePointsArray" :key="item.index">
      <adjuster :basePoint="item" @change="change" :enable="enables[item.index]"></adjuster>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { IBasePoint, IBasePoints } from '../../../store/basePoints/types';
import Adjuster from './basePoints/adjuster.vue';

const namespace: string = "basePointsState";

@Component({
  components: {
    Adjuster,
  }
})
export default class BasePointsAdjuster extends Vue {
  @Prop() private target!: string;
  @Prop() private speedLock!: boolean

  @Action('initialize', { namespace })
  private initialize!: (target: string) => void;
  @Getter('basePoints', { namespace })
  private basePoints!: IBasePoints;
  @Getter('basePointsArray', { namespace })
  private basePointsArray!: IBasePoint[];

  @Mutation('setHP', { namespace })
  private setHP!: (value: number) => void;
  @Mutation('setAttack', { namespace })
  private setAttack!: (value: number) => void;
  @Mutation('setDefense', { namespace })
  private setDefense!: (value: number) => void;
  @Mutation('setSpAttack', { namespace })
  private setSpAttack!: (value: number) => void;
  @Mutation('setSpDefense', { namespace })
  private setSpDefense!: (value: number) => void;
  @Mutation('setSpeed', { namespace })
  private setSpeed!: (value: number) => void;

  get enables(): boolean[] {
    return [
      !this.speedLock,
      !this.speedLock,
      !this.speedLock,
      !this.speedLock,
      !this.speedLock,
      true,
    ];
  }

  private setters: ((value: number) => void)[] = [
    this.setHP,
    this.setAttack,
    this.setDefense,
    this.setSpAttack,
    this.setSpDefense,
    this.setSpeed,
  ];

  created() {
    this.initialize(this.target);
  }

  change(index: number, value: number) {
    this.setters[index](value);
  }

  @Watch('basePoints', { deep: true })
  basePointsChanged() {
    this.$emit('change', this.basePoints);
  }
}
</script>

<style scoped>
</style>

<template>
  <div class="moves">
    わざ一覧
    <div class="row mb-1">
      <b-dropdown  class="col-1" id="physicals-dropdown" text="物理">
        <b-dropdown-item v-for="item in physicals" :key="item.name" @click="change(item)">{{ item.name }}</b-dropdown-item>
      </b-dropdown>
      <b-dropdown class="col-1" id="specials-dropdown" text="特殊">
        <b-dropdown-item v-for="item in specials" :key="item.name" @click="change(item)">{{ item.name }}</b-dropdown-item>
      </b-dropdown>
      <!-- 
      <div class="col-1">
        <label><input type="radio" v-model="isPhysical" value="true">物理</label>
        <label><input type="radio" v-model="isPhysical" value="false">特殊</label>
        -->
      </div>
      <div class="col-8">
        {{ selected.name }} {{ selected.power }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
import { Action, Getter } from "vuex-class";
import { MoveInfo } from './store/types'

const namespace: string = "moves"

// わざ一覧を取得
@Component
export default class Moves extends Vue {
  @Prop() private target!: string;
  @Action("getMoves", { namespace }) getMoves!: (string) => Promise<string>;
  @Getter("physicals", { namespace })
  private physicals!: MoveInfo[];
  @Getter("specials", { namespace })
  private specials!: MoveInfo[];

  private isPhysical: boolean = true;
  private selected: MoveInfo = MoveInfo.empty();

  @Watch('target')
  targetChanged() {
    this.getMoves(this.target)
  }

  change(info: MoveInfo) {
    this.selected = info;
  }

  
}
</script>
<style scoped>
</style>

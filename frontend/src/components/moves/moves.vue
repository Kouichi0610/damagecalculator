<template>
  <div class="moves">
    わざ一覧
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
import { Action /*State, Action, Getter, Mutation*/ } from "vuex-class";

const namespace: string = "moves"

// わざ一覧を取得
@Component
export default class Moves extends Vue {
  @Prop() private target!: string;

  @Action("getMoves", { namespace }) getMoves!: (string) => Promise<string>;

  // TODO:async
  @Watch('target')
  targetChanged() {
    this.getMoves(this.target)
  }
  
}
</script>
<style scoped>
</style>

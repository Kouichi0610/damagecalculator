/*
  対象となるポケモンの情報を表示する
  種族値
  とくせい
  わざ
*/
<template>
  <div class="target">
    <p>Name:{{name}}</p>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from 'vue-property-decorator';
import { State, Action } from 'vuex-class';

const namespace: string = "target";

@Component({
  components: {
  }
})
export default class Target extends Vue {
  @State('target') target: TargetState;
  @Action('getSpecies', { namespace })
  private getSpecies!: (string) => Promise<boolean>;

  @Prop() private name: string;

  @Watch('name')
  private nameChanged() {
    this.getSpecies(this.name);
  }


}
</script>

<style scoped>
</style>

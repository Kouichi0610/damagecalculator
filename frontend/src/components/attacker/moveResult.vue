<template>
  <div class="move-result">
    MoveResult:{{target}} Index:{{index}}
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Watch, Component } from "vue-property-decorator";
//import { State, Action, Getter, Mutation } from "vuex-class";

@Component({
  components:{
  },
})
export default class MoveResult extends Vue {
  @Prop() private target!: string;
  @Prop() private index!:number;

  @Watch("target")
  watchChanged() {
    console.log('watchChanged:' + this.target);
  }
  created() {
    console.log('index:' + this.index + ' name:' + this.target)
  }
  
}
</script>

<style scoped>
</style>

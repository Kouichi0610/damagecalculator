<template>
  <div class="move-result">
    MoveResult:{{target}} Index:{{index}}
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
//import { State, Action, Getter, Mutation } from "vuex-class";

@Component({
  components:{
  },
})
export default class MoveResult extends Vue {
  @Prop() private target: string;
  @Prop() private index:number;
  
}
</script>

<style scoped>
</style>

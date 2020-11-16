<template>
  <div class="move-result">
    Species:{{target.name}}  Index:{{index}}
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
//import { State, Action, Getter, Mutation } from "vuex-class";

import { Species } from '../target/store/species'

@Component({
  components:{
  },
})
export default class MoveResult extends Vue {
  @Prop() private index!: number;
  @Prop() private target!: Species;

  mounted() {
    console.log('' + this.target.name);
  }

  
}
</script>

<style scoped>
</style>

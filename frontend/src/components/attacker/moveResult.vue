<template>
  <div class="move-result">
    Index:{{index}}
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
//import { State, Action, Getter, Mutation } from "vuex-class";

import { Species } from '../target/store/species'
import { DefenderDamages } from '../target/store/defenderDamages'


@Component({
  components:{
  },
})
export default class MoveResult extends Vue {
  @Prop() private index!: number;
  @Prop() private damages!: DefenderDamages;

  mounted() {
  }

  
}
</script>

<style scoped>
</style>

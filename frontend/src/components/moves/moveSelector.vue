<template>
  <div class="container">
    <div class="row">
      <b-dropdown class="col-md-2" id="physicals-dropdown" text="物理">
        <b-dropdown-item v-for="item in physicals" :key="item.name" @click="select(item)">{{ item.name }}</b-dropdown-item>
      </b-dropdown>
      <b-dropdown class="col-md-2" id="specials-dropdown" text="特殊">
        <b-dropdown-item v-for="item in specials" :key="item.name" @click="select(item)">{{ item.name }}</b-dropdown-item>
      </b-dropdown>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { Move } from '../../store/attacker/types'

/*
  わざ選択
*/
@Component
export default class MoveSelector extends Vue {
  @Prop()
  physicals!: Move[];
  @Prop()
  specials!: Move[];

  select(move: Move) {
    this.$emit('select', move);
  }
}
</script>

<style scoped>
</style>

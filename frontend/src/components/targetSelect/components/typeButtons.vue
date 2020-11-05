<template>
  <div class="typeButtons">
    <div class="btn-group-vertical">
      <b-button-group>
        <type-button type="すべて" @selected="onSelectType"></type-button>
      </b-button-group>
      <b-button-group>
        <type-button type="ノーマル" @selected="onSelectType"></type-button>
        <type-button type="ほのお" @selected="onSelectType"></type-button>
        <type-button type="みず" @selected="onSelectType"></type-button>
        <type-button type="でんき" @selected="onSelectType"></type-button>
        <type-button type="くさ" @selected="onSelectType"></type-button>
        <type-button type="こおり" @selected="onSelectType"></type-button>
      </b-button-group>
      <b-button-group>
        <type-button type="かくとう" @selected="onSelectType"></type-button>
        <type-button type="どく" @selected="onSelectType"></type-button>
        <type-button type="じめん" @selected="onSelectType"></type-button>
        <type-button type="ひこう" @selected="onSelectType"></type-button>
        <type-button type="エスパー" @selected="onSelectType"></type-button>
        <type-button type="むし" @selected="onSelectType"></type-button>
      </b-button-group>
      <b-button-group>
        <type-button type="いわ" @selected="onSelectType"></type-button>
        <type-button type="ゴースト" @selected="onSelectType"></type-button>
        <type-button type="ドラゴン" @selected="onSelectType"></type-button>
        <type-button type="あく" @selected="onSelectType"></type-button>
        <type-button type="はがね" @selected="onSelectType"></type-button>
        <type-button type="フェアリー" @selected="onSelectType"></type-button>
      </b-button-group>
      <species-total :initialTotal="500" @changed="onChangeTotal"></species-total>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

import TypeButton from './typeButton.vue'

@Component({
  components: {
    TypeButton,
  }
})
export default class TypeButtons extends Vue {
  @Prop()
  private initialType!: string;

  onSelectType(type: string) {
    this.$emit('selected', type);
  }
}
</script>
<style scoped>
</style>

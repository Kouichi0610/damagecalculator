// 性格一覧、選択
<template>
  <div class="nature">
    <div class="row mb-1">
      <div class="col-1"></div>
      <b-dropdown class="col-1" id="nature-dropdown" text="性格">
        <b-dropdown-item v-for="item in natures" :key="item.name" @click="current = item.name">{{ item.name }} {{item.description}}</b-dropdown-item>
      </b-dropdown>
      <div class="col-1">{{current}}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { State, Action, Getter /* Mutation*/ } from 'vuex-class';

const namespace: string = "nature";

@Component
export default class Nature extends Vue {
  @State('nature') nature: NatureState;
  @Action('initialize', { namespace })
  private initialize!: () => Promise<boolean>;
  @Getter('initialized', { namespace })
  private initialized!: () => boolean;
  @Getter('natures', { namespace })
  private natures!: () => Nature[];

  private current: string = '';

  created() {
    if (this.initialized) {
      return;
    }
    this.initialize();
  }
}
</script>

<style scoped>
</style>

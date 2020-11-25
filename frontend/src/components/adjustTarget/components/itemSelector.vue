<template>
  <div class="item-selector">
    <div class="row mb-1">
      <div class="col-3">
        <b-dropdown class="col-1" id="items-dropdown" text="もちもの">
          <b-dropdown-item v-for="item in items" :key="item.name" @click="setCurrent(item)">{{ item.name }} {{ item.description }}</b-dropdown-item>
        </b-dropdown>
      </div>
      <div class="col-3">
        {{ current.name }}
      </div>
      <div class="col-3">
        {{ current.description }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Watch, Prop } from 'vue-property-decorator';
import { State, Action, Getter, Mutation } from "vuex-class";
import { Item } from '../../../store/items/types'

const namespace: string = "itemState";

@Component
export default class ItemSelector extends Vue {
  @Prop()
  private target!: string;

  @Action('initialize', { namespace })
  private initialize!: (name: string) => void;
  @Getter('items', { namespace })
  private items!: Item[];
  @Getter('current', { namespace })
  private current!: Item;
  @Mutation('setCurrent', { namespace })
  private setCurrent!: (current: Item) => void;

  created() {
    this.initialize(this.target);
  }

  @Watch('current')
  itemChanged() {
    this.$emit('change', this.current);
  }

}
</script>

<style scoped>
</style>

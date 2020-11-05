<template>
  <div class="candidates">
    <template v-if="candidates.length > 0">
      <ul class="list-group" v-for="candidate in candidates" :key="candidate.name">
        <candidate :candidate="candidate" @selected="onSelect"></candidate>
      </ul>
    </template>
    <template v-else>
      no candidates.
    </template>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';
//import { Species } from '../store/types'
import Candidate from './candidate.vue'

@Component({
  components: {
    Candidate,
  }
})
export default class Candidates extends Vue {
  @Prop()
  private candidates!: Species[];

  onSelect(name: string) {
    this.$emit('selected', name);
  }
}
</script>

<style scoped>
.candidates {
  width:600px;
  height:600px;
  overflow: scroll;
}
</style>

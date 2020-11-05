// 耐久調整
<template>
  <div class="defender">
    <target :show="true"></target>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Target from '../components/target/target.vue'

@Component({
  components: {
    Target,
  }
})
export default class Defender extends Vue {
}
</script>

<style scoped>

</style>

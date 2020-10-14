<template>
    <div class="sandboxts">
        SandBox from TypeScript.
        <StatsEditor :species="species"></StatsEditor>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import StatsEditor from '../components/stats/statsEditor.vue'

@Component({
      components: {
          StatsEditor,
    },
})
export default class SandBoxTs extends Vue {
    species: number[] = [100, 25, 180, 35, 120, 60];
}
</script>

<style scoped>
</style>

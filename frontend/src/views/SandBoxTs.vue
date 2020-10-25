<template>
    <div class="sandboxts">
        <TargetSelect></TargetSelect>
        SandBox from TypeScript.(1.1)
        <StatsEditor></StatsEditor>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import StatsEditor from '../components/stats/statsEditor.vue'
import TargetSelect from '../components/targetSelect/targetSelect.vue'

@Component({
      components: {
          StatsEditor,
          TargetSelect,
    },
})
export default class SandBoxTs extends Vue {
    species: number[] = [100, 25, 180, 35, 120, 60];
}
</script>

<style scoped>
</style>

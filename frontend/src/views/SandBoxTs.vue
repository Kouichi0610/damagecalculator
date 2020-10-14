<template>
    <div class="sandboxts">
        SandBox from TypeScript.
        <Nature selected="むじゃき" @changed="changeNature"></Nature>
        <BasePoints></BasePoints>
        <Species :params="species"></Species>
        <Stats :params="stats"></Stats>
        <individuals @changed="changeIndividuals"></individuals>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import Nature from '../components/stats/components/nature.vue'
import BasePoints from '../components/stats/components/basePoints.vue'
import Individuals from '../components/stats/components/individuals.vue'
import Species from '../components/stats/components/species.vue'
import Stats from '../components/stats/components/stats.vue'

@Component({
      components: {
        Nature,
        BasePoints,
        Individuals,
        Species,
        Stats,
    },
})
export default class SandBoxTs extends Vue {
    stats: number[] = [300, 120, 60, 188, 25, 33];
    species: number[] = [100, 25, 180, 35, 120, 60];
    // interfaceの$emitも問題なし
    private changeNature(n: INature): void {
        console.log('changeNature:' + n.Name);

    }
    private changeIndividuals(params: number[]) {
        console.log('changeIndividuals:' + params);
    }
}
</script>

<style scoped>
</style>

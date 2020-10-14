// 
<template>
    <div class="statsEditor">
        <nature :selected="nature" @changed="updateNature"></nature>
        <individuals :params="individuals" @changed="updateIndividuals"></individuals>
        <div class="row mb-1">
            <stats class="col-4" :params="stats"></stats>
            <species class="col-2" :params="species"></species>
            <base-points class="col-2" :params="basePoints" @changed="updateBasePoints"></base-points>
        </div>
    </div>
</template>

/*
*/

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

import Nature from './components/nature.vue'
import BasePoints from './components/basePoints.vue'
import Individuals from './components/individuals.vue'
import Species from './components/species.vue'
import Stats from './components/stats.vue'

import * as nature from '../../domain/nature'
import * as stats from '../../domain/stats'

// TODO: $emit stats
@Component({
      components: {
        Nature,
        BasePoints,
        Individuals,
        Species,
        Stats,
    },
})
export default class StatsEditor extends Vue {
    private individuals: number[] = [31,31,31,31,31,31];
    private stats: number[] = [0,0,0,0,0,0];
    private basePoints: number[] = [0,0,0,0,0,0];
    private nature: nature.INature = nature.NatureFactory('てれや');
    private calculator: stats.StatsCalculator = new stats.StatsCalculator();

    created() {
        this.calcStats();
    }

    @Prop()
    private species!: number[];

    private get natureName(): string {
        return this.nature.Name();
    }

    private updateNature(nature: nature.INature) {
        this.nature = nature;
        console.log('next:' + this.nature.Name());
        this.calcStats();
    }
    private updateIndividuals(params: number[]) {
        this.individuals = params;
        console.log('individuals:' + this.individuals);
        this.calcStats();
    }
    private updateBasePoints(params: number[]) {
        this.basePoints = params;
        console.log('basePoints:' + this.basePoints);
        this.calcStats();
    }

    private calcStats(): number[] {
        this.stats = this.calculator.Calculate(this.nature, 50, this.species, this.individuals, this.basePoints);
        console.log('Stats:' + this.stats);
        return this.stats;
    }
}
</script>

<style>

</style>
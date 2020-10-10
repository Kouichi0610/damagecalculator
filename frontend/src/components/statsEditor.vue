/*
    能力値編集コンポーネント
    ・能力値
    ・種族値
    ・性格
    ・個体値
    ・基礎ポイント
*/
<template>
    <div class="statsEditor">
        <stats :params="status"></stats>
        <species :params="species"></species>
        <nature :current="nature" :natures="natures" @update="updateNature"></nature>
        <individuals :params="individuals" @update="updateIndividuals"></individuals>
        <base-points :params="basepoints" @update="updateBasePoints"></base-points>
    </div>
</template>

<script>
import Nature from './statsComponents/nature'
import Species from './statsComponents/species'
import Individuals from './statsComponents/individuals'
import BasePoints from './statsComponents/basePoints'
import Stats from './statsComponents/stats'

import {natureNames, getNature} from '../domain/nature'
import {calcStats} from '../domain/stats'

export default {
    name: 'statsEditor',
    components: {
        'stats': Stats,
        'nature': Nature,
        'species': Species,
        'individuals': Individuals,
        'base-points': BasePoints,
    },
    // props species
    data: function() {
        return {
            status: [0,0,0,0,0,0],
            species: [95, 109, 105, 75, 85, 56],       // 種族値
            individuals: [31, 31, 31, 31, 31, 31],  // 個体値
            basepoints: [100, 0, 0, 0, 0, 0],   // 基礎ポイント
            nature: 'むじゃき',
            natures: [],
        }
    },
    created: function() {
        this.natures = natureNames();

        let n = getNature('ようき');
        console.log('result' + n.name
         + ' At:' + n.attack(100)
         + ' Df:' + n.defense(200)
         + ' SA:' + n.spAttack(300)
         + ' SD:' + n.spDefense(400)
         + ' SP:' + n.speed(500));
    },
    computed: {
    },
    methods: {
        updateIndividuals :function(slowest, weakest) {
            this.individuals[1] = weakest ? 0 : 31;
            this.individuals[5] = slowest ? 0 : 31;
            this.updateStatus();
        },
        updateNature: function(name) {
            this.nature = name;
            this.updateStatus();
        },
        updateBasePoints: function(values) {
            this.basepoints = values;
            this.updateStatus();
        },
        updateStatus: function() {
            let nature = getNature(this.nature);
            // TODO:level
            this.status = calcStats(50, this.species, this.individuals, this.basepoints, nature);
        }

    }
}
</script>

<style scoped>
</style>
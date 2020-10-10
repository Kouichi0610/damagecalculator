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
        <stats :params="stats"></stats>
        <species :params="species"></species>
        <nature :current="nature" :natures="natures"></nature>
        <individuals :params="individuals" @update="updateIndividuals"></individuals>
        <base-points :params="basepoints"></base-points>
    </div>
</template>

<script>
import Nature from './statsComponents/nature'
import Species from './statsComponents/species'
import Individuals from './statsComponents/individuals'
import BasePoints from './statsComponents/basePoints'
import Stats from './statsComponents/stats'

import {natureNames, getNature} from '../domain/nature'

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
            stats: [100, 200, 300, 400, 500, 600],    // 能力値
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
    methods: {
        updateIndividuals :function(slowest, weakest) {
            this.individuals[1] = weakest ? 0 : 31;
            this.individuals[5] = slowest ? 0 : 31;
        }
    }
}
</script>

<style scoped>
</style>
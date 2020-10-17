// 対象ポケモンの選択
<template>
    <div class="targetSelect">
        <p>ポケモン選択</p>
        <input @click="getSample" type="button" value="試作">
    </div>
</template>

<script lang="ts">
//import { Vue } from 'vue-property-decorator';
import Vue from 'vue';
import { State, /*State, Getter, */ Action } from 'vuex-class';
import Component from 'vue-class-component';

//import {Species, TargetsState} from '../store/targets/types'

const namespace: string = 'targets';


@Component
export default class TargetSelect extends Vue {
    @State('targets') targets: TargetsState;
    @Action('getCandidates', { namespace })
    private getCandidates!: (Filter) => Promise<boolean>;

    getSample(): void {
        console.log('todo:');
        this.getCandidates({
            types: 'ほのお',
            total: 530,
        });
    }
    
}
</script>

<style scoped>
</style>

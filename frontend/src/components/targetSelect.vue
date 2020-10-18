// 対象ポケモンの選択
<template>
    <div class="targetSelect">
        <p>ポケモン選択 {{ candidatesCount }}</p>
        <div class="btn-group-vertical">
            <b-button-group>
                <types-button type="すべて" @clicked="setType"></types-button>
            </b-button-group>
            <b-button-group>
                <types-button type="ノーマル" @clicked="setType"></types-button>
                <types-button type="ほのお" @clicked="setType"></types-button>
                <types-button type="みず" @clicked="setType"></types-button>
                <types-button type="でんき" @clicked="setType"></types-button>
                <types-button type="くさ" @clicked="setType"></types-button>
                <types-button type="こおり" @clicked="setType"></types-button>
            </b-button-group>
            <b-button-group>
                <types-button type="かくとう" @clicked="setType"></types-button>
                <types-button type="どく" @clicked="setType"></types-button>
                <types-button type="じめん" @clicked="setType"></types-button>
                <types-button type="ひこう" @clicked="setType"></types-button>
                <types-button type="エスパー" @clicked="setType"></types-button>
                <types-button type="むし" @clicked="setType"></types-button>
            </b-button-group>
            <b-button-group>
                <types-button type="いわ" @clicked="setType"></types-button>
                <types-button type="ゴースト" @clicked="setType"></types-button>
                <types-button type="ドラゴン" @clicked="setType"></types-button>
                <types-button type="あく" @clicked="setType"></types-button>
                <types-button type="はがね" @clicked="setType"></types-button>
                <types-button type="フェアリー" @clicked="setType"></types-button>
            </b-button-group>
        </div>

        <div class="input-group mb-3">
        <p>種族値合計</p>
        <div class="input-group-prepend">
            <div class="input-group-text">
            <input type="checkbox" v-model="enableTotal" aria-label="Checkbox for following text input">
            </div>
        </div>
        <input type="number" v-model="total" class="form-control" aria-label="Text input with checkbox">
        <p>以上</p>
        </div>

    </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { State, Getter, Action } from 'vuex-class';
import Component from 'vue-class-component';
import { Watch } from 'vue-property-decorator';
import TypesButton from './types/typesButton.vue'
//import {Species, TargetsState} from '../store/targets/types'
const namespace: string = 'targets';

@Component({
    components: {
        TypesButton,
    }
})
export default class TargetSelect extends Vue {
    @State('targets') targets: TargetsState;
    @Action('getCandidates', { namespace })
    private getCandidates!: (Filter) => Promise<boolean>;
    @Getter('candidatesCount', { namespace })
    private candidatesCount: number;

    private types: string = 'すべて';
    private enableTotal: boolean = true;
    private total: number = 480;

    @Watch('total')
    totalChanged(after: number, before: number) {
        if (before == after) {
            return;
        }
        this.updateCandidates();
    }
    @Watch('enableTotal')
    enableTotalChanged() {
        this.updateCandidates();
    }
    @Watch('types')
    typesChanged(after: string, before: string) {
        if (after == before) {
            return;
        }
        this.updateCandidates();
    }

    setType(type: string) {
        this.types = type;
    }

    updateCandidates() {
        console.log('update. types:' + this.types + " total:" + this.total + " enableTotal:" + this.enableTotal);
        let total = this.enableTotal ? this.total : 0;
        this.getCandidates({
            types: this.types,
            total: total,
        });
    }
}
</script>

<style scoped>
.targetSelect {
    width: 480px;
}
</style>

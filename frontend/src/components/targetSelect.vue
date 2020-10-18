// 対象ポケモンの選択
<template>
    <div class="targetSelect">
        <p>ポケモン選択 {{types}} {{total}}</p>
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
        <input type="number" v-model="total" class="form-control" aria-label="Text input with checkbox">
        <p>以上</p>
        </div>

        <!-- 候補一覧 -->
        <div class="candidates">
            候補 {{ candidatesCount }}
            <ul class="list-group" v-for="candidate in candidates" :key="candidate.name">
                <li class="list-group-item">
                    <b-button class="TypeButton" v-on:click="onButton(candidate)">{{ candidate.name }}</b-button>
                    {{ candidate.types }} HP:{{candidate.hp}} AT:{{candidate.attack}} DF:{{candidate.defense}} SA:{{candidate.spAttack}} SD:{{candidate.spDefense}} SP:{{candidate.speed}}
                </li>
            </ul>
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
//import router from '../router'

const namespace: string = 'targets';

// TODO:選択したポケモンを別のstoreに(あるいはルートに)渡す

@Component({
    components: {
        TypesButton,
    }
})
export default class TargetSelect extends Vue {
    @State('targets') targets: TargetsState;
    // 候補一覧を取得
    @Action('getCandidates', { namespace })
    private getCandidates!: (Filter) => Promise<boolean>;
    // 候補者数
    @Getter('candidatesCount', { namespace })
    private candidatesCount: number;
    // 候補一覧
    @Getter('candidates', { namespace })
    private candidates: Species[];
    @Getter('defaultTypes', { namespace })
    private defaultTypes: string;
    @Getter('defaultTotal', { namespace })
    private defaultTotal: number;

    private types: string = 'すべて';
    private total: number = 0;

    created() {
        this.types = this.defaultTypes;
        this.total = this.defaultTotal;
    }

    @Watch('total')
    totalChanged(after: number, before: number) {
        if (before == after) {
            return;
        }
        this.updateCandidates();
    }
    @Watch('types')
    typesChanged(after: string, before: string) {
        if (after == before) {
            return;
        }
        this.updateCandidates();
    }

    @Watch('candidates')
    candidatesChanged() {
        console.log('changed.');
        for (var i = 0; i < this.candidates.length; i++) {
            let c = this.candidates[i];
            console.log('' + c.name + ' types:' + c.types);
        }
    }

    setType(type: string) {
        this.types = type;
    }

    updateCandidates() {
        console.log('Update.');
        this.getCandidates({
            types: this.types,
            total: this.total,
        });
    }

    onButton(candidate: Species) {
        console.log('Select:' + candidate.name);
        // TODO:攻撃調整に遷移
        //router.push({ path: 'sandboxts' });
    }
}
</script>

<style scoped>
.candidates {
    width: 600px;
    height: 800px;
    overflow: scroll;
}
.targetSelect {
    width: 480px;
}
</style>

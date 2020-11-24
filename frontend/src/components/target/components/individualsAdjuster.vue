/*
    個体値調整
*/
<template>
    <div class="individuals">
        <p>個体値 HP:{{individuals.hp()}} 攻撃:{{individuals.attack()}} 防御:{{individuals.defense()}} 特攻:{{individuals.spAttack()}} 特防:{{individuals.spDefense()}} 素早さ:{{individuals.speed()}}</p>
        <b-form-checkbox id="check-slowest" v-model="slowest" name="check-slowest">最遅</b-form-checkbox>
        <b-form-checkbox id="check-weakest" v-model="weakest" name="check-weakest">最弱</b-form-checkbox>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from 'vue-property-decorator';
import { Individuals } from '../store/individuals'

@Component
export default class IndividualsAdjuster extends Vue {
    private slowest: boolean = false;
    private weakest: boolean = false;
    @Prop() private individuals!: Individuals;

    created() {
        this.slowest = this.individuals.isSlowest();
        this.weakest = this.individuals.isWeakest();
    }

    @Watch('slowest')
    private slowestChanged() {
        // this.individuals.slowest()を直接呼び出しても動作する
        // TODO:ことの問題点は
        // 状態変更はstoreに任せるべきか
        this.$emit('slowest', this.slowest);
    }
    @Watch('weakest')
    private weakestChanged() {
        this.$emit('weakest', this.weakest);
    }
}
</script>

<style scoped>
</style>

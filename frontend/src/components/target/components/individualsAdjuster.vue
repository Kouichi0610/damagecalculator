/*
    個体値調整
*/
<template>
    <div class="individuals">
        <p>個体値 HP:{{individuals.hp}} 攻撃:{{individuals.at}} 防御:{{individuals.df}} 特攻:{{individuals.sa}} 特防:{{individuals.sd}} 素早さ:{{individuals.sp}}</p>
        <b-form-checkbox id="check-slowest" v-model="slowest" name="check-slowest">最遅</b-form-checkbox>
        <b-form-checkbox id="check-weakest" v-model="weakest" name="check-weakest">最弱</b-form-checkbox>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from 'vue-property-decorator';

@Component
export default class IndividualsAdjuster extends Vue {
    private slowest: boolean = false;
    private weakest: boolean = false;
    @Prop() private individuals!: Individuals;

    created() {
        this.slowest = this.individuals.sp == 0;
        this.weakest = this.individuals.at == 0;
    }

    @Watch('slowest')
    private slowestChanged() {
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

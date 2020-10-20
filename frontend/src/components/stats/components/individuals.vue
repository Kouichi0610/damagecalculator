/*
    個体値コンポーネント
    
*/
<template>
    <div class="individuals">
        <p>個体値 HP:{{disp[0]}} 攻撃:{{disp[1]}} 防御:{{disp[2]}} 特攻:{{disp[3]}} 特防:{{disp[4]}} 素早さ:{{disp[5]}}</p>
        <b-form-checkbox id="checkbox-1" v-model="slowest" name="checkbox-1">最遅</b-form-checkbox>
        <b-form-checkbox id="checkbox-2" v-model="weakest" name="checkbox-2">最弱</b-form-checkbox>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from 'vue-property-decorator';

@Component
export default class Individuals extends Vue {
    private slowest: boolean = false;
    private weakest: boolean = false;
    @Prop() private params!: number[];

    private disp: number[] = [31, 31, 31, 31, 31, 31];

    created() {
        this.disp = this.params;
        this.slowest = this.params[5] == 0;
        this.weakest = this.params[1] == 0;
    }

    @Watch('slowest')
    private slowestChanged() {
        this.disp[5] = this.slowest ? 0 : 31;
        this.$emit('changed', this.disp);
    }
    @Watch('weakest')
    private weakestChanged() {
        this.disp[1] = this.weakest ? 0 : 31;
        this.$emit('changed', this.disp);
    }
}
</script>

<style scoped>
</style>

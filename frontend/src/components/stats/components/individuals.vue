/*
    個体値コンポーネント
    
*/
<template>
    <div class="individuals">
        <p>個体値 HP:{{params[0]}} 攻撃:{{params[1]}} 防御:{{params[2]}} 特攻:{{params[3]}} 特防:{{params[4]}} 素早さ:{{params[5]}}</p>
        <b-form-checkbox id="checkbox-1" v-model="slowest" name="checkbox-1">最遅</b-form-checkbox>
        <b-form-checkbox id="checkbox-2" v-model="weakest" name="checkbox-2">最弱</b-form-checkbox>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';

@Component
export default class Individuals extends Vue {
    private slowest: boolean = false;
    private weakest: boolean = false;
    private params: number[] = [31, 31, 31, 31, 31, 31];

    @Watch('slowest')
    private slowestChanged() {
        this.params[5] = this.slowest ? 0 : 31;
        this.$emit('changed', this.params);
    }
    @Watch('weakest')
    private weakestChanged() {
        this.params[1] = this.weakest ? 0 : 31;
        this.$emit('changed', this.params);
    }
}
</script>

<style scoped>
</style>

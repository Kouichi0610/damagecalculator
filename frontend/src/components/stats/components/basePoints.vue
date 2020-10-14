/*
    基礎ポイント
    HP,攻撃,防御,特攻,特防,素早さそれぞれに割り振ることができる
    最大508まで
    各252まで

    TODO:ボタンによる増減
    TODO:縦に
    TODO:v-for
*/
<template>
    <div class="basePoints">
        <p>基礎ポイント:{{ total }}</p>
            <fieldset>
                <div class="row mb-1">
                    <BasePointSlider v-bind:values="vals" index="0" @value-changed="valueChanged"></BasePointSlider><br>
                    <BasePointSlider v-bind:values="vals" index="1" @value-changed="valueChanged"></BasePointSlider><br>
                    <BasePointSlider v-bind:values="vals" index="2" @value-changed="valueChanged"></BasePointSlider><br>
                    <BasePointSlider v-bind:values="vals" index="3" @value-changed="valueChanged"></BasePointSlider><br>
                    <BasePointSlider v-bind:values="vals" index="4" @value-changed="valueChanged"></BasePointSlider><br>
                    <BasePointSlider v-bind:values="vals" index="5" @value-changed="valueChanged"></BasePointSlider><br>
                </div>
            </fieldset>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import BasePointSlider from './basePointSlider.vue';

@Component({
      components: {
        BasePointSlider,
    },
})
export default class BasePoints extends Vue {
    private vals: number[] = [0, 0, 0, 0, 0, 0];
    private total: number = 0;

    private valueChanged(index: number, val: number) {
        this.vals[index] = val;
        this.total = 0;
        for (var i = 0; i < 6; i++) {
            this.total += this.vals[i];
        }
        this.$emit('changed', this.vals);
    }
}
</script>

<style scoped>
</style>

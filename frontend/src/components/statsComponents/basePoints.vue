/*
    基礎ポイント
    HP,攻撃,防御,特攻,特防,素早さそれぞれに割り振ることができる
    最大508まで
    各252まで

    TODO:v-for対応
    TODO:ボタンによる増減
*/
<template>
    <div class="basePoints">
        <p>基礎ポイント:{{ total }}</p>

        <fieldset>
            <div class="row mb-1">
                <base-point-slider v-bind:tag="tags[0]" v-bind:values="vals" index="0" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
                <base-point-slider v-bind:tag="tags[1]" v-bind:values="vals" index="1" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
                <base-point-slider v-bind:tag="tags[2]" v-bind:values="vals" index="2" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
                <base-point-slider v-bind:tag="tags[3]" v-bind:values="vals" index="3" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
                <base-point-slider v-bind:tag="tags[4]" v-bind:values="vals" index="4" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
                <base-point-slider v-bind:tag="tags[5]" v-bind:values="vals" index="5" fixed="false" @value-changed="valueChanged"></base-point-slider><br>
            </div>
        </fieldset>
    </div>
</template>

<script>
import BasePointSlider from './basePointSlider.vue'

export default {
    name: 'basePoints',
    components: {
        'base-point-slider': BasePointSlider,
    },
    props: ['params'],
    created: function() {
        this.vals = this.params;
    },
    data: function() {
        return {
            values: [
                {index:0, value:this.params[0]},
                {index:1, value:this.params[1]},
                {index:2, value:this.params[2]},
                {index:3, value:this.params[3]},
                {index:4, value:this.params[4]},
                {index:5, value:this.params[5]},
                ],
            vals: [0, 0, 0, 0, 0, 0],
            tags: ['HP','攻撃','防御','特攻','特防','素早'],
            hp: 0,
            total: 0,
        }
    },
    methods: {
        valueChanged: function(v, i) {
            this.vals[i] = v;
            this.total = this.calcTotal();
            this.$emit('update', this.vals);
        },
        calcTotal: function() {
            let sum = function(arr) {
                return arr.reduce(function(prev, current) {
                    return prev + current;
                });
            };
            return sum(this.vals);
        }
    },
    computed: {
    },
    watch: {
    }
}
</script>

<style scoped>
</style>

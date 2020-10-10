/*
    ダメージ結果
    割合
    乱数
    急所時
*/
/*
    :max="max"
    v-bind:max="max"のシュガーシンタックス

*/
<template>
    <span class="damageResult">
        ダメージ結果
        <b-progress class="mt-2" :max="progressMax" show-value animated>
            <b-progress-bar :value="min" :label="alpha" variant="danger"></b-progress-bar>
            <b-progress-bar :value="max - min" variant="warning"></b-progress-bar>
        </b-progress>
        割合{{ min }}～{{ max }}%
    </span>
</template>

<script>
export default {
    name: 'damageResult',
    props:['min', 'max'],
    data() {
        return {
            value: 45,
            progressMax: 100,
            alpha: 'test' + 15,
        }
    },
    methods: {
    }

}
</script>

<style scoped>
</style>

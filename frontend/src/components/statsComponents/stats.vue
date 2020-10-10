// 能力値コンポーネント
// TODO:Slider使う
<template>
    <div class="stats">
        <p>能力値 HP:{{params[0]}} 攻撃:{{params[1]}} 防御:{{params[2]}}  特攻:{{params[3]}} 特防:{{params[4]}} 素早さ:{{params[5]}}</p>
        <stats-progress label="HP:" :current="params[0]"></stats-progress>
        <stats-progress label="攻撃:" :current="params[1]"></stats-progress>
        <stats-progress label="防御:" :current="params[2]"></stats-progress>
        <stats-progress label="特攻:" :current="params[3]"></stats-progress>
        <stats-progress label="特防:" :current="params[4]"></stats-progress>
        <stats-progress label="素早さ:" :current="params[5]"></stats-progress>
    </div>
</template>

<script>
import StatsProgress from './statsProgress'

export default {
    name: 'stats',
    components: {
        'stats-progress': StatsProgress,
    },
    props: ['params'],
}
</script>

<style scoped>
.stats {
    width: 500px;
}
</style>
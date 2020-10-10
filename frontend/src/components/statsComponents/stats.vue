// 能力値コンポーネント
// TODO:Slider使う
<template>
    <div class="stats">
        <p>能力値 HP:{{params[0]}} 攻撃:{{params[1]}} 防御:{{params[2]}}  特攻:{{params[3]}} 特防:{{params[4]}} 素早さ:{{params[5]}}</p>
    </div>
</template>

<script>
export default {
    name: 'stats',
    props: ['params'],
}
</script>

<style scoped>
</style>
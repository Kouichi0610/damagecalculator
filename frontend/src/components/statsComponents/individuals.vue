/*
    個体値コンポーネント
    めざめるパワーを考慮する必要はなく、基本的にはすべてさいこうを前提とする。
    ジャイロボール、トリックルーム->最遅
    イカサマ対策->最弱

    TODO:v-form-groupについて
*/
<template>
    <div class="individuals">
        <p>{{ params }}</p>
        <p>個体値 HP:{{params[0]}} 攻撃:{{params[1]}} 防御:{{params[2]}} 特攻:{{params[3]}} 特防:{{params[4]}} 素早さ:{{params[5]}}</p>
        <b-form-checkbox
        id="checkbox-1"
        v-model="slowest"
        name="checkbox-1"
        >
        最遅
        </b-form-checkbox>
        <b-form-checkbox
        id="checkbox-2"
        v-model="weakest"
        name="checkbox-2"
        >
        最弱
        </b-form-checkbox>
    </div>
</template>

<script>
export default {
    name: 'individuals',
    props: ['params'],
    data: function() {
        return {
            slowest: false,
            weakest: false,
        }
    },
    created: function() {
    },
    watch: {
        slowest: function() {
            this.$emit('update', this.slowest, this.weakest);
        },
        weakest: function() {
            this.$emit('update', this.slowest, this.weakest);
        }
    },
}
</script>

<style scoped>

</style>
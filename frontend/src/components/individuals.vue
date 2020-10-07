/*
    個体値コンポーネント
    めざめるパワーを考慮する必要はなく、基本的にはすべてさいこうを前提とする。
    ジャイロボール、トリックルーム->最遅
    イカサマ対策->最弱

    TODO:v-form-groupについて
*/
<template>
    <span class="individuals">
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
        {{ toString }}
    </span>
</template>

<script>
export default {
    name: 'individuals',
    data: function() {
        return {
            slowest: false,
            weakest: false,
            hp:31,
            at:31,
            df:31,
            sa:31,
            sd:31,
            sp:31,
        }
    },
    watch: {
        slowest: function() {
            this.sp = this.slowest ? 0 : 31;
            console.log('slowest:' + this.slowest);

        },
        weakest: function() {
            this.at = this.weakest ? 0 : 31;
        }
    },
    computed: {
        toString: function() {
            return '個体値 HP:' + this.hp + ' 攻撃:' + this.at + ' 防御:' + this.df + ' 特攻:' + this.sa + '特防:' + this.sd + ' 素早さ:' + this.sp;
        }
    }
}
</script>

<style scoped>

</style>
/*
    個体値スライダー
*/
<template>
    <span class="basePointSlider">
        {{ tag }}:<input type="range" id="slider" v-model.number="current" step="4" min="0" max="252">{{ current }}
    </span>
</template>

<script>
export default {
    name: 'basePointSlider',
    data: function() {
        return {
            current: 0,
        }
    },
    // パラメータ名、初期値一覧、インデックス
    // propsは直接書き換えられない ($emitで変更通知)
    props: ['tag', 'values', 'index'],
    created: function() {
        console.log('tag:' + this.tag + ' values:' + this.values + ' index:' + this.index);
        this.current = this.values[this.index];
    },
    watch: {
        current: function() {
            const max = 508;
            var total = 0;
            for (var i = 0; i < 6; i++) {
                if (i == this.index) {
                    continue;
                }
                total += this.values[i];
            }
            let remain = max - total;
            if (this.current > remain) {
                this.current = remain;
            }
            this.$emit('value-changed', this.current, this.index);
        }
    },
    methods: {
    }
    
}
</script>

<style scoped>
</style>
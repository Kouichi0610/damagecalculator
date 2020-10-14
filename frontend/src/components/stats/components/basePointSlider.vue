// 基礎ポイントスライダー
<template>
    <span class="row mb-1">
        <div class="col-8">
            <input type="range" class="form-control-change" id="slider" v-model.number="current" step="4" min="0" max="252" :disabled="disabled">
        </div>
        <div class="col-3">{{ current }}</div>
    </span>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';

@Component
export default class BasePointSlider extends Vue {
    @Prop()
    private values!: number[];
    @Prop()
    private index!: number;

    private current: number = 0;
    private disabled: boolean = false;

    @Watch('current')
    public currentChanged(/*newVal: number*/) {
        const max = 508;
        var total: number = 0;
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
        this.$emit('value-changed', this.index, this.current);
    }
    /*
    name: 'basePointSlider',
    data: function() {
        return {
            current: 0,
            disabled:false,
        }
    },
    // パラメータ名、初期値一覧、インデックス
    // propsは直接書き換えられない ($emitで変更通知)
    props: ['tag', 'values', 'index', 'fixed'],
    created: function() {
        this.disabled = (this.fixed == 'true');
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
    */
}
</script>

<style scoped>
</style>
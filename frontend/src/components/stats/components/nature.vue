// 性格設定コンポーネント
<template>
    <div class="nature">
        <b-form-select v-model="current" :options="natures" size="sm" class="mt-3"></b-form-select>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';

import { NatureFactory, NatureNames } from '../../../domain/nature'

@Component
export default class Nature extends Vue {
    @Prop() private selected!: string;
    private current: string = "てれや";
    private natures: string[];

    created(): void {
        this.current = this.selected;
        this.natures = NatureNames();
    }

    @Watch('current')
    public currentChanged(newVal: string/*, oldVal: string*/) {
        let n = NatureFactory(newVal);
        this.$emit("changed", n);
    }    
}
</script>

<style scoped>
</style>

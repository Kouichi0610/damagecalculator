// 性格設定コンポーネント
<template>
    <div class="nature">
        <!-- {{ selected.Name() }} -->
        <b-form-select @change="changed" v-model="current" :options="natures" size="sm" class="mt-3"></b-form-select>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

import * as nature from '../../../domain/nature'

@Component
export default class Nature extends Vue {
    private current: string;
    private natures: string[];

    @Prop() private selected!: nature.INature;

    created(): void {
        this.current = this.selected.Name();
        this.natures = nature.NatureNames();
        console.log('start:' + this.current);
    }

    // dropdown変更時
    changed(val: string) {
        this.current = val;

        let n = nature.NatureFactory(this.current);
        console.log('Nature to ' + n.Name);
        this.$emit("changed", n);
    }
}
</script>

<style scoped>
</style>

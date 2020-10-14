// TypeScriptコンポーネント試作
<template>
    <div class="hellochild">
        <button>{{ buttonName }}</button>
        <h2>{{ samplemessage }}</h2>

    </div>
    
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class HelloChild extends Vue {
    private buttonName: string
    @Prop() private samplemessage!: string;
    created() {
        this.buttonName = 'MyButton.';
    }
}
</script>
<style scoped>
</style>
